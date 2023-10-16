package label

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/google/uuid"
	"golang.org/x/exp/constraints"
)

type Label struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name LabelName `json:"name,omitempty"`
}

type ValueLabel[V constraints.Ordered | rfc3339date.Rfc3339Date | rfc3339date.Rfc3339DateTime] struct {
	Id    uuid.UUID `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value V         `json:"value,omitempty"`
}
