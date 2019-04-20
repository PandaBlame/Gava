package classfile

import "github.com/PandaBlame/Gava/java"

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttributeV8 struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttributeV8) Read(reader *java.ClassReader) (err error) {
	self.exceptionIndexTable, err = reader.ReadUint16s()
	return
}

func (self *ExceptionsAttributeV8) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
