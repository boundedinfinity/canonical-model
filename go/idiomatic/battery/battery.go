package battery

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type Battery struct {
	Id   ider.Id `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
	// Dimensions measurement.Dimensions `json:"dimensions,omitempty"`
	// Electrical measurement.Electrical `json:"electrical,omitempty"`
}
