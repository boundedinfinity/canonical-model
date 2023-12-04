package mailing_address

import (
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
)

type MailingAddressItem model.Item[mailing_address.MailingAddress]

type MailingAddressEvent = model.ConcreteEvent[mailing_address.MailingAddress]
