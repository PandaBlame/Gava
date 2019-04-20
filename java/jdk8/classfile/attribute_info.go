package classfile

import "github.com/PandaBlame/Gava/java"

var (
	attrDeprecated = &DeprecatedAttributeV8{}
	attrSynthetic  = &SyntheticAttributeV8{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

func readAttributes(reader *java.ClassReader, cp java.ConstantPool) (attributes []java.AttributeInfo, err error) {
	attributesCount, err := reader.ReadUint16()
	if err != nil {
		return
	}
	attributes = make([]java.AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i], err = readAttribute(reader, cp)
		if err != nil {
			return nil, err
		}
	}
	return
}

func readAttribute(reader *java.ClassReader, cp java.ConstantPool) (attrInfo java.AttributeInfo, err error) {
	attrNameIndex, err := reader.ReadUint16()
	if err != nil {
		return
	}
	attrName, err := cp.GetUtf8(attrNameIndex)
	if err != nil {
		return
	}
	attrLen, err := reader.ReadUint32()
	if err != nil {
		return
	}
	attrInfo = newAttributeInfo(attrName, attrLen, cp)
	err = attrInfo.Read(reader)
	return
}

func newAttributeInfo(attrName string, attrLen uint32, cp java.ConstantPool) java.AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttributeV8{}
	case "Code":
		return &CodeAttributeV8{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttributeV8{}
	case "Deprecated":
		return attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttributeV8{cp: cp}
	case "Exceptions":
		return &ExceptionsAttributeV8{}
	case "InnerClasses":
		return &InnerClassesAttributeV8{}
	case "LineNumberTable":
		return &LineNumberTableAttributeV8{}
	case "LocalVariableTable":
		return &LocalVariableTableAttributeV8{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttributeV8{}
	// case "MethodParameters":
	// case "RuntimeInvisibleAnnotations":
	// case "RuntimeInvisibleParameterAnnotations":
	// case "RuntimeInvisibleTypeAnnotations":
	// case "RuntimeVisibleAnnotations":
	// case "RuntimeVisibleParameterAnnotations":
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttributeV8{cp: cp}
	case "SourceFile":
		return &SourceFileAttributeV8{cp: cp}
	// case "SourceDebugExtension":
	// case "StackMapTable":
	case "Synthetic":
		return attrSynthetic
	default:
		return &UnParsedAttributeV8{attrName, attrLen, nil}
	}
}
