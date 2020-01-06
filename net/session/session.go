package session

import (
	"GoCraft/net/packets"
	"GoCraft/net/packets/client"
	"GoCraft/net/packets/server"
	"GoCraft/net/types"
	"bufio"
	"fmt"
)

type Handler func(packets.Packet)

var handlers = map[types.VarInt]map[types.VarInt]Handler{}

type Session struct {
	state types.VarInt
	out   *bufio.Writer
}

func (s *Session) OnPacket(p packets.Packet) {
	handler := handlers[s.state][p.GetID()]
	handler(p)
}

func (s *Session) GetState() types.VarInt {
	return s.state
}

func (s *Session) SetOutput(out *bufio.Writer) {
	s.out = out

	handlers[packets.HANDSHAKE] = map[types.VarInt]Handler{
		0x0: func(p packets.Packet) {
			hs := p.(*client.Handshake)
			s.state = hs.NextState
			fmt.Println("Handshake state change")
		},
	}

	handlers[packets.STATUS] = map[types.VarInt]Handler{
		0x0: func(p packets.Packet) {
			fmt.Println("Status packet received")
			// Write response
			response := &server.Response{}
			response.JsonResponse = `{"version":{"name":"1.15.1","protocol":575},"players":{"max":100,"online":5,"sample":[{"name":"thinkofdeath","id":"4566e69f-c907-48ee-8d71-d7ba5aa00d20"}]},"description":{"text":"Good news everyone!"}}`

			err := packets.Write(s.out, response)
			if err != nil {
				fmt.Println(err)
			}
		},
		0x1: func(p packets.Packet) {
			ping := p.(*client.Ping)
			pong := &server.Pong{Payload: ping.Payload}
			err := packets.Write(s.out, pong)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Recieved Ping, sent Pong")
		},
	}

	handlers[packets.LOGIN] = map[types.VarInt]Handler{

	}

	handlers[packets.PLAY] = map[types.VarInt]Handler{
		0x3: func(packet packets.Packet) {
			p := packet.(*client.ChatMessage)
			fmt.Printf("ChatMessage %s\n", p.Message)
		},
	}
}

func NewSession() Session {
	return Session{types.VarInt(packets.HANDSHAKE), nil}
}
