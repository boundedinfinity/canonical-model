package email_address

import (
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/email_address"
)

type EmailAddressItem model.Item[email_address.EmailAddress]

type EmailAddressEvent = model.ConcreteEvent[email_address.EmailAddress]
