package main

import (
	"GoCraft/net/packets/client"
	"GoCraft/net/types"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const (
	Newline = "\n"
	FileHeader = `
	package client

	import (
		"GoCraft/net/types"
		"bufio"
	)`
	ReadTemplate = `
	func (t *%s) Read(in *bufio.Reader) error {
		%s
		return nil
	}`
	ReadSingleTemplate = `
	val%s, err := %s.Read(in)
	if err != nil {
		return err
	}
	t.%s = val%s.(%s)`
	WriteTemplate = `
	func (t *%s) Write() error {
		return nil
	}`
	GetIDTemplate = `
	func (t *%s) GetID() types.VarInt {
		return %d
	}`
	DefaultInstPostfix = "Default"
	TypesPkgSuffix     = "types."
	StructTag = "packet"
	)

//go:generate go run gen.go
//go:generate gofmt -w ./net/packets/client/client_gen.go
func main() {
	fmt.Println("Generating packet code")

	packetSlice := []interface{}{
		client.Handshake{},
		client.Request{},
		client.ChatMessage{},
	}

	code := make([]byte, 0)

	code = append(code, FileHeader...)

	for _, p := range packetSlice {
		sum := getSummary(p)
		readBlock := fmt.Sprintf(ReadTemplate, sum.name, getReadBody(sum))
		writeBlock := getWriteBody(sum)
		idBlock := getIDBody(sum)

		code = append(code, readBlock...)
		code = append(code, Newline...)
		code = append(code, writeBlock...)
		code = append(code, Newline...)
		code = append(code, idBlock...)
		code = append(code, Newline...)
	}

	err := ioutil.WriteFile("./net/packets/client/client_gen.go", code, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

type (
	packetSummary struct {
		name string
		id types.VarInt
		fields []reflect.StructField
	}
)

func getSummary(packet interface{}) packetSummary {
	t := reflect.TypeOf(packet)
	packetId := types.VarInt(0)
	fields := make([]reflect.StructField, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		id := field.Tag.Get(StructTag)

		if len(id) > 0 {
			intId, err := strconv.ParseInt(id, 0, 32)

			if err != nil {
				log.Fatal(err)
			}

			packetId = types.VarInt(intId)
			continue
		}

		fields = append(fields, field)
	}

	return packetSummary{t.Name(), packetId, fields}
}

func getReadBody(sum packetSummary) string {
	lines := make([]string, 0)

	for _, field := range sum.fields {
		readInstName := TypesPkgSuffix + field.Type.Name() + DefaultInstPostfix
		readVal := fmt.Sprintf(ReadSingleTemplate, field.Name, readInstName, field.Name, field.Name, TypesPkgSuffix+ field.Type.Name())
		lines = append(lines, readVal)
	}

	return strings.Join(lines, Newline)
}

func getWriteBody(sum packetSummary) string {

	return fmt.Sprintf(WriteTemplate, sum.name)
}

func getIDBody(sum packetSummary) string {

	return fmt.Sprintf(GetIDTemplate, sum.name, sum.id)
}