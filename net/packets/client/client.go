package client

import (
	"GoCraft/net/types"
)

type (
	// State Handshake
	Handshake struct {
		ID              int               `packet:"0x0"`
		ProtocolVersion types.VarInt
		ServerAddress   types.CraftString
		ServerPort      types.CraftShort
		NextState       types.VarInt
	}

	// State Status
	Request struct {
		ID int `packet:"0x0"`
	}

	Ping struct {
		ID int `packet:"0x01"`
		Payload types.CraftLong
	}

	// State Login
	LoginStart struct {
		ID int `packet:"0x0"`
		Name types.CraftString
	}

	EncryptionResponse struct {
		ID int `packet:"0x1"`

	}

	// State Play
	ChatMessage struct {
		ID int `packet:"0x03"`
		Message types.CraftString
	}
)


