package main

import (
	"GoCraft/internal"
	"GoCraft/pkg/gocraft"
	"fmt"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting...")
	logger := zap.NewExample()
	defer logger.Sync()
	conf := gocraft.LoadConfig(logger)

	resolver := internal.NewResolver(conf, logger)
	if err := resolver.Server().Listen(); err != nil {
		logger.Fatal("server listen failed", zap.String("msg", errors.Message(err)))
	}
}
