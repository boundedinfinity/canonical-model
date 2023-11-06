package banking

//go:generate enumer -path=./account-number-format.enum.go

type AccountNumberFormat string

type nameFormats struct {
	NoSpace     AccountNumberFormat `enum:"no-space"`
	SpaceAfter3 AccountNumberFormat `enum:"space-after-3"`
}
