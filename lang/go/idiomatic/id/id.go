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
	return reflecter.IsZero[Id](t)
}

func (t Id) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(t.String())
}

var Ids = ids{}

type ids struct {
	zero Id
}

func (this ids) New() Id {
	return Id{
		UUID: uuid.New(),
	}
}

func (this ids) MustParse(s string) Id {
	return Id{
		UUID: uuid.MustParse(s),
	}
}

func (this ids) Parse(s string) (Id, error) {
	id, err := uuid.Parse(s)
	return Id{UUID: id}, err
}

func (this ids) IsEmpty(id Id) bool {
	return id == this.zero
}
