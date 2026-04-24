package dental

import "github.com/boundedinfinity/canonical_model/go/idiomatic/business"

type Insurer struct {
	Name     string                 `json:"name"`
	Business business.BusinessModel `json:"business"`
}
