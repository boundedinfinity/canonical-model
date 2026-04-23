package multifactor_secret

type Kind string

type kinds struct {
	Unknown               Kind
	BiometricsFingerPrint Kind
	BiometricsFace        Kind
	BiometricsVoice       Kind
	QuickResponseCode     Kind
	OneTimePassword       Kind
}

var Kinds = kinds{
	Unknown:               "unknown",
	BiometricsFingerPrint: "biometrics-fingerprint",
	BiometricsFace:        "biometrics-face",
	BiometricsVoice:       "biometrics-voice",
	QuickResponseCode:     "quick-response-code",
	OneTimePassword:       "one-time-password",
}
