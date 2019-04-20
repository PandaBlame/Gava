package classfile

import "github.com/PandaBlame/Gava/java"

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfoV8 struct {
	cp          java.ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfoV8) Read(reader *java.ClassReader) (err error) {
	self.stringIndex, err = reader.ReadUint16()
	return
}
func (self *ConstantStringInfoV8) String() (string, error) {
	return self.cp.GetUtf8(self.stringIndex)
}
