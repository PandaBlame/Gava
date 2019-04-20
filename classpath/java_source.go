package classpath

import "github.com/PandaBlame/Gava/java"

type JavaSource interface {
	ReadFile(path string) (has bool, data []byte, err error)
	ReadClass(className string, parser java.ClassFileParser) (has bool, classFile java.ClassFile, err error)
}

type JavaSourceInfo struct {
	FilePath string
}
