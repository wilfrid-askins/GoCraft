package server

import (
	"github.com/hashicorp/consul/api"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Description string `yaml:"description"`
		MaxPlayers  int    `yaml:"max_players"`
	}
)

func LoadConfig(logger *zap.Logger) Config {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		logger.Fatal("failed to connect to consul", zap.Field{Key:"msg", String: errors.Message(err)})
	}

	var out []byte
	_, err = client.Raw().Query("server", &out, nil)
	if err != nil {
		logger.Fatal("failed to query consul", zap.Field{Key:"msg", String: errors.Message(err)})
	}

	config := Config{}
	err = yaml.Unmarshal(out, &config)
	if err != nil {
		logger.Fatal("failed to unmarshal consul config", zap.Field{Key:"msg", String: errors.Message(err)})
	}

	return config
}

// see https://github.com/Tnze/go-mc/tree/master/nbt
// consider http://cassandra.apache.org/
// consider https://kafka.apache.org/intro
// use redis

// look at https://wiki.openstreetmap.org/wiki/Downloading_data#All_data_at_once

