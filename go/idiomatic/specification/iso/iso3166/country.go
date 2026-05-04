package iso3166

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://en.wikipedia.org/wiki/ISO_3166-1#Codes
// ISO 3166-1 Name, Alpha-2 code, Alpha-3 code, Numeric code, Link to ISO 3166-2, Independent

type Country struct {
	Name         string        `json:"name,omitempty,omitzero"`
	Alpha2       string        `json:"alpha-2,omitempty,omitzero"`
	Alpha3       string        `json:"alpha-3,omitempty,omitzero"`
	Numeric      string        `json:"numeric,omitempty,omitzero"`
	Independent  string        `json:"independent,omitempty,omitzero"`
	Reference    string        `json:"reference,omitempty,omitzero"`
	Subdivisions []Subdivision `json:"subdivisions,omitempty,omitzero"`
}

type Subdivision struct {
	Country       *Country `json:"-"`
	Name          string   `json:"name,omitempty,omitzero"`
	Code          string   `json:"code,omitempty,omitzero"`
	Category      string   `json:"category,omitempty,omitzero"`
	Parent        string   `json:"parent,omitempty,omitzero"`
	Reference     string   `json:"reference,omitempty,omitzero"`
	AlternateCode string   `json:"alternate-code,omitempty,omitzero"`
	PreviousCode  string   `json:"previous-code,omitempty,omitzero"`
}

