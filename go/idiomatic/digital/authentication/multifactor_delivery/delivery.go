package multifactor_delivery

type Delivery interface {
	GetKind() Kind
	GetName() string
}
