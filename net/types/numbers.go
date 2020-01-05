package types

import (
	"bufio"
	"encoding/binary"
)

const (
	varIntMax = 4
	varIntValue = 0b0111_1111
	varIntNextFlag = 0b1000_0000
	)

type (
	VarInt uint32
	CraftShort int16
)

var (
	VarIntDefault = new(VarInt)
	CraftShortDefault = new(CraftShort)
)

func (v *VarInt) Read(input *bufio.Reader) (interface{}, error) {
	num := uint32(0)

	for cur := uint32(0); cur <= varIntMax; cur++ {

		part, err := input.ReadByte()
		if err != nil {
			return VarInt(0), err
		}

		num |= uint32(part & varIntValue) << (7 * cur)
		if part & varIntNextFlag == 0 {
			break
		}
	}

	return VarInt(num), nil
}

func (v *VarInt) Write(out *bufio.Writer) error {
	value := *v
	for {
		temp := value & varIntValue
		value = value >> 7
		if value != 0 {
			temp |= 0b10000000
		}
		_, err := out.Write([]byte{ byte(temp) })
		//err := binary.Write(out, binary.LittleEndian, temp)

		if err != nil {
			return err
		}

		if value == 0 {
			break
		}
	}

	return nil
}

func (v *CraftShort) Read(input *bufio.Reader) (interface{}, error) {
	val1, err := input.ReadByte()
	if err != nil {
		return 0, err
	}

	val2, err := input.ReadByte()
	if err != nil {
		return 0, err
	}

	// TODO change int16
	num := int64(val2) | int64(val1) << 8

	return CraftShort(num), nil
}

// TODO check this works
func (v *CraftShort) Write(out *bufio.Writer) error {
	return binary.Write(out, binary.LittleEndian, *v)
}
