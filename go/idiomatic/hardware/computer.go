package hardware

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/id"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Contact struct {
	Id           ider.Id         `json:"id,omitempty"`
	SerialNumber id.SerialNumber `json:"serial-number,omitempty"`
}
