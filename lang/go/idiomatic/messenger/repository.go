package messenger

import (
	"github.com/boundedinfinity/rfc3339date"
)

type MessageRepository interface {
	Store(...Message) error
	Load(rfc3339date.Rfc3339DateTime) (chan Message, error)
	Close() error
}
