package persistence

type Item[T any] struct {
	Value T
	Audit Audit
}
