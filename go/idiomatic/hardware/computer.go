package hardware

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/identifier"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type Contact struct {
	Id           ider.Id                 `json:"id,omitempty"`
	SerialNumber identifier.SerialNumber `json:"serial-number,omitempty"`
}
