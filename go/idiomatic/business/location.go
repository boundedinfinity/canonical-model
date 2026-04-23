package business

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/location"
	"github.com/boundedinfinity/canonical_model/idiomatic/phone"
)

type Location struct {
	Id       ider.Id            `json:"id"`
	Location location.Location  `json:"location"`
	Roles    []RoleModel        `json:"roles"`
	Emails   []email.Email      `json:"emails"`
	Phones   []phone.PhoneModel `json:"phones"`
}
