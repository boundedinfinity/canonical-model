package battery

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/measurement"
)

type Battery struct {
	Id         id.Id                  `json:"id,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Dimensions measurement.Dimensions `json:"dimensions,omitempty"`
	Electrical measurement.Electrical `json:"electrical,omitempty"`
}
