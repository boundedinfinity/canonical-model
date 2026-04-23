package image

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/digital/mime_type"
	"github.com/boundedinfinity/canonical_model/idiomatic/model/specifications/rfc/rfc2396"
)

var _ Image = &ImageReference{}

type ImageReference struct {
	MimeType mime_type.MimeTypeModel `json:"mime_type"`
	Url      rfc2396.Url             `json:"url"`
}

func (_ ImageReference) GetKind() Kind {
	return Kinds.Reference
}

func (this *ImageReference) GetMimeType() mime_type.MimeTypeModel {
	return this.MimeType
}
