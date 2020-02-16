package types

import (
	"bufio"
	"io"
)

type (
	String string
	Bytes  []byte
)

var (
	StringDefault = new(String)
	BytesDefault  = new(Bytes)
)

func (cs *String) Read(input *bufio.Reader) (interface{}, error) {
	length, err := VarIntDefault.Read(input)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length.(VarInt))
	_, err = io.ReadFull(input, strBytes)
	if err != nil {
		return nil, err
	}

	return String(string(strBytes)), nil
}

func (cs *String) Write(out *bufio.Writer) error {
	length := VarInt(len([]byte(*cs)))
	err := length.Write(out)
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(*cs))
	if err != nil {
		return err
	}

	return nil
}

func (bs *Bytes) Read(input *bufio.Reader) (interface{}, error) {
	length, err := VarIntDefault.Read(input)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length.(VarInt))
	_, err = io.ReadFull(input, strBytes)
	if err != nil {
		return nil, err
	}

	return String(string(strBytes)), nil
}

func (bs *Bytes) Write(out *bufio.Writer) error {
	_, err := out.Write(*bs)
	if err != nil {
		return err
	}

	return nil
}
