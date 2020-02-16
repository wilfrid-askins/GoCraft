package client

import (
	"GoCraft/pkg/gocraft/net/types"
	"bufio"
)

func (p *Handshake) Read(in *bufio.Reader) error {

	valProtocolVersion, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.ProtocolVersion = valProtocolVersion.(types.VarInt)

	valServerAddress, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.ServerAddress = valServerAddress.(types.CraftString)

	valServerPort, err := types.CraftShortDefault.Read(in)
	if err != nil {
		return err
	}
	p.ServerPort = valServerPort.(types.CraftShort)

	valNextState, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.NextState = valNextState.(types.VarInt)
	return nil
}

func (p *Handshake) Write(out *bufio.Writer) error {
	var err error

	err = p.ProtocolVersion.Write(out)
	if err != nil {
		return err
	}

	err = p.ServerAddress.Write(out)
	if err != nil {
		return err
	}

	err = p.ServerPort.Write(out)
	if err != nil {
		return err
	}

	err = p.NextState.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *Handshake) GetID() types.VarInt {
	return 0
}

func (p *Request) Read(in *bufio.Reader) error {

	return nil
}

func (p *Request) Write(out *bufio.Writer) error {

	return nil
}

func (p *Request) GetID() types.VarInt {
	return 0
}

func (p *ChatMessage) Read(in *bufio.Reader) error {

	valMessage, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Message = valMessage.(types.CraftString)
	return nil
}

func (p *ChatMessage) Write(out *bufio.Writer) error {
	var err error

	err = p.Message.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *ChatMessage) GetID() types.VarInt {
	return 3
}

func (p *Ping) Read(in *bufio.Reader) error {

	valPayload, err := types.CraftLongDefault.Read(in)
	if err != nil {
		return err
	}
	p.Payload = valPayload.(types.CraftLong)
	return nil
}

func (p *Ping) Write(out *bufio.Writer) error {
	var err error

	err = p.Payload.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *Ping) GetID() types.VarInt {
	return 1
}

func (p *LoginStart) Read(in *bufio.Reader) error {

	valName, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Name = valName.(types.CraftString)
	return nil
}

func (p *LoginStart) Write(out *bufio.Writer) error {
	var err error

	err = p.Name.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *LoginStart) GetID() types.VarInt {
	return 0
}
