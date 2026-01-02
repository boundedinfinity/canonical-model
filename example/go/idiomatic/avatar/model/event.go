package model

import (
	"encoding/json"

	"github.com/boundedinfinity/rfc3339date"
)

type EventTypeDiscriminator struct {
	Header EventHeader `json:"header,omitempty"`
}

type EventHeader struct {
	Kind           string                      `json:"type,omitempty"`
	ReceivedAt     rfc3339date.Rfc3339DateTime `json:"received-at,omitempty"`
	SequenceNumber int                         `json:"sequence-number,omitempty"`
}

type RawEvent struct {
	Header EventHeader     `json:"header,omitempty"`
	Value  json.RawMessage `json:"value,omitempty"`
}

type TypedEvent struct {
	Header EventHeader `json:"header,omitempty"`
	Value  any         `json:"value,omitempty"`
}

func (t TypedEvent) Raw() RawEvent {
	bs, _ := json.Marshal(t.Value)

	return RawEvent{
		Header: t.Header,
		Value:  bs,
	}
}

type ConcreteEvent[T any] struct {
	Header EventHeader `json:"header,omitempty"`
	Value  T           `json:"value,omitempty"`
}

func (t ConcreteEvent[T]) Typed() TypedEvent {
	return TypedEvent{
		Header: t.Header,
		Value:  t.Value,
	}
}

func (t ConcreteEvent[T]) Raw() RawEvent {
	bs, _ := json.Marshal(t.Value)

	return RawEvent{
		Header: t.Header,
		Value:  bs,
	}
}
