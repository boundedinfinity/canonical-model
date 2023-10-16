package internal

import (
	"github.com/boundedinfinity/schema/idiomatic/label"
	"github.com/google/uuid"
)

type WithIdDesc struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
}

type WithIdDescFormat[T ~string] struct {
	WithIdDesc
	Format T `json:"format,omitempty"`
}

type WithIdDescLabel[T ~string] struct {
	WithIdDesc
	Labels []label.Label
}
