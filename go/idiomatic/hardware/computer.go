package hardware

import "github.com/boundedinfinity/canonical_model/idiomatic/id"

type Contact struct {
	Id           id.Id           `json:"id,omitempty"`
	SerialNumber id.SerialNumber `json:"serial-number,omitempty"`
}
