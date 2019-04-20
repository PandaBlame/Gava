package classpath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/PandaBlame/Gava/java"
)

type Classpath struct {
	jreSources    []JavaSource
	jreExtSources []JavaSource
	userSources   []JavaSource
	JreLibPath    string
	UserLibPath   string
	Parser        java.ClassFileParser
}

func (self *Classpath) Load() (err error) {
	return self.parseClasspath()
}

func (self *Classpath) ReadClass(className string) (cf java.ClassFile, err error) {
	if has, cf, err := self.read(self.jreSources, className); has {
		return cf, err
	}
	if has, cf, err := self.read(self.jreExtSources, className); has {
		return cf, err
	}
	if has, cf, err := self.read(self.userSources, className); has {
		return cf, err
	}
	return nil, nil
}

func (self *Classpath) read(sources []JavaSource, className string) (has bool, cf java.ClassFile, err error) {
	for _, source := range sources {
		has, cf, err = source.ReadClass(className, self.Parser)
		if has {
			return has, cf, err
		}
	}
	return
}

func (self *Classpath) parseClasspath() (err error) {
	if err = self.getJreDir(); err != nil {
		return
	}
	err = self.parseJreLib()
	if err != nil {
		return
	}
	err = self.parseJreExtLib()
	if err != nil {
		return
	}
	err = self.parseUserLib()
	return
}

func (self *Classpath) parseJreLib() (err error) {
	jreLibPath := filepath.Join(self.JreLibPath, "lib", "*")
	sources, err := self.parseWildcardSource(jreLibPath)
	if err != nil {
		return err
	}
	self.jreSources = append(self.jreSources, sources...)
	return
}

func (self *Classpath) parseJreExtLib() (err error) {
	jreExtPath := filepath.Join(self.JreLibPath, "lib", "ext", "*")
	sources, err := self.parseWildcardSource(jreExtPath)
	if err != nil {
		return err
	}
	self.jreExtSources = append(self.jreExtSources, sources...)
	return
}

func (self *Classpath) parseUserLib() (err error) {
	if self.UserLibPath == "" {
		self.UserLibPath = "." + string(os.PathSeparator) + "*"
	}
	for _, path := range strings.Split(self.UserLibPath, string(os.PathListSeparator)) {
		if strings.HasSuffix(path, ".class") || strings.HasSuffix(path, ".CLASS") {
			source, err := self.parseClassSource(path)
			if err != nil {
				return err
			}
			self.userSources = append(self.userSources, source)
		}
		if strings.HasSuffix(path, "*") {
			sources, err := self.parseWildcardSource(path)
			if err != nil {
				return err
			}
			self.userSources = append(self.userSources, sources...)
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
			strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			source, err := self.parseJarSource(path)
			if err != nil {
				return err
			}
			self.userSources = append(self.userSources, source)
		}
		source, err := self.parseDirSource(path)
		if err != nil {
			return err
		}
		self.userSources = append(self.userSources, source)
	}
	return
}

func (self *Classpath) parseClassSource(path string) (source JavaSource, err error) {
	_, err = filepath.Abs(path)
	if err != nil {
		return
	}
	source = &JavaSourceClass{
		JavaSourceInfo: &JavaSourceInfo{path},
	}
	return
}

func (self *Classpath) parseJarSource(path string) (source JavaSource, err error) {
	_, err = filepath.Abs(path)
	if err != nil {
		return
	}
	source = &JavaSourceJar{
		JavaSourceInfo: &JavaSourceInfo{path},
	}
	return
}

func (self *Classpath) parseDirSource(path string) (source JavaSource, err error) {
	_, err = filepath.Abs(path)
	if err != nil {
		return
	}
	source = &JavaSourceDir{
		&JavaSourceInfo{path},
	}
	return
}

func (self *Classpath) parseWildcardSource(path string) (sources []JavaSource, err error) {
	baseDir := path[:len(path)-1]
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		return
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != absPath {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			sources = append(sources, &JavaSourceJar{
				&JavaSourceInfo{filepath.Join(baseDir, info.Name())},
			})
		}
		return nil
	}
	filepath.Walk(absPath, walkFn)
	return
}

func (self *Classpath) getJreDir() error {
	if self.JreLibPath != "" && self.exists(self.JreLibPath) {
		return nil
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		self.JreLibPath = filepath.Join(jh, "jre")
		return nil
	}
	return errors.New("Can not find jre folder!")
}

func (self *Classpath) exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