type countries struct {
	all                                    []Country
	countryNames                           map[string]*Country
	countryAlpha2s                         map[string]*Country
	countryAlpha3s                         map[string]*Country
	countryNumerics                        map[string]*Country
	subdivisionNames                       map[string]*Subdivision
	subdivisionCodes                       map[string]*Subdivision
	Afghanistan                            Country
	AlandIslands                           Country
	AmericanSamoa                          Country
	Anguilla                               Country
	Antarctica                             Country
	Albania                                Country
	Aruba                                  Country
	Algeria                                Country
	Andorra                                Country
	Angola                                 Country
	AntiguaAndBarbuda                      Country
	Argentina                              Country
	Armenia                                Country
	Australia                              Country
	Austria                                Country
	Azerbaijan                             Country
	Bermuda                                Country
	Bonaire                                Country
	BouvetIsland                           Country
	BritishIndianOceanTerritory            Country
	Bahamas                                Country
	Bahrain                                Country
	Bangladesh                             Country
	Barbados                               Country
	Belarus                                Country
	Belgium                                Country
	Belize                                 Country
	Benin                                  Country
	Bhutan                                 Country
	Bolivia                                Country
	BosniaAndHerzegovina                   Country
	Botswana                               Country
	Brazil                                 Country
	BruneiDarussalam                       Country
	Bulgaria                               Country
	BurkinaFaso                            Country
	Burundi                                Country
	Crotia                                 Country
	ChristmasIsland                        Country
	CocosIslands                           Country
	Columbia                               Country
	CaymanIslands                          Country
	CaboVerde                              Country
	Cambodia                               Country
	Cameroon                               Country
	Canada                                 Country
	CentralAfricanRepublic                 Country
	Chad                                   Country
	Chile                                  Country
	CookIslands                            Country
	Curacao                                Country
	China                                  Country
	Colombia                               Country
	Comoros                                Country
	Congo                                  Country
	DemocraticRepublicOfTheCongo           Country
	CostaRica                              Country
	CoteDivoire                            Country
	Croatia                                Country
	Cuba                                   Country
	Cyprus                                 Country
	Czechia                                Country
	Denmark                                Country
	Djibouti                               Country
	Dominica                               Country
	DominicanRepublic                      Country
	Ecuador                                Country
	Egypt                                  Country
	ElSalvador                             Country
	EquatorialGuinea                       Country
	Eritrea                                Country
	Estonia                                Country
	Eswatini                               Country
	Ethiopia                               Country
	Fiji                                   Country
	Finland                                Country
	France                                 Country
	FalklandIslands                        Country
	FaroeIslands                           Country
	FrenchGuiana                           Country
	FrenchPolynesia                        Country
	FrenchSouthernTerritories              Country
	Gabon                                  Country
	Gambia                                 Country
	Georgia                                Country
	Germany                                Country
	Ghana                                  Country
	Greece                                 Country
	Grenada                                Country
	Guatemala                              Country
	Guinea                                 Country
	GuineaBissau                           Country
	Guyana                                 Country
	Girbaltar                              Country
	Greenland                              Country
	Guadeloupe                             Country
	Guam                                   Country
	Guernsey                               Country
	Haiti                                  Country
	HolySee                                Country
	Honduras                               Country
	Hungary                                Country
	HeardIslandAndMcDonaldIslands          Country
	HongKong                               Country
	IsleOfMan                              Country
	Iceland                                Country
	India                                  Country
	Indonesia                              Country
	Iran                                   Country
	Iraq                                   Country
	Ireland                                Country
	Israel                                 Country
	Italy                                  Country
	Jamaica                                Country
	Japan                                  Country
	Jordan                                 Country
	Jersey                                 Country
	Kazakhstan                             Country
	Kenya                                  Country
	Kiribati                               Country
	Kuwait                                 Country
	Kyrgyzstan                             Country
	DemocraticPeoplesRepublicOfKorea       Country
	RepublicOfKorea                        Country
	LaoPeoplesDemocraticRepublic           Country
	Latvia                                 Country
	Lebanon                                Country
	Lesotho                                Country
	Liberia                                Country
	Libya                                  Country
	Liechtenstein                          Country
	Lithuania                              Country
	Luxembourg                             Country
	Madagascar                             Country
	Malawi                                 Country
	Malaysia                               Country
	Maldives                               Country
	Mali                                   Country
	Malta                                  Country
	MarshallIslands                        Country
	Mauritania                             Country
	Mauritius                              Country
	Mexico                                 Country
	Micronesia                             Country
	Moldova                                Country
	Monaco                                 Country
	Mongolia                               Country
	Montenegro                             Country
	Morocco                                Country
	Mozambique                             Country
	Myanmar                                Country
	Macao                                  Country
	Martinique                             Country
	Mayotte                                Country
	Montserrat                             Country
	Namibia                                Country
	Nauru                                  Country
	Nepal                                  Country
	Netherlands                            Country
	NewZealand                             Country
	Nicaragua                              Country
	Niger                                  Country
	Nigeria                                Country
	NorthMacedonia                         Country
	Norway                                 Country
	NewCaledonia                           Country
	Niue                                   Country
	NorfolkIsland                          Country
	NorthernMarianaIslands                 Country
	Oman                                   Country
	Pakistan                               Country
	Palau                                  Country
	PalestineState                         Country
	Panama                                 Country
	PapuaNewGuinea                         Country
	Paraguay                               Country
	Peru                                   Country
	Philippines                            Country
	Poland                                 Country
	Portugal                               Country
	Pitcairn                               Country
	PuertoRico                             Country
	Qatar                                  Country
	Reunion                                Country
	Romania                                Country
	Russia                                 Country
	Rwanda                                 Country
	SaintKittsAndNevis                     Country
	SaintLucia                             Country
	SaintVincentAndTheGrenadines           Country
	Samoa                                  Country
	SouthGeorgiaAndTheSouthSandwichIslands Country
	SvalbardAndJanMayen                    Country
	SanMarino                              Country
	SaoTomeAndPrincipe                     Country
	SaudiArabia                            Country
	Senegal                                Country
	Serbia                                 Country
	Seychelles                             Country
	SierraLeone                            Country
	Singapore                              Country
	Slovakia                               Country
	Slovenia                               Country
	SolomonIslands                         Country
	Somalia                                Country
	SouthAfrica                            Country
	SouthSudan                             Country
	Spain                                  Country
	SriLanka                               Country
	Sudan                                  Country
	Suriname                               Country
	Sweden                                 Country
	Switzerland                            Country
	Syria                                  Country
	SaintBarthelemy                        Country
	SaintHelena                            Country
	SaintMartin                            Country
	SaintPierreAndMiquelon                 Country
	SintMaarten                            Country
	Taiwan                                 Country
	Tajikistan                             Country
	Tanzania                               Country
	Thailand                               Country
	TimorLeste                             Country
	Togo                                   Country
	Tonga                                  Country
	TrinidadAndTobago                      Country
	Tunisia                                Country
	Turkey                                 Country
	Turkmenistan                           Country
	Tuvalu                                 Country
	Tokelau                                Country
	TurksAndCaicosIslands                  Country
	UnitedStatesMinorOutlyingIslands       Country
	Uganda                                 Country
	Ukraine                                Country
	UnitedArabEmirates                     Country
	UnitedKingdom                          Country
	UnitedStates                           Country
	Uruguay                                Country
	Uzbekistan                             Country
	Vanuatu                                Country
	Venezuela                              Country
	BritishVirginIslands                   Country
	UnitedStatesVirginIslands              Country
	WallisAndFutuna                        Country
	WesternSahara                          Country
	Vietnam                                Country
	Yemen                                  Country
	Zambia                                 Country
	Zimbabwe                               Country
}

