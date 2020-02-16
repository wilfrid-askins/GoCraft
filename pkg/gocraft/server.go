package gocraft

import (
	net2 "GoCraft/pkg/gocraft/net"
	"GoCraft/pkg/gocraft/net/packets"
	"GoCraft/pkg/gocraft/play"
	"encoding/json"
	"log"
	"net"
)

type Server struct {
	config Config
	state  *play.State
}

func NewServer(config Config, playState *play.State) *Server {
	return &Server{config, playState}
}

func (se *Server) Listen() {
	// start server
	listener, err := net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		log.Fatal(err)
	}

	// listen
	for {
		// accept handler
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		sess := net2.NewSession(se)
		handler := packets.NewHandler(&sess)
		go handler.Listen(conn)
	}
}

func (se *Server) GetSummary() string {
	text, err := json.Marshal(se.config.GetSummary(se.state))
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}
