package play

import (
	"GoCraft/pkg/gocraft/server/net/session"
)

type State struct {
	sessions []session.Session
}

func NewState() *State {
	return &State{sessions: []session.Session{}}
}
