package note

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type Notes struct {
	Id    ider.Id `json:"id"`
	Notes []Note  `json:"notes"`
}
