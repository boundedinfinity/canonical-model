package model

type Contraint[T any] func(t T) error

type Validator interface {
	Validate(groups ...string) error
}
