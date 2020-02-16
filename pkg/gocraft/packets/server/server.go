package server

import (
	"GoCraft/pkg/gocraft/packets/types"
)

type (
	// State Handshake
	Response struct {
		ID           int `packet:"0x0"`
		JsonResponse types.String
	}

	// State Status
	Pong struct {
		ID      int `packet:"0x01"`
		Payload types.Long
	}

	// State Login
	EncryptionRequest struct {
		ID int `packet:"0x01"`
		ServerID types.String
		PublicKeyLength types.VarInt
		PublicKey types.Bytes
		VerifyTokenLength types.VarInt
		VerifyToken types.Bytes
	}

	SetCompression struct {
		ID        int `packet:"0x03"`
		Threshold types.VarInt
	}

	LoginSuccess struct {
		ID       int `packet:"0x02"`
		UUID     types.String
		Username types.String
	}

	// State Play

)

var (
	ResponseDefault = Response{}
)
