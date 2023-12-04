package marshaler

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model/people"
)

// type EventMarshaler marshaler.KindMarshaler[string, model.EventTypeDiscriminator]

type EventMarshaler struct {
	km *marshaler.KindMarshaler[string, model.EventTypeDiscriminator]
}

func (t EventMarshaler) Marshal(val any) ([]byte, error) {
	return t.km.Marshal(val)
}

func (t EventMarshaler) Unmarshal(bs []byte) error {
	return t.km.Unmarshal(bs)
}

func (t *EventMarshaler) Formatted(v bool) {
	t.km.Formatted = v
}

func NewMarshaller() *EventMarshaler {
	km := EventMarshaler{
		km: marshaler.NewKind[string, model.EventTypeDiscriminator](),
	}

	km.km.RegisterDescriminator(
		model.EventTypeDiscriminator{},
		func(d model.EventTypeDiscriminator) string { return d.Header.Type },
	)

	km.km.RegisterType(people.PersonEvent{}.Value.TypeName(), people.PersonEvent{})
	km.km.RegisterType(people.NameEvent{}.Value.TypeName(), people.NameEvent{})

	return &km
}
