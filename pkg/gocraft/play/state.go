package play

import (
	"GoCraft/pkg/gocraft/net"
)

type State struct {
	sessions []net.Session
}

func NewState() *State {
	return &State{sessions: []net.Session{}}
}
