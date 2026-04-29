package business

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/address/email"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
)

type Location struct {
	Id       ider.Id            `json:"id"`
	Location location.Location  `json:"location"`
	Roles    []RoleModel        `json:"roles"`
	Emails   []email.Email      `json:"emails"`
	Phones   []phone.PhoneModel `json:"phones"`
}
