package classpath

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/PandaBlame/Gava/java"
)

type JavaSourceJar struct {
	*JavaSourceInfo
}

func (self *JavaSourceJar) ReadClass(className string, parser java.ClassFileParser) (has bool, classFile java.ClassFile, err error) {
	has, data, err := self.ReadFile(className + ".class")
	if !has || err != nil {
		return
	}
	classFile, err = parser.ParseClassFile(ioutil.NopCloser(bytes.NewReader(data)))
	return
}

func (self *JavaSourceJar) ReadFile(path string) (has bool, data []byte, err error) {
	absPath, err := filepath.Abs(self.FilePath)
	if err != nil {
		return
	}
	zipRC, err := zip.OpenReader(absPath)
	if err != nil {
		return
	}
	defer zipRC.Close()
	for _, f := range zipRC.File {
		if f.Name == path {
			has = true
			file, err := f.Open()
			if err != nil {
				return has, nil, err
			}
			data, err = ioutil.ReadAll(file)
			return has, data, err
		}
	}
	return
}
