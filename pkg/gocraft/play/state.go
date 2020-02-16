package play

import (
	"GoCraft/pkg/gocraft/packets"
)

type State struct {
	sessions []packets.Session
}

func NewState() *State {
	return &State{sessions: []packets.Session{}}
}
