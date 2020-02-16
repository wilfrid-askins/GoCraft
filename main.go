package main

import (
	"GoCraft/pkg/gocraft"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting...")
	logger := zap.NewExample()
	defer logger.Sync()
	conf := gocraft.LoadConfig(logger)

	resolver := gocraft.NewResolver(conf, logger)
	resolver.Server().Listen()
}
