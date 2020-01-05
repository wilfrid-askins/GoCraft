package packets

import (
	"GoCraft/net/types"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

type (
	Handler struct {
		receiver Receiver
	}

	Receiver interface {
		OnPacket(Packet)
	}
)

func (h *Handler) Listen(conn net.Conn) {
	input := bufio.NewReader(conn)
	defer conn.Close()

	state := uint32(HANDSHAKE)

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
		packet := StateToPacketLookup[state][packetID]

		fmt.Println("Reading payload")
		err = packet.Read(payload)

		if err != nil {
			fmt.Println(err)
		}

		h.receiver.OnPacket(packet)

		//// Send response
		//conn.Write([]byte("hello"))
	}
}

func NewHandler(receiver Receiver) Handler {
	return Handler{receiver}
}