package mailing_address

type Zip string

func (t Zip) String() string {
	return string(t)
}
