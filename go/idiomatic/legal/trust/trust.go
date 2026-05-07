package trust

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/legal/benificiary"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/legal/grantor"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/legal/trustee"
)

type Trust struct {
	Kind        Kind
	Grantor     grantor.Grantor
	Trustee     trustee.Trustee
	Beneficiary benificiary.Benificiary
}
