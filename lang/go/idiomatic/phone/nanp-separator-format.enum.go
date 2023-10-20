package phone

//go:generate enumer -path=./nanp-separator-format.enum.go

type NapaPhoneSeparatorFormat string

type napaPhoneSeparatorFormats struct {
	DotSeparated         NapaPhoneSeparatorFormat
	DashSeparated        NapaPhoneSeparatorFormat
	ParenthesesAndDashes NapaPhoneSeparatorFormat
	None                 NapaPhoneSeparatorFormat
}
