package ider

import "github.com/google/uuid"

var (
	Zero = Id{}
)

type Id uuid.UUID

func (this Id) String() string {
	return uuid.UUID(this).String()
}

func (this Id) IsZero() bool {
	return this == Zero
}

func New() Id {
	return Id(uuid.New())
}

func Parse(s string) (Id, error) {
	id, err := uuid.Parse(s)
	return Id(id), err
}

func MustParse(s string) Id {
	return Id(uuid.MustParse(s))
}
