package mime_type

import "errors"

var (
	ErrMimeType = errors.New("MIME Type")
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MIME Type Model
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MimeTypeModel struct {
	MimeTypes      []string
	FileExtensions []string
	Abbreviations  []string
	Name           string
	Description    string
}

func (this MimeTypeModel) Primary() mimeTypePrimary {
	return mimeTypePrimary{parent: this}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Primary
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type mimeTypePrimary struct {
	parent MimeTypeModel
}

func (this mimeTypePrimary) PrimaryMimeType() string {
	if len(this.parent.MimeTypes) > 0 {
		return this.parent.MimeTypes[0]
	}
	return ""
}

func (this mimeTypePrimary) PrimaryFileExtention() string {
	if len(this.parent.FileExtensions) > 0 {
		return this.parent.FileExtensions[0]
	}
	return ""
}

func (this mimeTypePrimary) PrimaryAbbreviation() string {
	if len(this.parent.Abbreviations) > 0 {
		return this.parent.Abbreviations[0]
	}
	return ""
}
