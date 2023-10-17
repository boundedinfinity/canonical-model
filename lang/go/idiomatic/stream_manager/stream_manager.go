package stream_manager

import (
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
)

type StreamManager struct {
	mm *marshaler.MashalerManager
	m  map[string][]func(any, error) error
}

func New() *StreamManager {
	return &StreamManager{
		mm: marshaler.New(),
		m:  make(map[string][]func(any, error) error),
	}
}

func (t *StreamManager) Register(instance any, fn func(any, error) error) error {
	if err := t.mm.Register(instance); err != nil {
		return err
	}

	name := t.fqpp(instance)

	if _, ok := t.m[name]; !ok {
		t.m[name] = make([]func(any, error) error, 0)
	}

	t.m[name] = append(t.m[name], fn)

	return nil
}

func (t StreamManager) Handle(data []byte) error {
	instance, err := t.mm.Unmarshal(data)

	if err != nil {
		return err
	}

	name := t.fqpp(instance)

	if fns, ok := t.m[name]; ok {
		for _, fn := range fns {
			if err := fn(instance, nil); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (t StreamManager) fqpp(instance any) string {
	typ := reflect.TypeOf(instance)
	return fmt.Sprintf("%v/%v", typ.PkgPath(), typ.Name())
}
