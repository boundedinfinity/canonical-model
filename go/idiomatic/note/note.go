package note

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type Note struct {
	Id        ider.Id                     `json:"id"`
	Text      string                      `json:"text"`
	CreatedAt rfc3339date.Rfc3339DateTime `json:"created-at"`
}
