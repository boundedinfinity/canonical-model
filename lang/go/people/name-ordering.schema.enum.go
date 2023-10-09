package people

//go:generate enumer -path=./name-ordering.schema.enum.go

type NameOrdering string

const (
	GivenNameFamilyName NameOrdering = "Given name, Family name"
	FamilyNameGivenName NameOrdering = "Family name, Given name"
)
