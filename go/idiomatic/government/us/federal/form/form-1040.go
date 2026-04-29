package form

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// https://www.irs.gov/pub/irs-pdf/f1040.pdf

type Form1040 struct {
	Id ider.Id `json:"id,omitempty"`
}

type Form1040_2023 struct {
	Id ider.Id `json:"id,omitempty"`
}

type Form1040_2024 struct {
	Id ider.Id `json:"id,omitempty"`
}
