package model

type model[T any] struct {
	instance   *T
	contraints []Contraint[T]
}

func Model[T any](instance *T) *model[T] {
	return &model[T]{
		instance:   instance,
		contraints: make([]Contraint[T], 0),
	}
}
