package context

import (
	"github.com/Lajule/dac/ent"
)

type Key string

type Value struct {
	Version string
	Client  *ent.Client
}

const (
	KeyName = Key("dac")
)
