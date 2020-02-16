package client

import (
	"GoCraft/pkg/gocraft/packets/types"
)

type (
	// State Handshake
	Handshake struct {
		ID              int `packet:"0x0"`
		ProtocolVersion types.VarInt
		ServerAddress   types.String
		ServerPort      types.Short
		NextState       types.VarInt
	}

	// State Status
	Request struct {
		ID int `packet:"0x0"`
	}

	Ping struct {
		ID      int `packet:"0x01"`
		Payload types.Long
	}

	// State Login
	LoginStart struct {
		ID   int `packet:"0x0"`
		Name types.String
	}

	EncryptionResponse struct {
		ID                 int `packet:"0x1"`
		SharedSecretLength types.VarInt
		SharedSecret       types.Bytes
		VerifyTokenLength  types.VarInt
		VerifyToken        types.Bytes
	}

	// State Play
	ChatMessage struct {
		ID      int `packet:"0x03"`
		Message types.String
	}
)
