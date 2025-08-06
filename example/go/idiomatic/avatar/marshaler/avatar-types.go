package marshaler

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model/email_address"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model/people"
)

func NewAvatarTypes() *TypeMapper {
	tm := NewTypeMapper()

	tm.RegisterValue(people.PersonEvent{}.Value.TypeName(), people.PersonEvent{})
	tm.RegisterValue(people.NameEvent{}.Value.TypeName(), people.NameEvent{})
	tm.RegisterValue(people.PrefixEvent{}.Value.TypeName(), people.PrefixEvent{})
	tm.RegisterValue(people.SuffixEvent{}.Value.TypeName(), people.SuffixEvent{})
	tm.RegisterValue(email_address.EmailAddressEvent{}.Value.TypeName(), email_address.EmailAddressEvent{})

	return tm
}

func registerAvatarTypes(tm *marshaler.KindMarshaler[model.EventTypeDiscriminator]) {
	tm.RegisterValue(people.PersonEvent{}.Value.TypeName(), people.PersonEvent{})
	tm.RegisterValue(people.NameEvent{}.Value.TypeName(), people.NameEvent{})
	tm.RegisterValue(people.PrefixEvent{}.Value.TypeName(), people.PrefixEvent{})
	tm.RegisterValue(people.SuffixEvent{}.Value.TypeName(), people.SuffixEvent{})
	tm.RegisterValue(email_address.EmailAddressEvent{}.Value.TypeName(), email_address.EmailAddressEvent{})
}
