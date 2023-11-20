package modeller

type modellerBuilder[T any] struct {
	instance   *T
	contraints []Contraint[T]
}

func Modeller[T any](instance *T) *modellerBuilder[T] {
	return &modellerBuilder[T]{
		instance:   instance,
		contraints: make([]Contraint[T], 0),
	}
}
