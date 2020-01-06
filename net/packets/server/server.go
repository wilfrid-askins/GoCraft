package server

import "GoCraft/net/types"

type (
	Response struct {
		ID int `packet:"0x0"`
		JsonResponse types.CraftString
	}

	Pong struct {
		ID int `packet:"0x01"`
		Payload types.CraftLong
	}
)

var (
	ResponseDefault = Response{}
)
