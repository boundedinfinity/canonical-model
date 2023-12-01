package messenger

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
)

func New() *Messenger {
	repository, err := NewFileMessageRepository("", "")

	if err != nil {
		panic(err)
	}

	return &Messenger{
		marshaler:  NewMarshaller(),
		repository: repository,
		handlers:   map[string]func(){},
	}
}

type Messenger struct {
	marshaler  *marshaler.TypedMarshaler
	repository MessageRepository
	handlers   map[string]func()
}

func (t *Messenger) Close() error {
	var errs []error

	if err := t.repository.Close(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (t *Messenger) Send(item any) error {
	bs, err := t.marshaler.Marshal(item)

	fmt.Println(string(bs))

	return err
}
