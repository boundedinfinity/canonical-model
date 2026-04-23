package gpg

type Capability string

type capabilities struct {
	Unknown      Capability
	Sign         Capability
	Certify      Capability
	Encrypt      Capability
	Authenticate Capability
}

var Capabilities = capabilities{
	Unknown:      "unknown",
	Sign:         "S",
	Certify:      "C",
	Encrypt:      "E",
	Authenticate: "A",
}
