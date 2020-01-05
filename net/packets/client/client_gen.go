package client

import (
	"GoCraft/net/types"
	"bufio"
)

func (t *Handshake) Read(in *bufio.Reader) error {

	valProtocolVersion, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	t.ProtocolVersion = valProtocolVersion.(types.VarInt)

	valServerAddress, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	t.ServerAddress = valServerAddress.(types.CraftString)

	valServerPort, err := types.CraftShortDefault.Read(in)
	if err != nil {
		return err
	}
	t.ServerPort = valServerPort.(types.CraftShort)

	valNextState, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	t.NextState = valNextState.(types.VarInt)
	return nil
}

func (t *Handshake) Write() error {
	return nil
}

func (t *Handshake) GetID() types.VarInt {
	return 0
}

func (t *Request) Read(in *bufio.Reader) error {

	return nil
}

func (t *Request) Write() error {
	return nil
}

func (t *Request) GetID() types.VarInt {
	return 0
}

func (t *ChatMessage) Read(in *bufio.Reader) error {

	valMessage, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	t.Message = valMessage.(types.CraftString)
	return nil
}

func (t *ChatMessage) Write() error {
	return nil
}

func (t *ChatMessage) GetID() types.VarInt {
	return 3
}
