package messenger

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

func NewMarshaller() *marshaler.TypedMarshaler {
	marshaler := marshaler.NewTyped()

	marshaler.Register(
		contact.Contact{},
		people.Person{},
		people.Prefix{},
		people.Suffix{},
		people.Name{},
		mailing_address.MailingAddress{},
	)

	return marshaler
}
