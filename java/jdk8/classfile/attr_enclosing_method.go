package classfile

import "github.com/PandaBlame/Gava/java"

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttributeV8 struct {
	cp          java.ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (self *EnclosingMethodAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.classIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.methodIndex, err = reader.ReadUint16()
	return
}

func (self *EnclosingMethodAttributeV8) ClassName() (string, error) {
	return self.cp.GetClassName(self.classIndex)
}

func (self *EnclosingMethodAttributeV8) MethodNameAndDescriptor() (string, string, error) {
	if self.methodIndex > 0 {
		return self.cp.GetNameAndType(self.methodIndex)
	} else {
		return "", "", nil
	}
}
