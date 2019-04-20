package classfile

import "github.com/PandaBlame/Gava/java"

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttributeV8 struct {
	lineNumberTable []*LineNumberTableEntryV8
}

type LineNumberTableEntryV8 struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttributeV8) Read(reader *java.ClassReader) (err error) {
	lineNumberTableLength, err := reader.ReadUint16()
	if err != nil {
		return
	}
	self.lineNumberTable = make([]*LineNumberTableEntryV8, lineNumberTableLength)
	for i := range self.lineNumberTable {
		startPc, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		lineNumber, err := reader.ReadUint16()
		if err != nil {
			return err
		}
		self.lineNumberTable[i] = &LineNumberTableEntryV8{
			startPc:    startPc,
			lineNumber: lineNumber,
		}
	}
	return
}

func (self *LineNumberTableAttributeV8) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- {
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
