package main

import (
	"GoCraft/net/packets"
	"GoCraft/net/session"
	"fmt"
	"log"
	"net"
)

const (
	varIntMax = 4
	varIntValue = 0b0111_1111
	varIntNextFlag = 0b1000_0000
)

func main() {
	fmt.Println("Starting...")

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

		sess := session.NewSession()
		handler := packets.NewHandler(&sess)
		go handler.Listen(conn)
	}
}
