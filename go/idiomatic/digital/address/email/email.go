package email

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/specification/rfc/rfc5322"
)

type Email struct {
	Id      ider.Id              `json:"id"`
	Name    string               `json:"name"`
	Address rfc5322.EmailAddress `json:"address"`
}
