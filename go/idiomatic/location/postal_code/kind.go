package postal_code

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
	Unknown:                     "location.postal-code.unknown",
	ZoneImprovementPlan:         "location.postal-code.zone-improvement-plan",
	CodigoDeEnderecamentoPostal: "location.postal-code.codigo-de-enderecamento-postal",
	CodiceDiAvviamentoPostale:   "location.postal-code.codice-di-avviamento-postale",
	Eircode:                     "location.postal-code.eircode",
	NumeroPostalDAcheminement:   "location.postal-code.numero-postal-dacheminement",
	PostalIndexNumber:           "location.postal-code.postal-index-number",
	Postleitzahl:                "location.postal-code.postleitzahl",
	PostalCode:                  "location.postal-code.postal-code",
	PostCode:                    "location.postal-code.post-code",
	PostalIndex:                 "location.postal-code.postal-index",
	PostoveSmerovacieCislo:      "location.postal-code.postove-smerovacie-cislo",
	PostovniSmerovaciCislo:      "location.postal-code.postovni-smerovaci-cislo",
}
