package multifactor_secret

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Secret = &BiometricsFingerprintModel{}

type BiometricsFingerprintModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this BiometricsFingerprintModel) GetName() string {
	return this.Name
}

func (_ BiometricsFingerprintModel) GetKind() Kind {
	return Kinds.BiometricsFingerPrint
}
