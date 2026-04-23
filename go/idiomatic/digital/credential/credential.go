package credential

type Credential interface {
	GetKind() Kind
	GetName() string
}
