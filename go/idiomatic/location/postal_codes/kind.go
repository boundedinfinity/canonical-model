package postal_codes

// https://en.wikipedia.org/wiki/Postal_code

type Kind string

type kinds struct {
	Unknown                     Kind
	ZoneImprovementPlan         Kind // United States
	CodigoDeEnderecamentoPostal Kind // Brazil
	CodiceDiAvviamentoPostale   Kind // Italy
	Eircode                     Kind // Ireland
	NumeroPostalDAcheminement   Kind // France
	PostalIndexNumber           Kind // India
	Postleitzahl                Kind // Germany
	PostalCode                  Kind // Canada
	PostCode                    Kind // Netherlands
	PostalIndex                 Kind // Belarus, Moldova, Russia, Ukraine, etc.
	PostoveSmerovacieCislo      Kind // Slovakia
	PostovniSmerovaciCislo      Kind // Czech Republic
}

var Kinds = kinds{
	Unknown:                     "unknown",
	ZoneImprovementPlan:         "zone-improvement-plan",
	CodigoDeEnderecamentoPostal: "codigo-de-enderecamento-postal",
	CodiceDiAvviamentoPostale:   "codice-di-avviamento-postale",
	Eircode:                     "eircode",
	NumeroPostalDAcheminement:   "numero-postal-dacheminement",
	PostalIndexNumber:           "postal-index-number",
	Postleitzahl:                "postleitzahl",
	PostalCode:                  "postal-code",
	PostCode:                    "post-code",
	PostalIndex:                 "postal-index",
	PostoveSmerovacieCislo:      "postove-smerovacie-cislo",
	PostovniSmerovaciCislo:      "postovni-smerovaci-cislo",
}
