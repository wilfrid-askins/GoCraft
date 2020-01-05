package main

import (
	"GoCraft/net/packets"
	"GoCraft/net/packets/client"
	"GoCraft/net/types"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	varIntMax = 4
	varIntValue = 0b0111_1111
	varIntNextFlag = 0b1000_0000
)

type Handler func(packets.Packet)

var handlers = map[uint32]map[types.VarInt]Handler {
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

func main() {
	fmt.Println("Starting...")

	// start server
	listener, err := net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		log.Fatal(err)
	}

	// listen
	for {
		// accept connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	input := bufio.NewReader(conn)
	defer conn.Close()

	state := uint32(packets.HANDSHAKE)

	for {
		lenVal, err := types.VarIntDefault.Read(input)
		length := lenVal.(types.VarInt)

		if err == io.EOF {
			continue
		}

		buf := make([]byte, length)
		_, err = io.ReadFull(input, buf)
		if err != nil {
			fmt.Println(err)
		}

		payload := bufio.NewReader(bytes.NewReader(buf))
		packetType, err := types.VarIntDefault.Read(payload)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Recieved packet %d in state %d\n", packetType, state)

		packetID := packetType.(types.VarInt)
		packet := packets.StateToPacketLookup[state][packetID]

		fmt.Println("Reading payload")
		err = packet.Read(payload)

		if err != nil {
			fmt.Println(err)
		}

		handler := handlers[state][packetID]
		handler(packet)

		if state == packets.HANDSHAKE && packetID == 0 {
			hs := packet.(*client.Handshake)
			state = uint32(hs.NextState)
			fmt.Println("Handshake state change")
		}

		if state == packets.STATUS {
			fmt.Println("Status packet recieved")
		}

		//// Send response
		//conn.Write([]byte("hello"))
	}
}
