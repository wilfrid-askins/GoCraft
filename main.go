package main

import (
	"GoCraft/pkg/gocraft"
	"GoCraft/pkg/gocraft/server"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting...")
	logger := zap.L()
	conf := server.LoadConfig(logger)

	resolver := gocraft.NewResolver(conf, logger)
	resolver.Server().Listen()
}
