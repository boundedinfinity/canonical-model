package people

//go:generate enumer -path=./prefix-format.enum.go

type PrefixFormat string

type prefixFormats struct {
	Full         PrefixFormat
	Abbreviation PrefixFormat
}
