package business

import "github.com/boundedinfinity/canonical_model/go/idiomatic/digital/document"

type OperatingAgreement struct {
	Document document.Document `json:"document,omitempty"`
}
