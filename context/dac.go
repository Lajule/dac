package context

import (
	"github.com/Lajule/dac/ent"
)

type KeyType string

type ValueType struct {
	Version string
	Client  *ent.Client
}

const (
	KeyName = "DAC"

	Key = KeyType(KeyName)
)
