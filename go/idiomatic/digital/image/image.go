package image

import "github.com/boundedinfinity/canonical-model/go/idiomatic/digital/mime_type"

type Image interface {
	GetKind() Kind
	GetMimeType() mime_type.MimeTypeModel
}
