package classfile

import "github.com/PandaBlame/Gava/java"

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfoV8 struct {
	cp              java.ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []java.AttributeInfo
}

// read field or method table
func readMembers(reader *java.ClassReader, cp java.ConstantPool) (members []java.MemberInfo, err error) {
	memberCount, err := reader.ReadUint16()
	if err != nil {
		return
	}
	members = make([]java.MemberInfo, memberCount)
	for i := range members {
		members[i], err = readMember(reader, cp)
		if err != nil {
			return
		}
	}
	return
}

func readMember(reader *java.ClassReader, cp java.ConstantPool) (memberInfo java.MemberInfo, err error) {
	memberInfo = &MemberInfoV8{
		cp: cp,
	}
	err = memberInfo.Read(reader)
	return
}

func (self *MemberInfoV8) Read(reader *java.ClassReader) (err error) {
	self.accessFlags, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.nameIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.descriptorIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.attributes, err = readAttributes(reader, self.cp)
	return
}

func (self *MemberInfoV8) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfoV8) Name() (string, error) {
	return self.cp.GetUtf8(self.nameIndex)
}
func (self *MemberInfoV8) Descriptor() (string, error) {
	return self.cp.GetUtf8(self.descriptorIndex)
}

func (self *MemberInfoV8) CodeAttribute() java.CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.CodeAttribute:
			return attrInfo.(java.CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfoV8) ConstantValueAttribute() java.ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.ConstantValueAttribute:
			return attrInfo.(java.ConstantValueAttribute)
		}
	}
	return nil
}

func (self *MemberInfoV8) ExceptionsAttribute() java.ExceptionsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.ExceptionsAttribute:
			return attrInfo.(java.ExceptionsAttribute)
		}
	}
	return nil
}

func (self *MemberInfoV8) RuntimeVisibleAnnotationsAttributeData() []byte {
	return self.UnParsedAttributeData("RuntimeVisibleAnnotations")
}
func (self *MemberInfoV8) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return self.UnParsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (self *MemberInfoV8) AnnotationDefaultAttributeData() []byte {
	return self.UnParsedAttributeData("AnnotationDefault")
}

func (self *MemberInfoV8) UnParsedAttributeData(name string) []byte {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.UnParsedAttribute:
			unparsedAttr := attrInfo.(java.UnParsedAttribute)
			if unparsedAttr.Name() == name {
				return unparsedAttr.Info()
			}
		}
	}
	return nil
}
