package classfile

import "github.com/PandaBlame/Gava/java"

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTableAttributeV8 struct {
	localVariableTypeTable []*LocalVariableTypeTableEntryV8
}

type LocalVariableTypeTableEntryV8 struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (self *LocalVariableTypeTableAttributeV8) Read(reader *java.ClassReader) (err error) {
	localVariableTypeTableLength, err := reader.ReadUint16()
	if err != nil {
		return
	}
	self.localVariableTypeTable = make([]*LocalVariableTypeTableEntryV8, localVariableTypeTableLength)
	for i := range self.localVariableTypeTable {
		startPc, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		length, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		nameIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		signatureIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		index, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		self.localVariableTypeTable[i] = &LocalVariableTypeTableEntryV8{
			startPc:        startPc,
			length:         length,
			nameIndex:      nameIndex,
			signatureIndex: signatureIndex,
			index:          index,
		}
	}
	return
}
