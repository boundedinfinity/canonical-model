package credential

type Kind string

type kinds struct {
	Unknown          Kind
	EmailPassword    Kind
	UsernamePassword Kind
	SshKey           Kind
	Passkey          Kind
	GpgKey           Kind
	ApiKey           Kind
	SslCert          Kind
	TlsCert          Kind
	Shared           Kind
	Delegated        Kind
}

var Kinds = kinds{
	Unknown:          "unknown",
	EmailPassword:    "email-password",
	UsernamePassword: "username-password",
	SshKey:           "ssh-key",
	Passkey:          "passkey",
	GpgKey:           "gpg-key",
	ApiKey:           "api-key",
	SslCert:          "ssl-cert",
	TlsCert:          "tls-cert",
	Shared:           "shared",
	Delegated:        "delegated",
}
