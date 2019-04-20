package classfile

import "github.com/PandaBlame/Gava/java"

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttributeV8 struct {
	cp              java.ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.sourceFileIndex, err = reader.ReadUint16()
	return
}

func (self *SourceFileAttributeV8) FileName() (string, error) {
	return self.cp.GetUtf8(self.sourceFileIndex)
}
