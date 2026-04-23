package postal_codes

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type PostalCode interface {
	GetKind() Kind
	GetId() ider.Id
	String() string
}
