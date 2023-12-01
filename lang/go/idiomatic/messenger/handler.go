package messenger

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

func (t *Messenger) AddHandler(v any, fn func()) error {
	name := reflecter.Instances.QualifiedName(v)

	if _, ok := t.handlers[name]; !ok {
		t.handlers[name] = fn
	} else {
		fmt.Printf("already have handler for %v", name)
	}

	return nil
}

func (t *Messenger) HandleType(message any) error {
	switch value := message.(type) {
	case people.Person:
		fmt.Printf("handled %v", value)
	default:
		return fmt.Errorf("can't handle %v", message)
	}

	return nil
}

func (t *Messenger) HandleData(data []byte) error {
	message, err := t.marshaler.Unmarshal(data)

	if err != nil {
		return err
	}

	return t.HandleType(message)
}
