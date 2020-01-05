package types

import "bufio"

const (
	varIntMax = 4
	varIntValue = 0b0111_1111
	varIntNextFlag = 0b1000_0000

	VarIntDefault = VarInt(0)
	CraftShortDefault = CraftShort(0)
)

type (
	VarInt uint32
	CraftShort int16
)

func (v *VarInt) Read(input *bufio.Reader) (interface{}, error) {
	num := uint32(0)

	for cur := uint32(0); cur <= varIntMax; cur++ {

		part, err := input.ReadByte()
		if err != nil {
			return 0, err
		}

		num |= uint32(part & varIntValue) << (7 * cur)
		if part & varIntNextFlag == 0 {
			break
		}
	}

	return VarInt(num), nil
}

func (v *VarInt) Write() error {
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

func (v *CraftShort) Write() error {
	return nil
}
