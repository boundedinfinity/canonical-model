package digital_document

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// /////////////////////////////////////////////////////////////////
//  V1
// /////////////////////////////////////////////////////////////////

type DocumentV1 struct {
	Id             id.Id                   `json:"id,omitempty"`
	Location       string                  `json:"location,omitempty"`
	Name           string                  `json:"name,omitempty"`
	Type           mime_type.MimeType      `json:"type,omitempty"`
	CreatedOnDate  rfc3339date.Rfc3339Date `json:"created-on-date,omitempty"`
	UploadedOnDate rfc3339date.Rfc3339Date `json:"upload-on-date,omitempty"`
}

func (t DocumentV1) Upgrade() DocumentV2 {
	return DocumentV2{
		Id:         t.Id,
		Location:   t.Location,
		Name:       t.Name,
		Type:       t.Type,
		CreatedOn:  t.CreatedOnDate,
		UploadedOn: t.UploadedOnDate,
	}
}

// /////////////////////////////////////////////////////////////////
//  V2
// /////////////////////////////////////////////////////////////////

type DocumentV2 struct {
	Id         id.Id                   `json:"id,omitempty"`
	Location   string                  `json:"location,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Type       mime_type.MimeType      `json:"type,omitempty"`
	CreatedOn  rfc3339date.Rfc3339Date `json:"created-on,omitempty"`
	UploadedOn rfc3339date.Rfc3339Date `json:"upload-on,omitempty"`
}

func (t DocumentV2) Downgrade() DocumentV1 {
	return DocumentV1{
		Id:             t.Id,
		Location:       t.Location,
		Name:           t.Name,
		Type:           t.Type,
		CreatedOnDate:  t.CreatedOn,
		UploadedOnDate: t.UploadedOn,
	}
}
