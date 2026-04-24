package phone

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/label"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone/carrier"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone/number"
	"github.com/boundedinfinity/idiomatic/errorer"
)

var (
	ErrPhone = errorer.New("phone")
)

type PhoneModel struct {
	Id      ider.Id         `json:"id"`
	Name    string          `json:"name"`
	Number  number.Number   `json:"number"`
	Carrier carrier.Carrier `json:"carrier"`
	Kind    Kind            `json:"kind"`
	Labels  []label.Labels  `json:"labels"`
}
