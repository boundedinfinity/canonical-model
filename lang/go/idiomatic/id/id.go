package id

import "github.com/google/uuid"

type Id uuid.UUID

var (
	zero Id
)

func IsEmpty(id Id) bool {
	return id == zero
}