func (this countries) All() []Country {
	return this.all[:]
}

func (this countries) GetOk(term string) (Country, bool) {
	lc := stringer.ToLower(term)
	var country *Country
	var ok bool

	if country, ok = this.countryNames[lc]; ok {
		return *country, true
	}

	if country, ok = this.countryAlpha2s[lc]; ok {
		return *country, true
	}

	if country, ok = this.countryAlpha3s[lc]; ok {
		return *country, true
	}

	if country, ok = this.countryNumerics[lc]; ok {
		return *country, true
	}

	return Country{}, false
}

func (this countries) FindOk(term string) ([]Country, bool) {
	lc := stringer.ToLower(term)
	var found []Country
	var country *Country
	var ok bool

	if country, ok = this.countryNames[lc]; ok {
		found = append(found, *country)
	}

	if !ok {
		if country, ok = this.countryAlpha2s[lc]; ok {
			found = append(found, *country)
		}
	}

	if !ok {
		if country, ok = this.countryAlpha3s[lc]; ok {
			found = append(found, *country)
		}
	}

	if !ok {
		if country, ok = this.countryNumerics[lc]; ok {
			found = append(found, *country)
		}
	}

	if !ok {
		var name string

		for name, country = range this.countryNames {
			if stringer.Contains(name, lc) {
				found = append(found, *country)
			}
		}

		for name, country = range this.countryAlpha2s {
			if stringer.Contains(name, lc) {
				found = append(found, *country)
			}
		}

		for name, country = range this.countryAlpha3s {
			if stringer.Contains(name, lc) {
				found = append(found, *country)
			}
		}

		for name, country = range this.countryNumerics {
			if stringer.Contains(name, lc) {
				found = append(found, *country)
			}
		}
	}

	return found, len(found) > 0
}

func init() {
	Countries.countryNames = make(map[string]*Country)
	Countries.countryAlpha2s = make(map[string]*Country)
	Countries.countryAlpha3s = make(map[string]*Country)
	Countries.countryNumerics = make(map[string]*Country)
	Countries.subdivisionNames = make(map[string]*Subdivision)
	Countries.subdivisionCodes = make(map[string]*Subdivision)

	for i := range Countries.all {
		Countries.countryNames[stringer.ToLower(Countries.all[i].Name)] = &Countries.all[i]
		Countries.countryAlpha2s[stringer.ToLower(Countries.all[i].Alpha2)] = &Countries.all[i]
		Countries.countryAlpha3s[stringer.ToLower(Countries.all[i].Alpha3)] = &Countries.all[i]

		numeric := Countries.all[i].Numeric
		Countries.countryNumerics[numeric] = &Countries.all[i]

		numeric2 := stringer.TrimPrefix(numeric, "0")
		if numeric != numeric2 {
			Countries.countryNumerics[numeric2] = &Countries.all[i]
		}

		for j := range Countries.all[i].Subdivisions {
			Countries.all[i].Subdivisions[j].Country = &Countries.all[i]
			subName := stringer.ToLower(Countries.all[i].Subdivisions[j].Name)
			subCode := stringer.ToLower(Countries.all[i].Subdivisions[j].Code)

			Countries.subdivisionNames[subName] = &Countries.all[i].Subdivisions[j]
			Countries.subdivisionCodes[subCode] = &Countries.all[i].Subdivisions[j]

			subCode2 := stringer.Replace(subCode, "", "-")
			if subCode != subCode2 {
				Countries.subdivisionCodes[subCode2] = &Countries.all[i].Subdivisions[j]
			}
		}
	}
}
