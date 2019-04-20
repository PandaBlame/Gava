package classfile

import "github.com/PandaBlame/Gava/java"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnParsedAttributeV8 struct {
	name   string
	length uint32
	info   []byte
}

func (self *UnParsedAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.info, err = reader.ReadBytes(self.length)
	return
}

func (self *UnParsedAttributeV8) Name() string {
	return self.name
}

func (self *UnParsedAttributeV8) Info() []byte {
	return self.info
}
