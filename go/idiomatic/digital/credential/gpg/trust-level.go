package gpg

type TrustLevel string

type trustLevels struct {
	Unknown  TrustLevel
	Ultimate TrustLevel
	Full     TrustLevel
	Marginal TrustLevel
}

var TrustLevels = trustLevels{
	Unknown:  "unknown",
	Ultimate: "ultimate",
	Full:     "full",
	Marginal: "marginal",
}
