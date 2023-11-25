package digital_document

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Document struct {
	Id       id.Id              `json:"id,omitempty"`
	Location string             `json:"location,omitempty"`
	Name     string             `json:"name,omitempty"`
	Type     mime_type.MimeType `json:"type,omitempty"`
	Audit    audit.Audit        `json:"audit,omitempty"`
}
