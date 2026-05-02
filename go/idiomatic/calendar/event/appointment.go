package event

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

func Appointment() Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Appointment,
	}
}
