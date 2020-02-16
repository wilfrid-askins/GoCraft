package gocraft

import (
	"GoCraft/pkg/gocraft/play"
	"go.uber.org/zap"
)

type Resolver struct {
	config    Config
	playState *play.State
	server    *Server
	logger    *zap.Logger
}

func NewResolver(config Config, logger *zap.Logger) Resolver {
	return Resolver{config: config, logger: logger}
}

func (r *Resolver) Config() Config {
	return r.config
}

func (r *Resolver) PlayState() *play.State {
	if r.playState == nil {
		r.playState = play.NewState()
	}
	return r.playState
}

func (r *Resolver) Server() *Server {
	if r.server == nil {
		r.server = NewServer(r.Config(), r.PlayState())
	}
	return r.server
}
