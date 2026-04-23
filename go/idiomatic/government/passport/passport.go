package passport

type Passport interface {
	GetKind() Kind
	GetId() string
}
