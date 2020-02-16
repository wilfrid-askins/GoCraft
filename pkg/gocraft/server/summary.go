package server

import (
	"GoCraft/pkg/gocraft/server/play"
)

type (
	Summary struct {
		Version     `json:"version"`
		Players     `json:"players"`
		Description `json:"description"`
	}

	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	}

	Players struct {
		Max    int      `json:"max"`
		Online int      `json:"online"`
		Sample []Player `json:"players"`
	}

	Player struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	Description struct {
		Text string `json:"text"`
	}
)

func (config Config) GetSummary(state *play.State) Summary {

	return Summary{
		Version: Version{
			Name:     "1.15.1",
			Protocol: 575,
		},
		Players: Players{
			Max:    config.MaxPlayers,
			Online: 99,
			Sample: []Player{
				{"Test", "afdf"},
			},
		},
		Description: Description{Text: config.Description},
	}
}
