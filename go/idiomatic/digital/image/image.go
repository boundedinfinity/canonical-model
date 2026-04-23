package image

import "github.com/boundedinfinity/canonical_model/idiomatic/digital/mime_type"

type Image interface {
	GetKind() Kind
	GetMimeType() mime_type.MimeTypeModel
}
