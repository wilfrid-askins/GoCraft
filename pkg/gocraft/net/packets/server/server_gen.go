package server

import (
	"GoCraft/pkg/gocraft/net/types"
	"bufio"
)

func (p *Response) Read(in *bufio.Reader) error {

	valJsonResponse, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.JsonResponse = valJsonResponse.(types.CraftString)
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

	valPayload, err := types.CraftLongDefault.Read(in)
	if err != nil {
		return err
	}
	p.Payload = valPayload.(types.CraftLong)
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

func (p *LoginSuccess) Read(in *bufio.Reader) error {

	valUUID, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.UUID = valUUID.(types.CraftString)

	valUsername, err := types.CraftStringDefault.Read(in)
	if err != nil {
		return err
	}
	p.Username = valUsername.(types.CraftString)
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
