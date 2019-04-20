package classfile

import (
	"fmt"

	"github.com/PandaBlame/Gava/java"
	"github.com/PandaBlame/Gava/java/jdk8"
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
func readConstantInfo(reader *java.ClassReader, cp java.ConstantPool) (c java.ConstantInfo, err error) {
	tag, err := reader.ReadUint8()
	if err != nil {
		return
	}
	c = newConstantInfo(tag, cp)
	if c == nil {
		err = fmt.Errorf("java.lang.ClassFormatError: constant pool tag %v", tag)
		return
	}
	err = c.Read(reader)
	return
}

func newConstantInfo(tag uint8, cp java.ConstantPool) java.ConstantInfo {
	switch tag {
	case jdk8.JVM_CONSTANT_Integer:
		return &ConstantIntegerInfoV8{}
	case jdk8.JVM_CONSTANT_Float:
		return &ConstantFloatInfoV8{}
	case jdk8.JVM_CONSTANT_Long:
		return &ConstantLongInfoV8{}
	case jdk8.JVM_CONSTANT_Double:
		return &ConstantDoubleInfoV8{}
	case jdk8.JVM_CONSTANT_Utf8:
		return &ConstantUtf8InfoV8{}
	case jdk8.JVM_CONSTANT_String:
		return &ConstantStringInfoV8{cp: cp}
	case jdk8.JVM_CONSTANT_Class:
		return &ConstantClassInfoV8{cp: cp}
	case jdk8.JVM_CONSTANT_Fieldref:
		return &ConstantFieldRefInfoV8{ConstantMemberRefInfoV8{cp: cp}}
	case jdk8.JVM_CONSTANT_Methodref:
		return &ConstantMethodRefInfoV8{ConstantMemberRefInfoV8{cp: cp}}
	case jdk8.JVM_CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfoV8{ConstantMemberRefInfoV8{cp: cp}}
	case jdk8.JVM_CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfoV8{}
	case jdk8.JVM_CONSTANT_MethodType:
		return &ConstantMethodTypeInfoV8{}
	case jdk8.JVM_CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfoV8{}
	case jdk8.JVM_CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfoV8{}
	default:
		return nil
	}
}
