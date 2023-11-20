package modeller

type modellerBuilder[T any] struct {
	instance   *T
	contraints []Contraint
}

func Modeller[T any](instance *T) *modellerBuilder[T] {
	return &modellerBuilder[T]{
		instance:   instance,
		contraints: make([]Contraint, 0),
	}
}

type StringConstraint func(string) Contraint

func StructModeller[T any](instance *T) *modellerBuilder[T] {
	return &modellerBuilder[T]{
		instance: instance,
	}
}
