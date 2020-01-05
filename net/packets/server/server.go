package server

import "GoCraft/net/types"

type (
	Response struct {
		ID int `packet:"0x0"`
		JsonResponse types.CraftString `packet`
	}


)
