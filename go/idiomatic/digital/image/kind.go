package image

type Kind string

type kinds struct {
	Unknown   Kind
	Data      Kind
	Reference Kind
}

var Kinds = kinds{
	Unknown:   "image.unknown",
	Data:      "image.data",
	Reference: "image.reference",
}
