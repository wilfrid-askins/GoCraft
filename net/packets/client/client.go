package client

import (
	"GoCraft/net/types"
)

type (
	Handshake struct {
		ID              int               `packet:"0x0"`
		ProtocolVersion types.VarInt
		ServerAddress   types.CraftString
		ServerPort      types.CraftShort
		NextState       types.VarInt
	}

	Request struct {
		ID int `packet:"0x0"`
	}

	Ping struct {
		ID int `packet:"0x01"`
		Payload types.CraftLong
	}

	ChatMessage struct {
		ID int `packet:"0x03"`
		Message types.CraftString
	}
)


