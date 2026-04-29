package postal_codes

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type PostalCode interface {
	GetKind() Kind
	GetId() ider.Id
	String() string
}
