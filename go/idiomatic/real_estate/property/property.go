package property

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/image"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/mailing_address"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/real_estate/taxes"
	"github.com/boundedinfinity/rfc3339date"
)

type Property struct {
	Id       ider.Id                 `json:"id"`
	Address  mailing_address.Address `json:"address,omitempty"`
	Images   []image.Image           `json:"images,omitempty"`
	Features []PropertyFeature       `json:"features,omitempty"`
	BuiltIn  rfc3339date.Rfc3339Date `json:"built_in,omitempty"`
	Size     int                     `json:"size,omitempty"`
	Acreage  float64                 `json:"acreage,omitempty"`
	Taxes    []taxes.Tax             `json:"taxes,omitempty"`
}

type PropertyFeature struct {
	Feature Feature `json:"feature,omitempty"`
	Count   int     `json:"count,omitempty"`
}
