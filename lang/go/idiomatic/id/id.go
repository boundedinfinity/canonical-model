package id

import (
	"encoding/json"
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/google/uuid"
)

var (
	ErrIdInvalid = errors.New("invalid id")
)

type Id struct {
	uuid.UUID
}

func (t Id) String() string {
	return t.UUID.String()
}

func (t Id) Validate(groups ...string) error {
	return nil
}

func (t Id) IsZero() bool {
	return reflecter.Instances.IsZero(t)
}

func (t Id) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(t.String())
}

var (
	zero Id
)

func MustParse(s string) Id {
	return Id{
		UUID: uuid.MustParse(s),
	}
}

func Parse(s string) (Id, error) {
	var id Id

	if uuidId, err := uuid.Parse(s); err != nil {
		return id, err
	} else {
		return Id{UUID: uuidId}, nil
	}
}

func IsEmpty(id Id) bool {
	return id == zero
}
