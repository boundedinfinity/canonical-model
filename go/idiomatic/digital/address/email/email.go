package email

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/specifications/rfc/rfc5322"
)

type Email struct {
	Id      ider.Id              `json:"id"`
	Name    string               `json:"name"`
	Address rfc5322.EmailAddress `json:"address"`
}
