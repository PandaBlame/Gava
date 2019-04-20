package classfile

import "github.com/PandaBlame/Gava/java"

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttributeV8 struct {
	bootstrapMethods []*BootstrapMethodV8
}

func (self *BootstrapMethodsAttributeV8) Read(reader *java.ClassReader) (err error) {
	numBootstrapMethods, err := reader.ReadUint16()
	self.bootstrapMethods = make([]*BootstrapMethodV8, numBootstrapMethods)
	for i := range self.bootstrapMethods {
		bootstrapMethodRef, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		bootstrapArguments, err := reader.ReadUint16s()
		if err != nil {
			return err
		}
		self.bootstrapMethods[i] = &BootstrapMethodV8{
			bootstrapMethodRef: bootstrapMethodRef,
			bootstrapArguments: bootstrapArguments,
		}
	}
	return
}

type BootstrapMethodV8 struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
