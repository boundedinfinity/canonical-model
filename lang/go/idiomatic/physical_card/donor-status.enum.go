package physical_card

//go:generate enumer -path=./donor-status.enum.go

type DonorStatus string

type donorStatuses struct {
	IsDonor  DonorStatus `enum:"is-donor"`
	NotDonor DonorStatus `enum:"not-donor"`
}
