package server

import (
	"GoCraft/net/types"
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
