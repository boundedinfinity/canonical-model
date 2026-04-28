package mime_type

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	ErrMimeTypeFind     = errorer.New("Find")
	errMimeTypeFindFunc = errorer.Func(ErrMimeType, ErrMimeTypeFind)
)

func contains(slice []string, term string) bool {
	fn := func(item string) bool { return stringer.EqualIgnoreCase(item, term) }
	return slicer.ContainsFunc(fn, slice...)
}

func FindErr(terms ...string) ([]MimeTypeModel, error) {
	var found []MimeTypeModel
	var err error

	for i, term := range terms {
		for _, mt := range all {
			if stringer.EqualIgnoreCase(term, mt.Name) {
				found = append(found, mt)
				continue
			}

			if contains(mt.MimeTypes, term) {
				found = append(found, mt)
				continue
			}

			if contains(mt.FileExtensions, term) {
				found = append(found, mt)
				continue
			}

			if contains(mt.Abbreviations, term) {
				found = append(found, mt)
				continue
			}

			return found, errMimeTypeFindFunc("no MIME type found for terms[%d]: '%s'", i, term)
		}
	}

	return found, err
}

func FindOk(terms ...string) ([]MimeTypeModel, bool) {
	found := Find(terms...)
	return found, len(found) > 0
}

func Find(terms ...string) []MimeTypeModel {
	var found []MimeTypeModel

	for _, term := range terms {
		for _, mt := range all {
			if stringer.EqualIgnoreCase(term, mt.Name) {
				found = append(found, mt)
				continue
			}

			if contains(mt.MimeTypes, term) {
				found = append(found, mt)
				continue
			}

			if contains(mt.FileExtensions, term) {
				found = append(found, mt)
				continue
			}

			if contains(mt.Abbreviations, term) {
				found = append(found, mt)
				continue
			}
		}
	}

	return found
}
