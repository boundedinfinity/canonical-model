package iso639

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

var Languages = languages{
	English: Language{
		Name: "English",
		Code: "en",
	},
	French: Language{
		Name: "French",
		Code: "fr",
	},
	German: Language{
		Name: "German",
		Code: "de",
	},
	Spanish: Language{
		Name: "Spanish",
		Code: "es",
	},
}

func init() {
	Languages.all = []Language{
		Languages.English,
		Languages.French,
		Languages.German,
		Languages.Spanish,
	}

	Languages.mapper = make(map[string]*Language)

	add := func(lang *Language, strs ...string) {
		for _, str := range strs {
			str = stringer.Chain(str, stringer.RemoveDiacritics, stringer.ToLower)
			if ok := Languages.mapper[str]; ok != nil {
				Languages.mapper[str] = lang
			}
		}
	}

	for i := range Languages.all {
		add(&Languages.all[i], Languages.all[i].Name, Languages.all[i].Code)
	}
}
