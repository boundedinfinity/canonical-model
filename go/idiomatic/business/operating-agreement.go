package business

import "github.com/boundedinfinity/canonical_model/go/idiomatic/digital_document"

type OperatingAgreement struct {
	Document digital_document.DocumentV1 `json:"document,omitempty"`
}
