package internal

import (
	"GoCraft/pkg/gocraft"
	"GoCraft/pkg/gocraft/play"
	"go.uber.org/zap"
)

type Resolver struct {
	config    gocraft.Config
	playState *play.State
	server    *gocraft.Server
	logger    *zap.SugaredLogger
}

func NewResolver(config gocraft.Config, logger *zap.SugaredLogger) Resolver {
	return Resolver{config: config, logger: logger}
}

func (r *Resolver) Config() gocraft.Config {
	return r.config
}

func (r *Resolver) Logger() *zap.SugaredLogger {
	return r.logger
}

func (r *Resolver) PlayState() *play.State {
	if r.playState == nil {
		r.playState = play.NewState()
	}
	return r.playState
}

func (r *Resolver) Server() *gocraft.Server {
	if r.server == nil {
		r.server = gocraft.NewServer(r.Config(), r.PlayState(), r.Logger())
	}
	return r.server
}
