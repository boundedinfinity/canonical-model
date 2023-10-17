package phone

//go:generate enumer -path=./nanp-phone-format.enum.go

type NapaPhoneFormat string

type napaPhoneFormats struct {
	Full   NapaPhoneFormat
	Common NapaPhoneFormat
}
