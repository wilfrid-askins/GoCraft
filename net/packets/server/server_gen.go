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

func (t *Response) GetID() types.VarInt {
	return 0
}
