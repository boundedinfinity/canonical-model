package digital_document

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Document struct {
	Id   id.Id              `json:"id,omitempty"`
	Path string             `json:"path,omitempty"`
	Type mime_type.MimeType `json:"type,omitempty"`
}
