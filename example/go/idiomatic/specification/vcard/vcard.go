package vcard

import (
	"bytes"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
)

// https://en.wikipedia.org/wiki/VCard

func BuildVCard() *vCard40Builder {
	return &vCard40Builder{}
}

type vCard40Builder struct {
	contact contact.Contact
	address mailing_address.MailingAddress
}

func (t *vCard40Builder) Build() (VCard40, error) {
	var card VCard40

	if !t.address.IsZero() {
		items := []string{"ADR"}
		items = append(items, t.address.Lines...)
		items = append(items, t.address.City.String())
		items = append(items, t.address.State.String())
		items = append(items, t.address.Zip.String())
		card.ADR = stringer.Join(";", items...)
	}

	return VCard40{}, nil
}

func (t *vCard40Builder) Contact(contact contact.Contact) *vCard40Builder {
	t.contact = contact
	return t
}

func (t *vCard40Builder) ADR(address mailing_address.MailingAddress) *vCard40Builder {
	t.address = address
	return t
}

func (t *vCard40Builder) AGENT() *vCard40Builder {
	return t
}

func (t *vCard40Builder) ANNIVERSARY() *vCard40Builder {
	return t
}

func (t *vCard40Builder) BDAY() *vCard40Builder {
	return t
}

func (t *vCard40Builder) CALADURI() *vCard40Builder {
	return t
}

func (t *vCard40Builder) CALURI() *vCard40Builder {
	return t
}

func (t *vCard40Builder) CATEGORIES() *vCard40Builder {
	return t
}

func (t *vCard40Builder) CLASS() *vCard40Builder {
	return t
}

type VCard40 struct {
	ADR          string
	AGENT        string
	ANNIVERSARY  string
	BDAY         string
	CALADURI     string
	CALURI       string
	CATEGORIES   []string
	CLASS        string
	CLIENTPIDMAP string
	EMAIL        string
	FBURL        string
	FN           string
	GENDER       string
	GEO          string
	IMPP         string
	KEY          string
	KIND         string
	LABEL        string
	LANG         string
	LOGO         string
	MAILER       string
	MEMBER       string
	N            string
	NAME         string
	NICKNAME     string
	NOTE         string
	ORG          string
	PHOTO        string
	PRODID       string
	PROFILE      string
	RELATED      string
	REV          string
	ROLE         string
	SORT_STRING  string
	SOUND        string
	SOURCE       string
	TEL          string
	TITLE        string
	TZ           string
	UID          string
	URL          string
	VERSION      string
	XML          string
}

func (t VCard40) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	var err error

	add := func(s string) {
		buf.WriteString("BEGIN:VCARD")
	}

	add("BEGIN:VCARD")
	add("VERSION:4.0")

	add("END:VCARD")

	return buf.Bytes(), err
}
