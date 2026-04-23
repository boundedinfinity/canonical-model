package affix

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type AffixKind string

const (
	Prefix AffixKind = "preifx"
	Suffix AffixKind = "suffix"
)

type Affix struct {
	Id            ider.Id   `json:"id"`
	Name          string    `json:"name"`
	Kind          AffixKind `json:"kind"`
	Abbreviations []string  `json:"abbreviations"`
	Description   string    `json:"description"`
	Category      Category  `json:"category"`
}

func (this Affix) String() string {
	return this.Name
}
