package identifier

type Identifier interface {
	GetKind() Kind
	GetIdentifier() string
}
