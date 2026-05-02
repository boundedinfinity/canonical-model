package label

import "fmt"

var Generator = generator{}

type generator struct{}

func (_ generator) Months() Labels {
	labels := make(Labels, 12)

	for i := 1; i <= 12; i++ {
		labels[i] = Label{
			Name: fmt.Sprintf("%02d", i),
		}
	}

	return labels
}

func (_ generator) Years(start, end int) Labels {
	labels := make(Labels, end-start)

	for year := start; year <= end; year++ {
		labels[year] = Label{
			Name: fmt.Sprintf("%04d", year),
		}
	}

	return labels
}

func (_ generator) YearMonths(year int) Labels {
	labels := make(Labels, 12)

	for i := 1; i <= 12; i++ {
		labels[i] = Label{
			Name: fmt.Sprintf("%04d.%02d", year, i),
		}
	}

	return labels
}
