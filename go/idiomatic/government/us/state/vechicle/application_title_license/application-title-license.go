package application_title_license

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/county"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/mailing_address"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/vehicle"
)

type ApplicationTitleLicense struct {
	Owner      person.Person           `json:"owner,omitempty"`
	Address    mailing_address.Address `json:"address,omitempty"`
	County     county.County           `json:"county,omitempty"`
	Vehicle    vehicle.Vehicle         `json:"vehicle,omitempty"`
	LienHolder business.Business       `json:"lien-holder,omitempty"`
}
