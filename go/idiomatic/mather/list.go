package mather

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Sum[N idiomatic.Number](numbers ...N) N {
	fn := func(_ int, a N, v N) N { return a + v }
	return slicer.Reduce(fn, 0, numbers...)
}

func SumFn[E any, N idiomatic.Number](fn func(E) N, numbers ...E) N {
	wrap := func(_ int, a N, e E) N { return a + fn(e) }
	return slicer.Reduce(wrap, 0, numbers...)
}

func Mean[N idiomatic.Number](numbers ...N) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := Sum(numbers...)
	mean := sum / N(l)

	return mean
}

func MeanFn[E any, N idiomatic.Number](fn func(E) N, numbers ...E) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := SumFn(fn, numbers...)
	mean := sum / N(l)

	return mean
}

func Median[N idiomatic.Number](numbers ...N) N {
	sorted := slicer.Sort(numbers...)
	i := len(sorted) / 2
	return numbers[i]
}

func MedianFn[E any, O idiomatic.Ordered, N idiomatic.Number](fn func(E) O, numbers ...E) E {
	sorted := slicer.SortBy(fn, numbers...)
	i := len(sorted) / 2
	return numbers[i]
}

func MinOf[N idiomatic.Number](numbers ...N) (N, bool) {
	var min N
	var ok bool
	size := len(numbers)

	if size > 0 {
		min = numbers[0]
		ok = true

		for i := 1; i < size; i++ {
			if min > numbers[i] {
				min = numbers[i]
			}
		}
	}

	return min, ok
}

func MaxOf[V idiomatic.Number](numbers ...V) V {
	fn := func(_ int, a V, v V) V {
		if v > a {
			return v
		} else {
			return a
		}
	}

	head, _ := slicer.Head(numbers...)
	return slicer.Reduce(fn, head, numbers...)
}
