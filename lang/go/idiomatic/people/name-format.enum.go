package people

//go:generate enumer -path=./name-format.enum.go

type NameFormat string

type nameFormats struct {
	GivenNameFamilyName NameFormat
	FamilyNameGivenName NameFormat
}
