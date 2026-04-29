package multifactor_delivery

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
)

var _ Delivery = &ShortMessageServiceModel{}

type ShortMessageServiceModel struct {
	Id    ider.Id          `json:"id"`
	Name  string           `json:"name"`
	Phone phone.PhoneModel `json:"phone"`
}

func (this ShortMessageServiceModel) GetName() string {
	return this.Name
}

func (_ ShortMessageServiceModel) GetKind() Kind {
	return Kinds.ShortMessageService
}
