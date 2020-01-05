package types

import "bufio"

type (
	CraftType interface {
		Read(*bufio.Reader) (interface{}, error)
		Write() error
	}
)
