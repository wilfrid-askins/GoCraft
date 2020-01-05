package session

import (
	"GoCraft/net/packets"
	"GoCraft/net/packets/client"
	"GoCraft/net/types"
	"fmt"
)

type Handler func(packets.Packet)

var handlers = map[types.VarInt]map[types.VarInt]Handler {
	packets.HANDSHAKE: {
		0: func(packet packets.Packet) {
			p := packet.(*client.Handshake)
			fmt.Printf("Handshake %s %d\n", p.ServerAddress, p.ServerPort)
		},
	},
	packets.STATUS: {
		// Request
		0: func(packet packets.Packet) {
			// p := packet.(*client.Request)
			fmt.Printf("Request (no fields)\n")
		},
	},
	packets.LOGIN: {

	},
	packets.PLAY: {
		// Chat message
		3: func(packet packets.Packet) {
			p := packet.(*client.ChatMessage)
			fmt.Printf("ChatMessage %s\n", p.Message)
		},
	},
}

type Session struct {
	state types.VarInt
}

func (s *Session) OnPacket(p packets.Packet) {
	handler := handlers[s.state][p.GetID()]
	handler(p)

	if s.state == packets.HANDSHAKE && p.GetID() == 0 {
		hs := p.(*client.Handshake)
		s.state = hs.NextState
		fmt.Println("Handshake state change")
	}

	if s.state == packets.STATUS {
		fmt.Println("Status packet recieved")
	}
}

func (s *Session) GetState() types.VarInt {
	return s.state
}

func NewSession() Session {
	return Session{ types.VarInt(packets.HANDSHAKE)}
}
