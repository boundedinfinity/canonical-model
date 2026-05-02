package perscription

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/currency"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/medical/amount"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/medical/pharmaceutical"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
	"github.com/boundedinfinity/rfc3339date"
)

type PerscriptionModel struct {
	Id             ider.Id                            `json:"id"`
	PerscriptionId string                             `json:"perscription-id"`
	Patient        person.Person                      `json:"patient"`
	Physician      person.Person                      `json:"physician"`
	Pharmachy      business.Business                  `json:"pharmachy"`
	Directions     string                             `json:"directions"`
	Phramaceutical pharmaceutical.PhramaceuticalModel `json:"phramaceutical"`
	Count          int                                `json:"count"`
	Phone          phone.PhoneModel                   `json:"phone"`
	PickedUpDate   rfc3339date.Rfc3339Date            `json:"picked-up-date"`
	PromisedDate   rfc3339date.Rfc3339Date            `json:"promised-date"`
	RefillCount    int                                `json:"refill-count"`
	RefillBy       rfc3339date.Rfc3339Date            `json:"refill-by"`
	FormFactor     FormFactor                         `json:"form-factor"`
	Amount         amount.Amount                      `json:"amount"`
	Price          currency.Amount                    `json:"price"`
}
