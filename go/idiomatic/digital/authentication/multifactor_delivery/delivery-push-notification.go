package multifactor_delivery

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ Delivery = &PushNotificationModel{}

type PushNotificationModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this PushNotificationModel) GetName() string {
	return this.Name
}

func (_ PushNotificationModel) GetKind() Kind {
	return Kinds.PushNotification
}
