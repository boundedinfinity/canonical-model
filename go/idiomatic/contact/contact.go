package contact

// https://en.wikipedia.org/wiki/VCard

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/mailing_address"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
)

type Contact struct {
	Id              ider.Id                 `json:"id,omitempty"`
	Person          person.PersonModel      `json:"person,omitempty"`
	Relationships   []Relationship          `json:"relationships,omitempty"`
	Telephones      []phone.PhoneModel      `json:"phone,omitempty"`
	MailingAddress  mailing_address.Address `json:"mailing-address,omitempty"`
	DeliveryAddress mailing_address.Address `json:"delivery-address,omitempty"`
}

func (t Contact) Name() string {
	return t.Person.String()
}
