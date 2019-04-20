package classfile

import "github.com/PandaBlame/Gava/java"

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldRefInfoV8 struct{ ConstantMemberRefInfoV8 }
type ConstantMethodRefInfoV8 struct{ ConstantMemberRefInfoV8 }
type ConstantInterfaceMethodRefInfoV8 struct{ ConstantMemberRefInfoV8 }

type ConstantMemberRefInfoV8 struct {
	cp               java.ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfoV8) Read(reader *java.ClassReader) (err error) {
	self.classIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.nameAndTypeIndex, err = reader.ReadUint16()
	return
}

func (self *ConstantMemberRefInfoV8) ClassName() (string, error) {
	return self.cp.GetClassName(self.classIndex)
}

func (self *ConstantMemberRefInfoV8) NameAndDescriptor() (string, string, error) {
	return self.cp.GetNameAndType(self.nameAndTypeIndex)
}
