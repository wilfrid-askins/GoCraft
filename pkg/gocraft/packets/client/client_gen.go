package client

import (
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
)

func (p *Handshake) Read(in *bufio.Reader) error {

	valProtocolVersion, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.ProtocolVersion = valProtocolVersion.(types.VarInt)

	valServerAddress, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.ServerAddress = valServerAddress.(types.String)

	valServerPort, err := types.ShortDefault.Read(in)
	if err != nil {
		return err
	}
	p.ServerPort = valServerPort.(types.Short)

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

	valMessage, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Message = valMessage.(types.String)
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

	valPayload, err := types.LongDefault.Read(in)
	if err != nil {
		return err
	}
	p.Payload = valPayload.(types.Long)
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

	valName, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Name = valName.(types.String)
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

func (p *EncryptionResponse) Read(in *bufio.Reader) error {

	valSharedSecretLength, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.SharedSecretLength = valSharedSecretLength.(types.VarInt)

	valSharedSecret, err := types.BytesDefault.Read(in)
	if err != nil {
		return err
	}
	p.SharedSecret = valSharedSecret.(types.Bytes)

	valVerifyTokenLength, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.VerifyTokenLength = valVerifyTokenLength.(types.VarInt)

	valVerifyToken, err := types.BytesDefault.Read(in)
	if err != nil {
		return err
	}
	p.VerifyToken = valVerifyToken.(types.Bytes)
	return nil
}

func (p *EncryptionResponse) Write(out *bufio.Writer) error {
	var err error

	err = p.SharedSecretLength.Write(out)
	if err != nil {
		return err
	}

	err = p.SharedSecret.Write(out)
	if err != nil {
		return err
	}

	err = p.VerifyTokenLength.Write(out)
	if err != nil {
		return err
	}

	err = p.VerifyToken.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *EncryptionResponse) GetID() types.VarInt {
	return 1
}
