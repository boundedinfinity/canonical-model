package label

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/google/uuid"
)

type LabelName struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type LabelType interface {
	string | float64 | int | rfc3339date.Rfc3339Date | rfc3339date.Rfc3339DateTime
}

type Label[V LabelType] struct {
	Id          uuid.UUID `json:"id"`
	Name        LabelName `json:"name,omitempty"`
	Value       V         `json:"value,omitempty"`
	Description string    `json:"description,omitempty"`
}
