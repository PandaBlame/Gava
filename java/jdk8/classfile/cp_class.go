package classfile

import "github.com/PandaBlame/Gava/java"

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfoV8 struct {
	cp        java.ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfoV8) Read(reader *java.ClassReader) (err error) {
	self.nameIndex, err = reader.ReadUint16()
	return
}
func (self *ConstantClassInfoV8) Name() (string, error) {
	return self.cp.GetUtf8(self.nameIndex)
}
