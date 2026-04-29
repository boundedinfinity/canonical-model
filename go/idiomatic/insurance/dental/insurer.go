package dental

import "github.com/boundedinfinity/canonical-model/go/idiomatic/business"

type Insurer struct {
	Name     string            `json:"name"`
	Business business.Business `json:"business"`
}
