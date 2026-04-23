package ider

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type Idable interface {
	GetId() Id
}

func ContainsFunc[T Idable](target T) func(T) bool {
	targetId := target.GetId()

	if targetId.IsZero() {
		return func(_ T) bool { return false }
	}

	return func(id T) bool {
		return !id.GetId().IsZero() && targetId == id.GetId()
	}
}

func Contains[T Idable](items []T, target T) bool {
	return slicer.ContainsFn(items, ContainsFunc(target))
}
