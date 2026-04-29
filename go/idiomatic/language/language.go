package language

import "github.com/boundedinfinity/canonical-model/go/idiomatic/specification/iso/iso639"

type Language struct {
	Iso iso639.Language `json:"iso"`
}
