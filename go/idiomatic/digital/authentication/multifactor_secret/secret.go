package multifactor_secret

type Secret interface {
	GetKind() Kind
	GetName() string
}
