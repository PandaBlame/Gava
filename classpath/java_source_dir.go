package classpath

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/PandaBlame/Gava/java"
)

type JavaSourceDir struct {
	*JavaSourceInfo
}

func (self *JavaSourceDir) ReadClass(className string, parser java.ClassFileParser) (has bool, classFile java.ClassFile, err error) {
	has, data, err := self.ReadFile(className + ".class")
	if !has || err != nil {
		return
	}
	classFile, err = parser.ParseClassFile(ioutil.NopCloser(bytes.NewReader(data)))
	return
}

func (self *JavaSourceDir) ReadFile(path string) (has bool, data []byte, err error) {
	absPath, err := filepath.Abs(self.FilePath)
	if err != nil {
		return
	}
	fPath := filepath.Join(absPath, path)
	data, err = ioutil.ReadFile(fPath)
	has = err == nil
	return
}
