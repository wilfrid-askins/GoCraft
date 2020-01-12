package config

import "GoCraft/net/state"

type (
	Summary struct {
		Version `json:"version"`
		Players `json:"players"`
		Description `json:"description"`
	}

	Version struct {
		Name string `json:"name"`
		Protocol int `json:"protocol"`
	}

	Players struct {
		Max int `json:"max"`
		Online int `json:"online"`
		Sample []Player `json:"players"`
	}

	Player struct {
		Name string `json:"name"`
		Id string `json:"id"`
	}

	Description struct {
		Text string `json:"text"`
	}
)

func GetSummary(config Config, state state.State) Summary {

	return Summary{}
}
