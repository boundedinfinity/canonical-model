package show

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type Season struct {
	Id       ider.Id   `json:"id"`
	Season   int       `json:"season"`
	Episodes []Episode `json:"episodes"`
}
