package classfile

import (
	"errors"
	"fmt"

	"github.com/PandaBlame/Gava/java"
	"github.com/PandaBlame/Gava/java/jdk8"
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFileV8 struct {
	minorVersion uint16
	majorVersion uint16
	constantPool java.ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []java.MemberInfo
	methods      []java.MemberInfo
	attributes   []java.AttributeInfo
}

func (self *ClassFileV8) Read(reader *java.ClassReader) (err error) {
	err = self.readAndCheckMagic(reader)
	if err != nil {
		return
	}
	err = self.readAndCheckVersion(reader)
	if err != nil {
		return
	}
	self.constantPool, err = readConstantPool(reader)
	if err != nil {
		return
	}
	self.accessFlags, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.thisClass, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.superClass, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.interfaces, err = reader.ReadUint16s()
	if err != nil {
		return
	}
	self.fields, err = readMembers(reader, self.constantPool)
	if err != nil {
		return
	}
	self.methods, err = readMembers(reader, self.constantPool)
	if err != nil {
		return
	}
	self.attributes, err = readAttributes(reader, self.constantPool)
	return
}

func (self *ClassFileV8) readAndCheckMagic(reader *java.ClassReader) (err error) {
	magic, err := reader.ReadUint32()
	if err != nil {
		return
	}
	if magic != 0xCAFEBABE {
		err = errors.New("java.lang.ClassFormatError: magic")
	}
	return
}

func (self *ClassFileV8) readAndCheckVersion(reader *java.ClassReader) (err error) {
	self.minorVersion, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.majorVersion, err = reader.ReadUint16()
	if err != nil {
		return
	}
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == jdk8.JVM_CLASSFILE_MINOR_VERSION {
			return
		}
	}
	err = errors.New(fmt.Sprintf("java.lang.UnsupportedClassVersionError: %v", self.majorVersion))
	return
}

func (self *ClassFileV8) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFileV8) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFileV8) ConstantPool() java.ConstantPool {
	return self.constantPool
}
func (self *ClassFileV8) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFileV8) Fields() []java.MemberInfo {
	return self.fields
}
func (self *ClassFileV8) Methods() []java.MemberInfo {
	return self.methods
}

func (self *ClassFileV8) ClassName() (string, error) {
	return self.constantPool.GetClassName(self.thisClass)
}

func (self *ClassFileV8) SuperClassName() (string, error) {
	if self.superClass > 0 {
		return self.constantPool.GetClassName(self.superClass)
	}
	return "", nil
}

func (self *ClassFileV8) InterfaceNames() ([]string, error) {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		in, err := self.constantPool.GetClassName(cpIndex)
		if err != nil {
			return nil, err
		}
		interfaceNames[i] = in
	}
	return interfaceNames, nil
}

func (self *ClassFileV8) SourceFileAttribute() java.SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.SourceFileAttribute:
			return attrInfo.(java.SourceFileAttribute)
		}
	}
	return nil
}
