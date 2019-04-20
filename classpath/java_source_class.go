package classpath

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/PandaBlame/Gava/java"
)

type JavaSourceClass struct {
	*JavaSourceInfo
	className    string
	m            *sync.Mutex
	isFinishFlag uint32
}

func (self *JavaSourceClass) isFinish() bool {
	return atomic.LoadUint32(&self.isFinishFlag) == 1
}

func (self *JavaSourceClass) load(parser java.ClassFileParser) (classFile java.ClassFile, err error) {
	if self.isFinish() {
		return
	}
	self.m.Lock()
	defer self.m.Unlock()
	if self.isFinishFlag == 0 {
		defer atomic.StoreUint32(&self.isFinishFlag, 1)
		self.className = ""
		absPath, err := filepath.Abs(self.FilePath)
		if err != nil {
			return nil, err
		}
		file, err := os.Open(absPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		classFile, err = parser.ParseClassFile(file)
		if err != nil {
			return classFile, err
		}
	}
	return
}

func (self *JavaSourceClass) ReadClass(className string, parser java.ClassFileParser) (has bool, classFile java.ClassFile, err error) {
	has = className == self.className
	if self.isFinish() && !has {
		return
	}
	classFile, err = self.load(parser)
	if err != nil {
		return
	}
	self.className, err = classFile.ClassName()
	if err != nil {
		return
	}
	has = className == self.className
	return
}

func (self *JavaSourceClass) ReadFile(path string) (has bool, data []byte, err error) {
	return
}
