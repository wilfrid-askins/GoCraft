package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
		// accept connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	input := bufio.NewReader(conn)
	defer conn.Close()

	state := uint32(0)

	for {
		length, err := readVarInt(input)

		if err == io.EOF {
			continue
		}

		fmt.Println(length)

		buf := make([]byte, length)
		_, err = io.ReadFull(input, buf)
		if err != nil {
			fmt.Println(err)
		}

		payload := bufio.NewReader(bytes.NewReader(buf))
		packetType, err := readVarInt(payload)
		if err != nil {
			fmt.Println(err)
		}

		if packetType == 0 {

			if state == 0 {
				fmt.Println("Processing handshake")

				protocolVersion, err := readVarInt(payload)
				if err != nil {
					fmt.Println(err)
				}

				serverAddr, err := readString(payload)
				if err != nil {
					fmt.Println(err)
				}

				serverPort, err := readShort(payload)
				if err != nil {
					fmt.Println(err)
				}

				nextState, err := readVarInt(payload)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(protocolVersion)
				fmt.Println(serverAddr)
				fmt.Println(serverPort)
				fmt.Println(nextState)

				state = nextState

			} else {
				fmt.Println("Processing Request Packet")
			}
		}

		//// Send response
		//conn.Write([]byte("hello"))
	}
}

func readShort(input *bufio.Reader) (int64, error) {
	val1, err := input.ReadByte()
	if err != nil {
		return 0, err
	}

	val2, err := input.ReadByte()
	if err != nil {
		return 0, err
	}

	num := int64(val2) | int64(val1) << 8

	return num, nil
}

func readString(input *bufio.Reader) (string, error) {
	length, err := readVarInt(input)

	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	io.ReadFull(input, strBytes)

	return string(strBytes), nil
}

func readVarInt(input *bufio.Reader) (uint32, error) {
	num := uint32(0)

	for cur := uint32(0); cur <= varIntMax; cur++ {

		part, err := input.ReadByte()
		if err != nil {
			return 0, err
		}

		num |= uint32(part & varIntValue) << (7 * cur)
		if part & varIntNextFlag == 0 {
			break
		}
	}

	return num, nil
}