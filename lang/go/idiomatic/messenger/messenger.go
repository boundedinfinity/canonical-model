package messenger

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

func New() *Messenger {
	marshaler := marshaler.NewWrapped()

	marshaler.Register(
		contact.Contact{},
		people.Person{},
		people.Name{},
	)

	repository, err := NewFileMessageRepository("", "")

	if err != nil {
		panic(err)
	}

	return &Messenger{
		marshaler:  marshaler,
		repository: repository,
	}
}

type Messenger struct {
	marshaler  *marshaler.WrappedMarshaler
	repository MessageRepository
}

func (t *Messenger) Close() error {
	var errs []error

	if err := t.repository.Close(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (t *Messenger) Marshal(item any) ([]byte, error) {
	return t.marshaler.Marshal(item)
}

func (t *Messenger) Unmarshal(data []byte) (any, error) {
	return t.marshaler.Unmarshal(data)
}

func (t *Messenger) Send(item any) error {
	bs, err := t.marshaler.Marshal(item)

	fmt.Println(string(bs))

	return err
}

func (t *Messenger) Handle(data []byte) error {
	message, err := t.marshaler.Unmarshal(data)

	if err != nil {
		return err
	}

	switch value := message.(type) {
	case people.Person:
		fmt.Printf("handled %v", value)
	default:
		return fmt.Errorf("can't handle %v", string(data))
	}

	return nil
}
