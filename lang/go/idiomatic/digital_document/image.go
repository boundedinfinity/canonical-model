package digital_document

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Image struct {
	Id             id.Id                   `json:"id,omitempty"`
	Location       Location                `json:"location,omitempty"`
	Name           string                  `json:"name,omitempty"`
	Description    string                  `json:"description,omitempty"`
	Type           mime_type.MimeType      `json:"type,omitempty"`
	CreatedOnDate  rfc3339date.Rfc3339Date `json:"created-on-date,omitempty"`
	UploadedOnDate rfc3339date.Rfc3339Date `json:"upload-on-date,omitempty"`
}
