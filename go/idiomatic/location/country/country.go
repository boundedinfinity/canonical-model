package country

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/specification/iso/iso3166"
)

type Country struct {
	Name string          `json:"name"`
	Iso  iso3166.Country `json:"iso"`
}

type countries struct {
	all   []Country
	toIso map[string]*Country
}

func (this countries) FindOk(term string) ([]Country, bool) {

	return nil, false
}
