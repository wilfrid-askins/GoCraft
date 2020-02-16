package server

import (
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
)

func (p *Response) Read(in *bufio.Reader) error {

	valJsonResponse, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.JsonResponse = valJsonResponse.(types.String)
	return nil
}

func (p *Response) Write(out *bufio.Writer) error {
	var err error

	err = p.JsonResponse.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *Response) GetID() types.VarInt {
	return 0
}

func (p *Pong) Read(in *bufio.Reader) error {

	valPayload, err := types.LongDefault.Read(in)
	if err != nil {
		return err
	}
	p.Payload = valPayload.(types.Long)
	return nil
}

func (p *Pong) Write(out *bufio.Writer) error {
	var err error

	err = p.Payload.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pong) GetID() types.VarInt {
	return 1
}

func (p *EncryptionRequest) Read(in *bufio.Reader) error {

	valServerID, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.ServerID = valServerID.(types.String)

	valPublicKeyLength, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.PublicKeyLength = valPublicKeyLength.(types.VarInt)

	valPublicKey, err := types.BytesDefault.Read(in)
	if err != nil {
		return err
	}
	p.PublicKey = valPublicKey.(types.Bytes)

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

func (p *EncryptionRequest) Write(out *bufio.Writer) error {
	var err error

	err = p.ServerID.Write(out)
	if err != nil {
		return err
	}

	err = p.PublicKeyLength.Write(out)
	if err != nil {
		return err
	}

	err = p.PublicKey.Write(out)
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

func (p *EncryptionRequest) GetID() types.VarInt {
	return 1
}

func (p *SetCompression) Read(in *bufio.Reader) error {

	valThreshold, err := types.VarIntDefault.Read(in)
	if err != nil {
		return err
	}
	p.Threshold = valThreshold.(types.VarInt)
	return nil
}

func (p *SetCompression) Write(out *bufio.Writer) error {
	var err error

	err = p.Threshold.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *SetCompression) GetID() types.VarInt {
	return 3
}

func (p *LoginSuccess) Read(in *bufio.Reader) error {

	valUUID, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.UUID = valUUID.(types.String)

	valUsername, err := types.StringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Username = valUsername.(types.String)
	return nil
}

func (p *LoginSuccess) Write(out *bufio.Writer) error {
	var err error

	err = p.UUID.Write(out)
	if err != nil {
		return err
	}

	err = p.Username.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (p *LoginSuccess) GetID() types.VarInt {
	return 2
}
