package business

import "github.com/boundedinfinity/schema/idiomatic/digital_document"

type OperatingAgreement struct {
	Document digital_document.DocumentV1 `json:"document,omitempty"`
}
