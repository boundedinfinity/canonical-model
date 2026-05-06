package ioc

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

// https://en.wikipedia.org/wiki/List_of_FIFA_country_codes

var Countries = countries{
	Afghanistan: Country{
		Name: "Afghanistan",
		Code: "AFG",
	},
	Albania: Country{
		Name: "Albania",
		Code: "ALB",
	},
}

func init() {
	Countries.all = []Country{
		Countries.Afghanistan,
		Countries.Albania,
	}

	chain := []func(string) string{
		stringer.RemoveDiacritics[string],
		stringer.ToLower[string],
	}

	for _, country := range Countries.all {
		add := func(strs ...string) {
			for _, str := range strs {
				str = stringer.Chain(str, chain...)
				if _, ok := Countries.names[str]; !ok {
					Countries.names[str] = &country
				}
			}
		}

		add(country.Name, country.Code)
	}
}
