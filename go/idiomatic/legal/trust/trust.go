package trust

import "github.com/boundedinfinity/canonical-model/go/idiomatic/legal/benificiary"

type Trust struct {
	Kind        Kind
	Grantor     string
	Trustee     string
	Beneficiary benificiary.Benificiary
}
