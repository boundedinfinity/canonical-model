package multifactor_secret

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Secret = &BiometricsVoiceModel{}

type BiometricsVoiceModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this BiometricsVoiceModel) GetName() string {
	return this.Name
}

func (_ BiometricsVoiceModel) GetKind() Kind {
	return Kinds.BiometricsVoice
}
