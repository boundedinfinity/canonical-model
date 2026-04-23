package algorithm

type Kind string

type kinds struct {
	Unknown    Kind
	Rsa        Kind
	Dsa        Kind
	ElGamal    Kind
	Ecc        Kind
	Ed25519    Kind
	Curve25519 Kind
	Ecdsa      Kind
	Ecdh       Kind
	Brainpool  Kind
	Ed448      Kind
}

var Kinds = kinds{
	Unknown:    "unknown",
	Rsa:        "RSA",
	Dsa:        "DSA",
	ElGamal:    "ELGamal",
	Ecc:        "ECC",
	Ed25519:    "Ed25519",
	Curve25519: "Curve25519",
	Ecdsa:      "ECDSA",
	Ecdh:       "ECDH",
	Brainpool:  "Brainpool",
	Ed448:      "Ed448",
}
