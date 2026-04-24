package details

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type DateTime struct {
	Id            ider.Id  `json:"id"`
	Name          string   `json:"name"`
	Abbreviations []string `json:"abbreviations"`
	Description   string   `json:"description"`
	Start         rfc3339date.Rfc3339DateTime
	End           rfc3339date.Rfc3339DateTime
}

func (this DateTime) String() string {
	return this.Name
}
