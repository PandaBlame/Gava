package classfile

import "github.com/PandaBlame/Gava/java"

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttributeV8 struct {
	MarkerAttributeV8
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttributeV8 struct {
	MarkerAttributeV8
}

type MarkerAttributeV8 struct{}

func (self *MarkerAttributeV8) Read(reader *java.ClassReader) (err error) {
	return
}
