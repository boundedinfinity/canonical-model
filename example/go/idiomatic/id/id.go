package id

import (
	"encoding/json"
	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/schema/idiomatic/model"
	"github.com/google/uuid"
)

var (
	ErrIdInvalid = errors.New("invalid id")
)

var _ model.Validator = &Id{}

type Id struct {
	uuid.UUID
}

func (this Id) String() string {
	return this.UUID.String()
}

func (this Id) Validate() error {
	return nil
}

func (this Id) IsZero() bool {
	return reflecter.IsZero[Id](this)
}

func (this Id) MarshalJSON() ([]byte, error) {
	if this.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(this.String())
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
