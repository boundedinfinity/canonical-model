package audit

//go:generate enumer -path=./status.enum.go

type Status string

type statues struct {
	Active       Status `enum:"active"`
	Inactive     Status `enum:"inactive"`
	Discontinued Status `enum:"discontinued"`
}
