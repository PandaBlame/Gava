package classfile

import (
	"fmt"

	"github.com/PandaBlame/Gava/java"
)

type ConstantPoolV8 struct {
	infos []java.ConstantInfo
}

func readConstantPool(reader *java.ClassReader) (java.ConstantPool, error) {
	cp := &ConstantPoolV8{}
	err := cp.Read(reader)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func (self *ConstantPoolV8) Read(reader *java.ClassReader) error {
	cpCount, err := reader.ReadUint16()
	if err != nil {
		return err
	}
	self.infos = make([]java.ConstantInfo, int(cpCount))
	for i := 1; i < int(cpCount); i++ {
		self.infos[i], err = readConstantInfo(reader, self)
		if err != nil {
			return err
		}
		switch self.infos[i].(type) {
		case *ConstantLongInfoV8, *ConstantDoubleInfoV8:
			i++
		}
	}
	return nil
}

func (self *ConstantPoolV8) GetConstantInfo(index uint16) (java.ConstantInfo, error) {
	if cpInfo := self.infos[index]; cpInfo != nil {
		return cpInfo, nil
	}
	return nil, fmt.Errorf("java.lang.ClassFormatError: Invalid constant pool index: %v", index)
}

func (self *ConstantPoolV8) GetNameAndType(index uint16) (name string, _type string, err error) {
	cInfo, err := self.GetConstantInfo(index)
	if err != nil {
		return
	}
	switch cInfo.(type) {
	case *ConstantNameAndTypeInfoV8:
		ntInfo := cInfo.(*ConstantNameAndTypeInfoV8)
		name, err = self.GetUtf8(ntInfo.nameIndex)
		if err != nil {
			return
		}
		_type, err = self.GetUtf8(ntInfo.descriptorIndex)
		return
	default:
		return "", "", fmt.Errorf("java.lang.ClassFormatError: Wrong type constant pool index: %v", index)
	}
}

func (self *ConstantPoolV8) GetClassName(index uint16) (cn string, err error) {
	cInfo, err := self.GetConstantInfo(index)
	if err != nil {
		return
	}
	switch cInfo.(type) {
	case *ConstantClassInfoV8:
		classInfo := cInfo.(*ConstantClassInfoV8)
		return self.GetUtf8(classInfo.nameIndex)
	default:
		return "", fmt.Errorf("java.lang.ClassFormatError: Wrong type constant pool index: %v", index)
	}
}

func (self *ConstantPoolV8) GetUtf8(index uint16) (str string, err error) {
	cInfo, err := self.GetConstantInfo(index)
	if err != nil {
		return
	}
	switch cInfo.(type) {
	case *ConstantUtf8InfoV8:
		utf8Info := cInfo.(*ConstantUtf8InfoV8)
		return utf8Info.str, nil
	default:
		return "", fmt.Errorf("java.lang.ClassFormatError: Wrong type constant pool index: %v", index)
	}
}
