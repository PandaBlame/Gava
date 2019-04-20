package classfile

import "github.com/PandaBlame/Gava/java"

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttributeV8 struct {
	cp             java.ConstantPool
	signatureIndex uint16
}

func (self *SignatureAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.signatureIndex, err = reader.ReadUint16()
	return
}

func (self *SignatureAttributeV8) Signature() (string, error) {
	return self.cp.GetUtf8(self.signatureIndex)
}
