package people

//go:generate enumer -path=./suffix-format.enum.go

type SuffixFormat string

type suffixFormats struct {
	Full         SuffixFormat
	Abbreviation SuffixFormat
}
