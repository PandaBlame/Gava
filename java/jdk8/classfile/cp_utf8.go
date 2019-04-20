package classfile

import "fmt"
import (
	"unicode/utf16"

	"github.com/PandaBlame/Gava/java"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8InfoV8 struct {
	str string
}

func (self *ConstantUtf8InfoV8) Read(reader *java.ClassReader) error {
	length, err := reader.ReadUint16()
	if err != nil {
		return err
	}
	bytes, err := reader.ReadBytes(uint32(length))
	if err != nil {
		return err
	}
	self.str, err = decodeMUTF8(bytes)
	return err
}

func (self *ConstantUtf8InfoV8) String() string {
	return self.str
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytes []byte) (string, error) {
	utflen := len(bytes)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			count += 2
			if count > utflen {
				return "", fmt.Errorf("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				return "", fmt.Errorf("malformed input around byte %v", count)
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			count += 3
			if count > utflen {
				return "", fmt.Errorf("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				return "", fmt.Errorf("malformed input around byte %v", (count - 1))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			return "", fmt.Errorf("malformed input around byte %v", count)
		}
	}
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes), nil
}
