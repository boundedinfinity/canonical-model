package messenger

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type MessageHeader struct {
	Kind       string                      `json:"kind,omitempty"`
	Source     id.Id                       `json:"id,omitempty"`
	SentAt     rfc3339date.Rfc3339DateTime `json:"sent-at,omitempty"`
	ReceivedAt rfc3339date.Rfc3339DateTime `json:"received-at,omitempty"`
}

type messageRaw struct {
	Header  MessageHeader   `json:"header,omitempty"`
	Message json.RawMessage `json:"message,omitempty"`
}

type Message struct {
	Header  MessageHeader `json:"header,omitempty"`
	Message any           `json:"message,omitempty"`
}

func NewMessage(item any) Message {
	return Message{
		Header: MessageHeader{
			Kind:   reflecter.Instances.QualifiedName(item),
			SentAt: rfc3339date.DateTimes.Now(),
		},
		Message: item,
	}
}
