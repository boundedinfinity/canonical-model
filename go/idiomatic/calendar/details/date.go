package details

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type Date struct {
	Id            ider.Id  `json:"id"`
	Name          string   `json:"name"`
	Abbreviations []string `json:"abbreviations"`
	Description   string   `json:"description"`
	Start         rfc3339date.Rfc3339Date
	End           rfc3339date.Rfc3339Date
}

func (this Date) String() string {
	return this.Name
}

func Day(name string, abbreviations []string, description string, date rfc3339date.Rfc3339Date) Date {
	return Date{
		Name:          name,
		Abbreviations: abbreviations,
		Description:   description,
		Start:         date,
		End:           date,
	}
}
