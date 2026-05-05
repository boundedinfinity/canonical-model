package iso639

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type languages struct {
	all     []Language
	mapper  map[string]*Language
	English Language
	French  Language
	German  Language
	Spanish Language
}

func (this languages) FindOk(term string) ([]Language, bool) {
	term = stringer.Chain(term, stringer.RemoveDiacritics, stringer.ToLower)
	var found []Language
	var ok bool

	if lang, ok := this.mapper[term]; ok {
		found = append(found, *lang)
	}

	if !ok {
		for key, lang := range this.mapper {
			if stringer.Contains(key, term) {
				found = append(found, *lang)
				ok = true
			}
		}
	}

	return found, ok
}
