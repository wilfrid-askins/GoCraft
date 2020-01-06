package server

import "GoCraft/net/types"

type (
	// State Handshake
	Response struct {
		ID int `packet:"0x0"`
		JsonResponse types.CraftString
	}

	// State Status
	Pong struct {
		ID int `packet:"0x01"`
		Payload types.CraftLong
	}

	// State Login
	EncryptionRequest struct {
		ID int `packet:"0x01"`

	}

	SetCompression struct {
		ID int `packet:"0x03"`
		Threshold types.VarInt
	}

	LoginSuccess struct {
		ID int `packet:"0x02"`
		UUID types.CraftString
		Username types.CraftString
	}

	// State Play

)

var (
	ResponseDefault = Response{}
)
