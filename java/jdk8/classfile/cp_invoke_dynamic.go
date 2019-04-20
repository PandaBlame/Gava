package classfile

import "github.com/PandaBlame/Gava/java"

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfoV8 struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfoV8) Read(reader *java.ClassReader) (err error) {
	self.referenceKind, err = reader.ReadUint8()
	if err != nil {
		return
	}
	self.referenceIndex, err = reader.ReadUint16()
	return
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfoV8 struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfoV8) Read(reader *java.ClassReader) (err error) {
	self.descriptorIndex, err = reader.ReadUint16()
	return
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfoV8 struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfoV8) Read(reader *java.ClassReader) (err error) {
	self.bootstrapMethodAttrIndex, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.nameAndTypeIndex, err = reader.ReadUint16()
	return
}
