package session

import (
	packets2 "GoCraft/pkg/gocraft/server/net/packets"
	client2 "GoCraft/pkg/gocraft/server/net/packets/client"
	server2 "GoCraft/pkg/gocraft/server/net/packets/server"
	types2 "GoCraft/pkg/gocraft/server/net/types"
	"bufio"
	"fmt"
)

type Handler func(packets2.Packet)

var handlers = map[types2.VarInt]map[types2.VarInt]Handler{}

type Session struct {
	state      types2.VarInt
	out        *bufio.Writer
	summariser Summariser
}

type Summariser interface {
	GetSummary() string
}

func NewSession(summariser Summariser) Session {
	return Session{types2.VarInt(packets2.HANDSHAKE), nil, summariser}
}

func (s *Session) OnPacket(p packets2.Packet) {
	handler := handlers[s.state][p.GetID()]
	handler(p)
}

func (s *Session) GetState() types2.VarInt {
	return s.state
}

func (s *Session) SetOutput(out *bufio.Writer) {
	s.out = out

	handlers[packets2.HANDSHAKE] = map[types2.VarInt]Handler{
		0x0: func(p packets2.Packet) {
			hs := p.(*client2.Handshake)
			s.state = hs.NextState
			fmt.Println("Handshake state change")
		},
	}

	handlers[packets2.STATUS] = map[types2.VarInt]Handler{
		0x0: func(p packets2.Packet) {
			fmt.Println("Status packet received")
			summary := s.summariser.GetSummary()

			// Write response
			response := &server2.Response{}
			response.JsonResponse = types2.CraftString(summary)

			err := packets2.Write(s.out, response)
			if err != nil {
				fmt.Println(err)
			}
		},
		0x1: func(p packets2.Packet) {
			ping := p.(*client2.Ping)
			pong := &server2.Pong{Payload: ping.Payload}
			err := packets2.Write(s.out, pong)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Recieved Ping, sent Pong")
			// TODO close connection
		},
	}

	handlers[packets2.LOGIN] = map[types2.VarInt]Handler{
		0x0: func(p packets2.Packet) {
			lp := p.(*client2.LoginStart)
			fmt.Println("Login from " + lp.Name)

			success := &server2.LoginSuccess{
				UUID:     "31574af2-21ff-53hf-5832-94b63e5o6678",
				Username: lp.Name,
			}
			err := packets2.Write(s.out, success)
			if err != nil {
				fmt.Println(err)
			}
			// Login start
			// Optional: Send encryption request
		},
		0x1: func(p packets2.Packet) {
			// Optional: Encryption response
			// Optional: Send set compression
		},
		// Finally: Send login success
	}

	handlers[packets2.PLAY] = map[types2.VarInt]Handler{
		0x3: func(packet packets2.Packet) {
			p := packet.(*client2.ChatMessage)
			fmt.Printf("ChatMessage %s\n", p.Message)
		},
	}
}
