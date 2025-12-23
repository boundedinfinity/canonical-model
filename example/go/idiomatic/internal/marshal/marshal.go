package marshal

type Typed[T any] struct {
	Id   string `json:"type-id"`
	Type T      `json:"type"`
}

func Typedf[T any](typeId string, t T) Typed[T] {
	return Typed[T]{
		Id:   typeId,
		Type: t,
	}
}
