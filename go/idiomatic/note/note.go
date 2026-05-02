package note

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/label"
	"github.com/boundedinfinity/rfc3339date"
)

type Note struct {
	Id        ider.Id                     `json:"id"`
	Text      string                      `json:"text"`
	CreatedAt rfc3339date.Rfc3339DateTime `json:"created-at"`
	Labels    []label.Label               `json:"labels"`
}
