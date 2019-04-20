package classfile

import "github.com/PandaBlame/Gava/java"

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/
type InnerClassesAttributeV8 struct {
	classes []*InnerClassInfoV8
}

type InnerClassInfoV8 struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (self *InnerClassesAttributeV8) Read(reader *java.ClassReader) (err error) {
	numberOfClasses, err := reader.ReadUint16()
	if err != nil {
		return
	}
	self.classes = make([]*InnerClassInfoV8, numberOfClasses)
	for i := range self.classes {
		innerClassInfoIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		outerClassInfoIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		innerNameIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		innerClassAccessFlags, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		self.classes[i] = &InnerClassInfoV8{
			innerClassInfoIndex:   innerClassInfoIndex,
			outerClassInfoIndex:   outerClassInfoIndex,
			innerNameIndex:        innerNameIndex,
			innerClassAccessFlags: innerClassAccessFlags,
		}
	}
	return
}
