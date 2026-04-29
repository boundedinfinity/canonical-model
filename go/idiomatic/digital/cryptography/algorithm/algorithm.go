package algorithm

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type AlgorithmModel struct {
	Id          ider.Id    `json:"id"`
	Name        string     `json:"name"`
	MaximumBits int        `json:"maximum-bits"`
	MinimumBits int        `json:"minimum-bits"`
	Functions   []Function `json:"functions"`
}
