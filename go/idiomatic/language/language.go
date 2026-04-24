package language

import "github.com/boundedinfinity/canonical_model/go/idiomatic/specifications/iso/iso639"

type Language struct {
	Iso iso639.Language `json:"iso"`
}
