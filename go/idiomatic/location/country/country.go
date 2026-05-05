package country

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/specification/iso/iso3166"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Country struct {
	Name    string          `json:"name"`
	Aliases []string        `json:"aliases"`
	Iso     iso3166.Country `json:"iso"`
}

type countries struct {
	all                                    []Country
	names                                  map[string]*Country
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

func (this countries) FindOk(term string) ([]Country, bool) {
	term = stringer.Chain(term, stringer.RemoveDiacritics, stringer.ToLower)
	var found []Country
	var country *Country
	var ok bool

	if country, ok = this.names[term]; ok {
		found = append(found, *country)
	}

	if !ok {
		for name, country := range this.names {
			if stringer.Contains(name, term) {
				found = append(found, *country)
			}
		}
	}

	return found, ok
}
