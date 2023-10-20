package people

//go:generate enumer -path=./name-format.enum.go

type NameFormat string

type nameFormats struct {
	GivenNameFamilyName NameFormat `enum:"given-name-family-name"`
	FamilyNameGivenName NameFormat `enum:"family-name-given-name"`
}
