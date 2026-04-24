package image

import "github.com/boundedinfinity/canonical_model/go/idiomatic/digital/mime_type"

var _ Image = &ImageData{}

type ImageData struct {
	MimeType mime_type.MimeTypeModel `json:"mime_type"`
	Data     []byte                  `json:"data"`
}

func (_ ImageData) GetKind() Kind {
	return Kinds.Data
}

func (this *ImageData) GetMimeType() mime_type.MimeTypeModel {
	return this.MimeType
}
