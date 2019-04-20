package java

import (
	"encoding/binary"
	"io"
)

type ClassFile interface {
	Read(reader *ClassReader) (err error)
	MinorVersion() uint16
	MajorVersion() uint16
	ConstantPool() ConstantPool
	AccessFlags() uint16
	Fields() []MemberInfo
	Methods() []MemberInfo
	ClassName() (string, error)
	SuperClassName() (string, error)
	InterfaceNames() ([]string, error)
	SourceFileAttribute() SourceFileAttribute
}

type ConstantInfo interface {
	Read(reader *ClassReader) error
}

type ConstantPool interface {
	Read(reader *ClassReader) error
	GetConstantInfo(index uint16) (ConstantInfo, error)
	GetNameAndType(index uint16) (name string, _type string, err error)
	GetClassName(index uint16) (cn string, err error)
	GetUtf8(index uint16) (str string, err error)
}

type MemberInfo interface {
	Read(reader *ClassReader) error
	AccessFlags() uint16
	Name() (string, error)
	Descriptor() (string, error)
	CodeAttribute() CodeAttribute
	ConstantValueAttribute() ConstantValueAttribute
	ExceptionsAttribute() ExceptionsAttribute
	RuntimeVisibleAnnotationsAttributeData() []byte
	RuntimeVisibleParameterAnnotationsAttributeData() []byte
	AnnotationDefaultAttributeData() []byte
	UnParsedAttributeData(name string) []byte
}

type AttributeInfo interface {
	Read(reader *ClassReader) error
}

type CodeAttribute interface {
	Read(reader *ClassReader) error
	MaxStack() uint
	MaxLocals() uint
	Code() []byte
	ExceptionTable() []ExceptionTableEntry
	LineNumberTableAttribute() LineNumberTableAttribute
}

type ExceptionTableEntry interface {
	Read(reader *ClassReader) error
	StartPc() uint16
	EndPc() uint16
	HandlerPc() uint16
	CatchType() uint16
}

type LineNumberTableEntry interface {
	Read(reader *ClassReader) error
	StartPc() uint16
	LineNumber() uint16
}

type LineNumberTableAttribute interface {
	Read(reader *ClassReader) error
	GetLineNumber(pc int) int
}

type ConstantValueAttribute interface {
	Read(reader *ClassReader) error
	ConstantValueIndex() uint16
}

type ExceptionsAttribute interface {
	Read(reader *ClassReader) error
	ExceptionIndexTable() []uint16
}

type SourceFileAttribute interface {
	Read(reader *ClassReader) error
	FileName() (string, error)
}

type EnclosingMethodAttribute interface {
	Read(reader *ClassReader) error
	ClassName() (string, error)
	MethodNameAndDescriptor() (string, string, error)
}

type SignatureAttribute interface {
	Read(reader *ClassReader) error
	Signature() (string, error)
}

type UnParsedAttribute interface {
	Read(reader *ClassReader) error
	Name() string
	Info() []byte
}

type ClassReader struct {
	Data io.ReadCloser
}

func (self *ClassReader) ReadUint8() (uint8, error) {
	var val uint8
	err := binary.Read(self.Data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint16() (uint16, error) {
	var val uint16
	err := binary.Read(self.Data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint32() (uint32, error) {
	var val uint32
	err := binary.Read(self.Data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint64() (uint64, error) {
	var val uint64
	err := binary.Read(self.Data, binary.BigEndian, &val)
	return val, err
}

func (self *ClassReader) ReadUint16s() ([]uint16, error) {
	n, err := self.ReadUint16()
	if err != nil {
		return nil, err
	}
	val := make([]uint16, n)
	err = binary.Read(self.Data, binary.BigEndian, &val)
	return val, nil
}

func (self *ClassReader) ReadBytes(n uint32) ([]byte, error) {
	val := make([]byte, n)
	_, err := self.Data.Read(val)
	return val, err
}

type ClassFileParser interface {
	ParseClassFile(data io.ReadCloser) (cf ClassFile, err error)
}
