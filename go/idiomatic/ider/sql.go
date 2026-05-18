package ider

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

func (this *Id) Scan(src interface{}) error {
	return (*uuid.UUID)(this).Scan(src)
}

func (this Id) Value() (driver.Value, error) {
	return (uuid.UUID)(this).Value()
}
