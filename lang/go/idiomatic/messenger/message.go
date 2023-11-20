package messenger

import "github.com/boundedinfinity/rfc3339date"

type Message[T any] struct {
	Timestamp rfc3339date.Rfc3339DateTime
	Type      string
	Value     T
}
