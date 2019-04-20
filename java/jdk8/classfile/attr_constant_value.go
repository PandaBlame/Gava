package classfile

import "github.com/PandaBlame/Gava/java"

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttributeV8 struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.constantValueIndex, err = reader.ReadUint16()
	return
}

func (self *ConstantValueAttributeV8) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
