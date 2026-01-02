package marshaler

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
)

// type EventMarshaler marshaler.KindMarshaler[string, model.EventTypeDiscriminator]

type EventMarshaler struct {
	// tm *TypeMapper
	km *marshaler.KindMarshaler[model.EventTypeDiscriminator]
}

func (t EventMarshaler) Marshal(val any) ([]byte, error) {
	return t.km.Marshal(val)
}

func (t EventMarshaler) Unmarshal(bs []byte) (any, error) {
	return t.km.Unmarshal(bs)
}

func (t *EventMarshaler) Formatted(v bool) {
	t.km.Formatted = v
}

func NewMarshaller() *EventMarshaler {
	km := EventMarshaler{
		km: marshaler.NewKind(
			model.EventTypeDiscriminator{},
			func(d model.EventTypeDiscriminator) string { return d.Header.Kind },
		),
		// tm: NewAvatarTypes(),
	}

	// for _, info := range km.tm.Slice() {
	// 	km.km.RegisterValue(info.typeName, info.typ)
	// }

	registerAvatarTypes(km.km)

	return &km
}
