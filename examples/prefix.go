package main

type Description string

type PrefixId string

// Prefix The prefix of a person's name.
//
// This may includes things like a title, certification, rank,
// honorific, etc...
type Prefix struct {
	Id            PrefixId
	Name          string
	Abbreviations []string
	Description   Description
}

var (
	Prefixes = []Prefix{
		{
			Id:            "https://www.boundedinfinity.com/schema/people/prefix/mister",
			Name:          "Mister",
			Abbreviations: []string{"Mr"},
			Description: Description(`A title used before a surname or full name to address or refer 
            to a man without a higher, honorific or professional title.`),
		},
		{
			Id:            "https://www.boundedinfinity.com/schema/people/prefix/mistress",
			Name:          "Mistress",
			Abbreviations: []string{"Miss"},
			Description: Description(`A title used before a surname or full name to address or refer 
            to a unmarried or young woman without a higher, honorific or 
            professional title.`),
		},
		{
			Id:            "https://www.boundedinfinity.com/schema/people/prefix/mrs",
			Name:          "Missus",
			Abbreviations: []string{"Mrs"},
			Description: Description(`A title used before a surname or full name to address or refer 
            to a married or widowed woman without a higher, honorific or 
            professional title.`),
		},
	}
)
