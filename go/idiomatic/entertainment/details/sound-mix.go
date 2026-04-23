package details

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type SoundMix struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
