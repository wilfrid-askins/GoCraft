package packets

import (
	"GoCraft/pkg/gocraft/net/packets/client"
	"GoCraft/pkg/gocraft/net/types"
	"bufio"
	"bytes"
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

func Write(out *bufio.Writer, p Packet) error {
	pBuffer := &bytes.Buffer{}
	payload := bufio.NewWriter(pBuffer)
	pType := types.VarInt(0)
	err := pType.Write(payload)
	if err != nil {
		return err
	}

	err = p.Write(payload)
	if err != nil {
		return err
	}

	err = payload.Flush()
	if err != nil {
		return err
	}

	length := types.VarInt(pBuffer.Len())
	err = length.Write(out)
	if err != nil {
		return err
	}

	_, err = out.Write(pBuffer.Bytes())
	if err != nil {
		return err
	}

	err = out.Flush()
	if err != nil {
		return err
	}

	return nil
}
