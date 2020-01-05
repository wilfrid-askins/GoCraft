package main

import (
	"GoCraft/net/packets/client"
	"GoCraft/net/packets/server"
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
	package %s

	import (
		"GoCraft/net/types"
		"bufio"
	)`
	ReadTemplate = `
	func (p *%s) Read(in *bufio.Reader) error {
		%s
		return nil
	}`
	ReadSingleTemplate = `
	val%s, err := %s.Read(in)
	if err != nil {
		return err
	}
	p.%s = val%s.(%s)`
	WriteTemplate = `
	func (p *%s) Write(out *bufio.Writer) error {
		%s

		return nil
	}`
	DeclareError = `var err error`
	WriteSingleTemplate = `
	err = %s.Write(out)
	if err != nil {
		return err
	}`
	GetIDTemplate = `
	func (p *%s) GetID() types.VarInt {
		return %d
	}`
	DefaultInstPostfix = "Default"
	TypesPkgSuffix     = "types."
	StructTag = "packet"
	)

//go:generate go run gen.go
//go:generate gofmt -w ./net/packets/client/client_gen.go
//go:generate gofmt -w ./net/packets/server/server_gen.go
func main() {
	fmt.Println("Generating packet code")

	writeToFile("./net/packets/client/client_gen.go", "client", []interface{}{
		client.Handshake{},
		client.Request{},
		client.ChatMessage{},
	})

	writeToFile("./net/packets/server/server_gen.go", "server", []interface{}{
		server.Response{},
	})
}

func writeToFile(filePath, pkgName string, packetSlice []interface{}) {
	fmt.Println("Writing to " + filePath)

	code := make([]byte, 0)
	code = append(code, fmt.Sprintf(FileHeader, pkgName)...)

	for _, p := range packetSlice {
		sum := getSummary(p)
		readBlock := fmt.Sprintf(ReadTemplate, sum.name, getReadBody(sum))
		writeBlock :=  fmt.Sprintf(WriteTemplate, sum.name, getWriteBody(sum))
		idBlock := getIDBody(sum)

		code = append(code, readBlock...)
		code = append(code, Newline...)
		code = append(code, writeBlock...)
		code = append(code, Newline...)
		code = append(code, idBlock...)
		code = append(code, Newline...)
	}

	err := ioutil.WriteFile(filePath, code, 0644)

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
	lines := make([]string, 0)

	if len(sum.fields) > 0 {
		lines = append(lines, DeclareError)
	}

	for _, field := range sum.fields {
		writeInstName := "p." + field.Name
		writeVal := fmt.Sprintf(WriteSingleTemplate, writeInstName)
		lines = append(lines, writeVal)
	}

	return strings.Join(lines, Newline)
}

func getIDBody(sum packetSummary) string {

	return fmt.Sprintf(GetIDTemplate, sum.name, sum.id)
}