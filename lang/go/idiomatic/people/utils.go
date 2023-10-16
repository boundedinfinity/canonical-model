package people

import (
	"fmt"

	"github.com/boundedinfinity/commons/slices"
	"github.com/boundedinfinity/optioner"
)

func toStrings[T any](ts []T) []string {
	return slices.Map(ts, func(t T) string { return fmt.Sprintf("%v", t) })
}

func sAppend[T any](slice []T, elems ...T) []T {
	if slice != nil {
		return append(slice, elems...)
	}

	return elems
}

func soAppend[T any](slice optioner.Option[[]T], elems ...T) optioner.Option[[]T] {
	if slice.Defined() && slice.Get() != nil {
		return optioner.Some(append(slice.Get(), elems...))
	}

	return optioner.Some(elems)
}
