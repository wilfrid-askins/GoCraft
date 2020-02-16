package gocraft

import (
	"GoCraft/pkg/gocraft/packets"
	"GoCraft/pkg/gocraft/play"
	"encoding/json"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
	"net"
)

type Server struct {
	config Config
	state  *play.State
	logger *zap.SugaredLogger
}

func NewServer(config Config, playState *play.State, logger *zap.SugaredLogger) *Server {
	return &Server{config, playState, logger}
}

func (se *Server) Listen() error {
	// start server
	listener, err := net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		return errors.Wrap(err, "failed to listen on port")
	}

	// listen
	for {
		// accept handler
		conn, err := listener.Accept()
		if err != nil {
			return errors.Wrap(err, "failed to accept listener")
		}

		sess := packets.NewSession(se, se.logger)
		handler := packets.NewHandler(&sess, se.logger)
		go handler.Listen(conn)
	}
}

func (se *Server) GetSummary() string {
	text, err := json.Marshal(se.config.GetSummary(se.state))
	if err != nil {
		se.logger.Fatalw("failed to unmarshal server summary", "msg", errors.Message(err))
	}
	return string(text)
}
