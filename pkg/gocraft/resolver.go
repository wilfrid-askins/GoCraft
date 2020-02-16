package gocraft

import (
	"GoCraft/pkg/gocraft/server"
	"GoCraft/pkg/gocraft/server/play"
	"go.uber.org/zap"
)

type Resolver struct {
	config    server.Config
	playState *play.State
	server    *server.Server
	logger    *zap.Logger
}

func NewResolver(config server.Config, logger *zap.Logger) Resolver {
	return Resolver{config: config, logger: logger}
}

func (r *Resolver) Config() server.Config {
	return r.config
}

func (r *Resolver) PlayState() *play.State {
	if r.playState == nil {
		r.playState = play.NewState()
	}
	return r.playState
}

func (r *Resolver) Server() *server.Server {
	if r.server == nil {
		r.server = server.NewServer(r.Config(), r.PlayState())
	}
	return r.server
}
