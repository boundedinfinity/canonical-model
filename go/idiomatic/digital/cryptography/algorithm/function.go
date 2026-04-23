package algorithm

type Function string

type functions struct {
	Unknown      Kind
	Encryption   Kind
	Signature    Kind
	KeyExchange  Kind
	Verification Kind
}

var Functions = functions{
	Unknown:      "unknown",
	Encryption:   "encryption",
	Signature:    "signature",
	Verification: "verification",
	KeyExchange:  "key-exchange",
}
