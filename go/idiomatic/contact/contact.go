package contact

// https://en.wikipedia.org/wiki/VCard

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/mailing_address"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/people"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone"
)

type Contact struct {
	Id              ider.Id                        `json:"id,omitempty"`
	Person          people.Person                  `json:"person,omitempty"`
	Relationships   []RelationShip                 `json:"relationships,omitempty"`
	Telephones      []phone.NanpNumber             `json:"telephones,omitempty"`
	MailingAddress  mailing_address.MailingAddress `json:"mailing-address,omitempty"`
	DeliveryAddress mailing_address.MailingAddress `json:"delivery-address,omitempty"`
}

func (t Contact) Name() string {
	return t.Person.String()
}
