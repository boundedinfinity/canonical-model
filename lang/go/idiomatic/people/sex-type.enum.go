package people

//go:generate enumer -path=./sex-type.enum.go

type SexType string

type sexTypes struct {
	Male   SexType `enum:"male"`
	Female SexType `enum:"female"`
}
