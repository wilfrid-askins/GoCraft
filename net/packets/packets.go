package packets

import (
	"GoCraft/net/packets/client"
	"GoCraft/net/types"
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
		Write() error
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
		},
		LOGIN: {

		},
		PLAY: {
			&client.ChatMessage{},
		},
	}
	StateToPacketLookup map[types.VarInt]map[types.VarInt]Packet
)

func init() {

	StateToPacketLookup = make(map[types.VarInt]map[types.VarInt]Packet)

	for state,ps := range StateToPackets {
		StateToPacketLookup[state] = make(map[types.VarInt]Packet)
		for _, p := range ps {
			StateToPacketLookup[state][p.GetID()] = p
		}
	}
}
