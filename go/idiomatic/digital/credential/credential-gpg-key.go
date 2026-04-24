package credential

// https://www.gnupg.org/

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/cryptography/algorithm"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/specifications/rfc/rfc3339date"
)

var _ Credential = &GpgKeyModel{}

type GpgKeyModel struct {
	Id             ider.Id                     `json:"id"`
	Name           string                      `json:"name"`
	GpgId          string                      `json:"gpg-id"`
	PublicKey      string                      `json:"public-key"`
	PrivateKey     string                      `json:"private-key"`
	PassPhrase     string                      `json:"pass-phrase"`
	Algorithm      algorithm.AlgorithmModel    `json:"algorithm"`
	ExpirationDate rfc3339date.Rfc3339DateTime `json:"expiration-date"`
	SubKeys        []*GpgKeyModel              `json:"sub-keys"`
	Parent         *GpgKeyModel                `json:"parent"`
}

func (this GpgKeyModel) GetName() string {
	return this.Name
}

func (_ GpgKeyModel) GetKind() Kind {
	return Kinds.GpgKey
}
