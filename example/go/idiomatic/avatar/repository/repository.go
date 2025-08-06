package repository

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
)

type MessageRepository interface {
	Store(...model.TypedEvent) error
	Load(rfc3339date.Rfc3339DateTime) error
	Close() error
}
