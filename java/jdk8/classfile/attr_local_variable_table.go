package classfile

import "github.com/PandaBlame/Gava/java"

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttributeV8 struct {
	localVariableTable []*LocalVariableTableEntryV8
}

type LocalVariableTableEntryV8 struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (self *LocalVariableTableAttributeV8) Read(reader *java.ClassReader) (err error) {
	localVariableTableLength, err := reader.ReadUint16()
	if err != nil {
		return
	}
	self.localVariableTable = make([]*LocalVariableTableEntryV8, localVariableTableLength)
	for i := range self.localVariableTable {
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
		descriptorIndex, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		index, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		self.localVariableTable[i] = &LocalVariableTableEntryV8{
			startPc:         startPc,
			length:          length,
			nameIndex:       nameIndex,
			descriptorIndex: descriptorIndex,
			index:           index,
		}
	}
	return
}
