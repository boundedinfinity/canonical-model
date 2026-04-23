package perscription

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/idiomatic/currency"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/medical/amount"
	"github.com/boundedinfinity/canonical_model/idiomatic/medical/pharmaceutical"
	"github.com/boundedinfinity/canonical_model/idiomatic/person"
	"github.com/boundedinfinity/canonical_model/idiomatic/phone"
	"github.com/boundedinfinity/rfc3339date"
)

type PerscriptionModel struct {
	Id             ider.Id                            `json:"id"`
	PerscriptionId string                             `json:"perscription-id"`
	Patient        person.PersonModel                 `json:"patient"`
	Physician      person.PersonModel                 `json:"physician"`
	Pharmachy      business.BusinessModel             `json:"pharmachy"`
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
