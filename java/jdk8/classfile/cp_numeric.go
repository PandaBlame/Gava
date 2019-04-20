package classfile

import (
	"math"

	"github.com/PandaBlame/Gava/java"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfoV8 struct {
	val int32
}

func (self *ConstantIntegerInfoV8) Read(reader *java.ClassReader) (err error) {
	bytes, err := reader.ReadUint32()
	if err != nil {
		return
	}
	self.val = int32(bytes)
	return
}
func (self *ConstantIntegerInfoV8) Value() int32 {
	return self.val
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfoV8 struct {
	val float32
}

func (self *ConstantFloatInfoV8) Read(reader *java.ClassReader) (err error) {
	bytes, err := reader.ReadUint32()
	if err != nil {
		return
	}
	self.val = math.Float32frombits(bytes)
	return
}
func (self *ConstantFloatInfoV8) Value() float32 {
	return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfoV8 struct {
	val int64
}

func (self *ConstantLongInfoV8) Read(reader *java.ClassReader) (err error) {
	bytes, err := reader.ReadUint64()
	if err != nil {
		return
	}
	self.val = int64(bytes)
	return
}
func (self *ConstantLongInfoV8) Value() int64 {
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfoV8 struct {
	val float64
}

func (self *ConstantDoubleInfoV8) Read(reader *java.ClassReader) (err error) {
	bytes, err := reader.ReadUint64()
	if err != nil {
		return
	}
	self.val = math.Float64frombits(bytes)
	return
}
func (self *ConstantDoubleInfoV8) Value() float64 {
	return self.val
}
