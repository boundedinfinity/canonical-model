package professional

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/person/affix"
)

type Credentials struct {
	Id          ider.Id       `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Prefix      []affix.Affix `json:"prefix"`
	Suffix      []affix.Affix `json:"suffix"`
}
