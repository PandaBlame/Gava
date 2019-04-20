package classfile

import (
	"io"

	"github.com/PandaBlame/Gava/java"
)

type ClassFileParserV8 struct {
}

func (self ClassFileParserV8) ParseClassFile(data io.ReadCloser) (cf java.ClassFile, err error) {
	defer data.Close()
	cr := &java.ClassReader{Data: data}
	cf = &ClassFileV8{}
	err = cf.Read(cr)
	return
}
