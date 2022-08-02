package main

//go:generate enumer -path=./name-ordering.schema.go

type NameOrdering string

const (
	GivenNameFamilyName = "Given name, Family name"
	FamilyNameGivenName = "Family name, Given name"
)
