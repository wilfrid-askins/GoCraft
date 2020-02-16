package packets

import (
	"GoCraft/pkg/gocraft/packets/client"
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
)

const (
	HANDSHAKE = iota
	STATUS
	LOGIN
	PLAY
)

type (
	Packet interface {
		Read(*bufio.Reader) error
		Write(*bufio.Writer) error
		GetID() types.VarInt
	}
)

var (
	StateToPackets = map[types.VarInt][]Packet{
		HANDSHAKE: {
			&client.Handshake{},
		},
		STATUS: {
			&client.Request{},
			&client.Ping{},
		},
		LOGIN: {
			&client.LoginStart{},
			&client.EncryptionResponse{},
		},
		PLAY: {
			&client.ChatMessage{},
		},
	}
	StateToPacketLookup map[types.VarInt]map[types.VarInt]Packet
)

func init() {
	StateToPacketLookup = make(map[types.VarInt]map[types.VarInt]Packet)

	for state, ps := range StateToPackets {
		StateToPacketLookup[state] = make(map[types.VarInt]Packet)
		for _, p := range ps {
			StateToPacketLookup[state][p.GetID()] = p
		}
	}
}
