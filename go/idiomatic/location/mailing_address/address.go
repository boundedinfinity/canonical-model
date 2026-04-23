package mailing_address

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type MailingAddressKind string

const (
	Unknown  MailingAddressKind = "unknown"
	Work     MailingAddressKind = "work"
	Personal MailingAddressKind = "personal"
)

type Address struct {
	Id    ider.Id            `json:"id"`
	Kind  MailingAddressKind `json:"kind"`
	Lines []string           `json:"lines"`
	City  string             `json:"city"`
	State string             `json:"state"`
	Zip   string             `json:"zip"`
}
