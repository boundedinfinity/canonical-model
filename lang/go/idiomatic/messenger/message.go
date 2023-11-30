package messenger

import (
	"encoding/json"

	"github.com/boundedinfinity/rfc3339date"
)

type RawMessage struct {
	Timestamp rfc3339date.Rfc3339DateTime
	Type      string
	Value     json.RawMessage
}

type StoredMessage struct {
	IngestedTimestamp rfc3339date.Rfc3339DateTime
	RawMessage
}

type TypedMessage[T any] struct {
	Timestamp rfc3339date.Rfc3339DateTime
	Type      string
	Value     T
}
