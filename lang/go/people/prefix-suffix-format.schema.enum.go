package people

//go:generate enumer -path=./prefix-suffix-format.schema.enum.go

type PrefixSuffixFormat string

const (
	Full         PrefixSuffixFormat = "Full"
	Abbreviation PrefixSuffixFormat = "Abbreviation"
)
