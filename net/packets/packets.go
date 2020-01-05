package packets

import "bufio"

const (
	STATUS = iota + 1
	LOGIN
	PLAY
)

type (
	Packet interface {
		Read(bufio.Reader) error
		Write()
	}
)

var (
	statePackets = map[int]Packet{

	}
)