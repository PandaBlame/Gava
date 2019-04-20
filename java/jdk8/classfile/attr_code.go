package classfile

import "github.com/PandaBlame/Gava/java"

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttributeV8 struct {
	cp             java.ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []java.ExceptionTableEntry
	attributes     []java.AttributeInfo
}

func (self *CodeAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.maxStack, err = reader.ReadUint16()
	if err != nil {
		return
	}
	self.maxLocals, err = reader.ReadUint16()
	if err != nil {
		return
	}
	codeLength, err := reader.ReadUint32()
	if err != nil {
		return
	}
	self.code, err = reader.ReadBytes(codeLength)
	if err != nil {
		return
	}
	self.exceptionTable, err = readExceptionTable(reader)
	if err != nil {
		return
	}
	self.attributes, err = readAttributes(reader, self.cp)
	return
}

func (self *CodeAttributeV8) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttributeV8) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttributeV8) Code() []byte {
	return self.code
}
func (self *CodeAttributeV8) ExceptionTable() []java.ExceptionTableEntry {
	return self.exceptionTable
}

func (self *CodeAttributeV8) LineNumberTableAttribute() java.LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case java.LineNumberTableAttribute:
			return attrInfo.(java.LineNumberTableAttribute)
		}
	}
	return nil
}

type ExceptionTableEntryV8 struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *java.ClassReader) (exceptionTable []java.ExceptionTableEntry, err error) {
	exceptionTableLength, err := reader.ReadUint16()
	exceptionTable = make([]java.ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		entry := &ExceptionTableEntryV8{}
		err := entry.Read(reader)
		if err != nil {
			return nil, err
		}
		exceptionTable[i] = entry
	}
	return
}

func (self *ExceptionTableEntryV8) Read(reader *java.ClassReader) (err error) {
	self.startPc, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	self.endPc, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	self.handlerPc, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	self.catchType, err = reader.ReadUint16()
	return
}

func (self *ExceptionTableEntryV8) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntryV8) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntryV8) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntryV8) CatchType() uint16 {
	return self.catchType
}
