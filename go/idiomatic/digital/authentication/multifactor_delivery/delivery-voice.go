package multifactor_delivery

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone"
)

var _ Delivery = &VoiceModel{}

type VoiceModel struct {
	Id    ider.Id          `json:"id"`
	Name  string           `json:"name"`
	Phone phone.PhoneModel `json:"phone"`
}

func (this VoiceModel) GetName() string {
	return this.Name
}

func (_ VoiceModel) GetKind() Kind {
	return Kinds.Voice
}
