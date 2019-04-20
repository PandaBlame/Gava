package classfile

import "github.com/PandaBlame/Gava/java"

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfoV8 struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfoV8) Read(reader *java.ClassReader) (err error) {
	self.nameIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.descriptorIndex, err = reader.ReadUint16()
	return
}
