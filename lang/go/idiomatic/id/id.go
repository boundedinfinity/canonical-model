package id

import "github.com/google/uuid"

type Id struct {
	uuid.UUID
}

func (t Id) String() string {
	return t.UUID.String()
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
