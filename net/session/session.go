package session

import (
	"GoCraft/net/packets"
	"GoCraft/net/packets/client"
	"GoCraft/net/packets/server"
	"GoCraft/net/types"
	"bufio"
	"bytes"
	"fmt"
)

type Handler func(packets.Packet)

var handlers = map[types.VarInt]map[types.VarInt]Handler {
	packets.HANDSHAKE: {
		0: func(packet packets.Packet) {
			p := packet.(*client.Handshake)
			fmt.Printf("Handshake %d\n", p.ProtocolVersion)
		},
	},
	packets.STATUS: {
		0: func(packet packets.Packet) {
			// p := packet.(*client.Request)
			fmt.Printf("Request (no fields)\n")
		},
		1: func(packet packets.Packet) {
			// TODO
			fmt.Printf("TODO PACKET state 1 0x01 recieved")
		},
	},
	packets.LOGIN: {

	},
	packets.PLAY: {
		3: func(packet packets.Packet) {
			p := packet.(*client.ChatMessage)
			fmt.Printf("ChatMessage %s\n", p.Message)
		},
	},
}

type Session struct {
	state types.VarInt
	out *bufio.Writer
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
		fmt.Println("Status packet received")
		// Write response
		response := server.Response{}
		response.JsonResponse = `{"version":{"name":"1.15.1","protocol":575},"players":{"max":100,"online":5,"sample":[{"name":"thinkofdeath","id":"4566e69f-c907-48ee-8d71-d7ba5aa00d20"}]},"description":{"text":"Hello world!"}}`

		pBuffer := new(bytes.Buffer)
		payload := bufio.NewWriter(pBuffer)
		pType := types.VarInt(0)
		err := pType.Write(payload)
		if err != nil {
			fmt.Println(err)
		}

		err = response.Write(payload)
		if err != nil {
			fmt.Println(err)
		}

		err = payload.Flush()
		if err != nil {
			fmt.Println(err)
		}

		length := types.VarInt(pBuffer.Len())
		err = length.Write(s.out)
		if err != nil {
			fmt.Println(err)
		}

		_, err = s.out.Write(pBuffer.Bytes())
		if err != nil {
			fmt.Println(err)
		}

		err = s.out.Flush()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Response sent")
	}
}

func (s *Session) GetState() types.VarInt {
	return s.state
}

func (s *Session) SetOutput(out *bufio.Writer) {
	s.out = out
}

func NewSession() Session {
	return Session{ types.VarInt(packets.HANDSHAKE), nil}
}
