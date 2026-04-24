package uri

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/specifications/rfc/rfc2396"
)

type Url struct {
	Id      ider.Id     `json:"id"`
	Name    string      `json:"name"`
	Address rfc2396.Url `json:"address"`
}
