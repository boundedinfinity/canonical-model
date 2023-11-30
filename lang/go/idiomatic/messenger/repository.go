package messenger

import (
	"github.com/boundedinfinity/rfc3339date"
)

type MessageRepository interface {
	Store(...RawMessage) error
	Load(rfc3339date.Rfc3339DateTime) (chan StoredMessage, error)
	Close() error
}
