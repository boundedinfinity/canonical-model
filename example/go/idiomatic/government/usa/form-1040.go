package usa

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// https://www.irs.gov/pub/irs-pdf/f1040.pdf

type Form1040 struct {
	Id id.Id `json:"id,omitempty"`
}

type Form1040_2023 struct {
	Id id.Id `json:"id,omitempty"`
}

type Form1040_2024 struct {
	Id id.Id `json:"id,omitempty"`
}
