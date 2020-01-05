package types

import (
	"bufio"
	"io"
)

const (
	CraftStringDefault = CraftString("")
)

type (
	CraftString string
)

func (cs *CraftString) Read(input *bufio.Reader) (interface{}, error) {

	length, err := VarIntDefault.Read(input)

	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length.(VarInt))
	_, err = io.ReadFull(input, strBytes)

	if err != nil {
		return nil, err
	}

	return CraftString(string(strBytes)), nil
}
