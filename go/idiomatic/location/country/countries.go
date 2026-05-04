package country

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/specification/iso/iso3166"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var Countries = countries{
	Afghanistan: Country{
		Name: "Afghanistan",
		Iso:  iso3166.Countries.Afghanistan,
	},
	AlandIslands: Country{
		Name: "Åland Islands",
		Iso:  iso3166.Countries.AlandIslands,
	},
	Albania: Country{
		Name: "Albania",
		Iso:  iso3166.Countries.Albania,
	},
	Algeria: Country{
		Name: "Algeria",
		Iso:  iso3166.Countries.Algeria,
	},
	AmericanSamoa: Country{
		Name: "American Samoa",
		Iso:  iso3166.Countries.AmericanSamoa,
	},
	Andorra: Country{
		Name: "Andorra",
		Iso:  iso3166.Countries.Andorra,
	},
	Angola: Country{
		Name: "Angola",
		Iso:  iso3166.Countries.Angola,
	},
	Anguilla: Country{
		Name: "Anguilla",
		Iso:  iso3166.Countries.Anguilla,
	},
	Antarctica: Country{
		Name: "Antarctica",
		Iso:  iso3166.Countries.Antarctica,
	},
	AntiguaAndBarbuda: Country{
		Name: "Antigua and Barbuda",
		Iso:  iso3166.Countries.AntiguaAndBarbuda,
	},
	Argentina: Country{
		Name: "Argentina",
		Iso:  iso3166.Countries.Argentina,
	},
	Armenia: Country{
		Name: "Armenia",
		Iso:  iso3166.Countries.Armenia,
	},
	Aruba: Country{
		Name: "Aruba",
		Iso:  iso3166.Countries.Aruba,
	},
	Australia: Country{
		Name: "Australia",
		Iso:  iso3166.Countries.Australia,
	},
	Austria: Country{
		Name: "Austria",
		Iso:  iso3166.Countries.Austria,
	},
	Azerbaijan: Country{
		Name: "Azerbaijan",
		Iso:  iso3166.Countries.Azerbaijan,
	},
	Bahamas: Country{
		Name: "Bahamas",
		Iso:  iso3166.Countries.Bahamas,
	},
	Bahrain: Country{
		Name: "Bahrain",
		Iso:  iso3166.Countries.Bahrain,
	},
	Bangladesh: Country{
		Name: "Bangladesh",
		Iso:  iso3166.Countries.Bangladesh,
	},
	Barbados: Country{
		Name: "Barbados",
		Iso:  iso3166.Countries.Barbados,
	},
	Belarus: Country{
		Name: "Belarus",
		Iso:  iso3166.Countries.Belarus,
	},
	Belgium: Country{
		Name: "Belgium",
		Iso:  iso3166.Countries.Belgium,
	},
	Belize: Country{
		Name: "Belize",
		Iso:  iso3166.Countries.Belize,
	},
	Benin: Country{
		Name: "Benin",
		Iso:  iso3166.Countries.Benin,
	},
	Bermuda: Country{
		Name: "Bermuda",
		Iso:  iso3166.Countries.Bermuda,
	},
	Bhutan: Country{
		Name: "Bhutan",
		Iso:  iso3166.Countries.Bhutan,
	},
	Bolivia: Country{
		Name: "Bolivia, Plurinational State of",
		Iso:  iso3166.Countries.Bolivia,
	},
	Bonaire: Country{
		Name: "Bonaire, Sint Eustatius and Saba",
		Iso:  iso3166.Countries.Bonaire,
	},
	BosniaAndHerzegovina: Country{
		Name: "Bosnia and Herzegovina",
		Iso:  iso3166.Countries.BosniaAndHerzegovina,
	},
	Botswana: Country{
		Name: "Botswana",
		Iso:  iso3166.Countries.Botswana,
	},
	BouvetIsland: Country{
		Name: "Bouvet Island",
		Iso:  iso3166.Countries.BouvetIsland,
	},
	Brazil: Country{
		Name: "Brazil",
		Iso:  iso3166.Countries.Brazil,
	},
	BritishIndianOceanTerritory: Country{
		Name: "British Indian Ocean Territory",
		Iso:  iso3166.Countries.BritishIndianOceanTerritory,
	},
	BruneiDarussalam: Country{
		Name: "Brunei Darussalam",
		Iso:  iso3166.Countries.BruneiDarussalam,
	},
	Bulgaria: Country{
		Name: "Bulgaria",
		Iso:  iso3166.Countries.Bulgaria,
	},
	BurkinaFaso: Country{
		Name: "Burkina Faso",
		Iso:  iso3166.Countries.BurkinaFaso,
	},
	Burundi: Country{
		Name: "Burundi",
		Iso:  iso3166.Countries.Burundi,
	},
	CaboVerde: Country{
		Name: "Cabo Verde",
		Iso:  iso3166.Countries.CaboVerde,
	},
	Colombia: Country{
		Name: "Colombia",
		Iso:  iso3166.Countries.Colombia,
	},
	Cameroon: Country{
		Name: "Cameroon",
		Iso:  iso3166.Countries.Cameroon,
	},
	Canada: Country{
		Name: "Canada",
		Iso:  iso3166.Countries.Canada,
	},
	CaymanIslands: Country{
		Name: "Cayman Islands",
		Iso:  iso3166.Countries.CaymanIslands,
	},
	CentralAfricanRepublic: Country{
		Name: "Central African Republic",
		Iso:  iso3166.Countries.CentralAfricanRepublic,
	},
	Chad: Country{
		Name: "Chad",
		Iso:  iso3166.Countries.Chad,
	},
	Chile: Country{
		Name: "Chile",
		Iso:  iso3166.Countries.Chile,
	},
	China: Country{
		Name: "China",
		Iso:  iso3166.Countries.China,
	},
	ChristmasIsland: Country{
		Name: "Christmas Island",
		Iso:  iso3166.Countries.ChristmasIsland,
	},
	CocosIslands: Country{
		Name: "Cocos (Keeling) Islands",
		Iso:  iso3166.Countries.CocosIslands,
	},
	Columbia: Country{
		Name: "Colombia",
		Iso:  iso3166.Countries.Colombia,
	},
	Comoros: Country{
		Name: "Comoros",
		Iso:  iso3166.Countries.Comoros,
	},
	Congo: Country{
		Name: "Congo",
		Iso:  iso3166.Countries.Congo,
	},
	DemocraticRepublicOfTheCongo: Country{
		Name: "Democratic Republic of the Congo",
		Iso:  iso3166.Countries.DemocraticRepublicOfTheCongo,
	},
	CookIslands: Country{
		Name: "Cook Islands",
		Iso:  iso3166.Countries.CookIslands,
	},
	CostaRica: Country{
		Name: "Costa Rica",
		Iso:  iso3166.Countries.CostaRica,
	},
	CoteDivoire: Country{
		Name: "Cote d'Ivoire",
		Iso:  iso3166.Countries.CoteDivoire,
	},
	Croatia: Country{
		Name: "Croatia",
		Iso:  iso3166.Countries.Croatia,
	},
	Cuba: Country{
		Name: "Cuba",
		Iso:  iso3166.Countries.Cuba,
	},
	Curacao: Country{
		Name: "Curacao",
		Iso:  iso3166.Countries.Curacao,
	},
	Cyprus: Country{
		Name: "Cyprus",
		Iso:  iso3166.Countries.Cyprus,
	},
	Czechia: Country{
		Name: "Czechia",
		Iso:  iso3166.Countries.Czechia,
	},
	Denmark: Country{
		Name: "Denmark",
		Iso:  iso3166.Countries.Denmark,
	},
	Djibouti: Country{
		Name: "Djibouti",
		Iso:  iso3166.Countries.Djibouti,
	},
	Dominica: Country{
		Name: "Dominica",
		Iso:  iso3166.Countries.Dominica,
	},
	DominicanRepublic: Country{
		Name: "Dominican Republic",
		Iso:  iso3166.Countries.DominicanRepublic,
	},
	Ecuador: Country{
		Name: "Ecuador",
		Iso:  iso3166.Countries.Ecuador,
	},
	Egypt: Country{
		Name: "Egypt",
		Iso:  iso3166.Countries.Egypt,
	},
	ElSalvador: Country{
		Name: "El Salvador",
		Iso:  iso3166.Countries.ElSalvador,
	},
	EquatorialGuinea: Country{
		Name: "Equatorial Guinea",
		Iso:  iso3166.Countries.EquatorialGuinea,
	},
	Eritrea: Country{
		Name: "Eritrea",
		Iso:  iso3166.Countries.Eritrea,
	},
	Estonia: Country{
		Name: "Estonia",
		Iso:  iso3166.Countries.Estonia,
	},
	Eswatini: Country{
		Name: "Eswatini",
		Iso:  iso3166.Countries.Eswatini,
	},
	Ethiopia: Country{
		Name: "Ethiopia",
		Iso:  iso3166.Countries.Ethiopia,
	},
	FalklandIslands: Country{
		Name: "Falkland Islands (Malvinas)",
		Iso:  iso3166.Countries.FalklandIslands,
	},
	FaroeIslands: Country{
		Name: "Faroe Islands",
		Iso:  iso3166.Countries.FaroeIslands,
	},
	Fiji: Country{
		Name: "Fiji",
		Iso:  iso3166.Countries.Fiji,
	},
	Finland: Country{
		Name: "Finland",
		Iso:  iso3166.Countries.Finland,
	},
	France: Country{
		Name: "France",
		Iso:  iso3166.Countries.France,
	},
	FrenchGuiana: Country{
		Name: "French Guiana",
		Iso:  iso3166.Countries.FrenchGuiana,
	},
	FrenchPolynesia: Country{
		Name: "French Polynesia",
		Iso:  iso3166.Countries.FrenchPolynesia,
	},
	FrenchSouthernTerritories: Country{
		Name: "French Southern Territories",
		Iso:  iso3166.Countries.FrenchSouthernTerritories,
	},
	Gabon: Country{
		Name: "Gabon",
		Iso:  iso3166.Countries.Gabon,
	},
	Gambia: Country{
		Name: "Gambia",
		Iso:  iso3166.Countries.Gambia,
	},
	Georgia: Country{
		Name: "Georgia",
		Iso:  iso3166.Countries.Georgia,
	},
	Germany: Country{
		Name: "Germany",
		Iso:  iso3166.Countries.Germany,
	},
	Ghana: Country{
		Name: "Ghana",
		Iso:  iso3166.Countries.Ghana,
	},
	Girbaltar: Country{
		Name: "Gibraltar",
		Iso:  iso3166.Countries.Girbaltar,
	},
	Greece: Country{
		Name: "Greece",
		Iso:  iso3166.Countries.Greece,
	},
	Greenland: Country{
		Name: "Greenland",
		Iso:  iso3166.Countries.Greenland,
	},
	Grenada: Country{
		Name: "Grenada",
		Iso:  iso3166.Countries.Grenada,
	},
	Guadeloupe: Country{
		Name: "Guadeloupe",
		Iso:  iso3166.Countries.Guadeloupe,
	},
	Guam: Country{
		Name: "Guam",
		Iso:  iso3166.Countries.Guam,
	},
	Guatemala: Country{
		Name: "Guatemala",
		Iso:  iso3166.Countries.Guatemala,
	},
	Guernsey: Country{
		Name: "Guernsey",
		Iso:  iso3166.Countries.Guernsey,
	},
	Guinea: Country{
		Name: "Guinea",
		Iso:  iso3166.Countries.Guinea,
	},
	GuineaBissau: Country{
		Name: "Guinea-Bissau",
		Iso:  iso3166.Countries.GuineaBissau,
	},
	Guyana: Country{
		Name: "Guyana",
		Iso:  iso3166.Countries.Guyana,
	},
	Haiti: Country{
		Name: "Haiti",
		Iso:  iso3166.Countries.Haiti,
	},
	HeardIslandAndMcDonaldIslands: Country{
		Name: "Heard Island and McDonald Islands",
		Iso:  iso3166.Countries.HeardIslandAndMcDonaldIslands,
	},
	HolySee: Country{
		Name: "Holy See",
		Iso:  iso3166.Countries.HolySee,
	},
	Honduras: Country{
		Name: "Honduras",
		Iso:  iso3166.Countries.Honduras,
	},
	HongKong: Country{
		Name: "Hong Kong",
		Iso:  iso3166.Countries.HongKong,
	},
	Hungary: Country{
		Name: "Hungary",
		Iso:  iso3166.Countries.Hungary,
	},
	Iceland: Country{
		Name: "Iceland",
		Iso:  iso3166.Countries.Iceland,
	},
	India: Country{
		Name: "India",
		Iso:  iso3166.Countries.India,
	},
	Indonesia: Country{
		Name: "Indonesia",
		Iso:  iso3166.Countries.Indonesia,
	},
	Iran: Country{
		Name: "Iran, Islamic Republic of",
		Iso:  iso3166.Countries.Iran,
	},
	Iraq: Country{
		Name: "Iraq",
		Iso:  iso3166.Countries.Iraq,
	},
	Ireland: Country{
		Name: "Ireland",
		Iso:  iso3166.Countries.Ireland,
	},
	IsleOfMan: Country{
		Name: "Isle of Man",
		Iso:  iso3166.Countries.IsleOfMan,
	},
	Israel: Country{
		Name: "Israel",
		Iso:  iso3166.Countries.Israel,
	},
	Italy: Country{
		Name: "Italy",
		Iso:  iso3166.Countries.Italy,
	},
	Jamaica: Country{
		Name: "Jamaica",
		Iso:  iso3166.Countries.Jamaica,
	},
	Japan: Country{
		Name: "Japan",
		Iso:  iso3166.Countries.Japan,
	},
	Jersey: Country{
		Name: "Jersey",
		Iso:  iso3166.Countries.Jersey,
	},
	Jordan: Country{
		Name: "Jordan",
		Iso:  iso3166.Countries.Jordan,
	},
	Kazakhstan: Country{
		Name: "Kazakhstan",
		Iso:  iso3166.Countries.Kazakhstan,
	},
	Kenya: Country{
		Name: "Kenya",
		Iso:  iso3166.Countries.Kenya,
	},
	Kiribati: Country{
		Name: "Kiribati",
		Iso:  iso3166.Countries.Kiribati,
	},
	DemocraticPeoplesRepublicOfKorea: Country{
		Name:    "Korea, Democratic People's Republic of",
		Aliases: []string{"North Korea"},
		Iso:     iso3166.Countries.DemocraticPeoplesRepublicOfKorea,
	},
	RepublicOfKorea: Country{
		Name:    "Korea, Republic of",
		Aliases: []string{"South Korea"},
		Iso:     iso3166.Countries.RepublicOfKorea,
	},
	Kuwait: Country{
		Name: "Kuwait",
		Iso:  iso3166.Countries.Kuwait,
	},
	Kyrgyzstan: Country{
		Name: "Kyrgyzstan",
		Iso:  iso3166.Countries.Kyrgyzstan,
	},
	LaoPeoplesDemocraticRepublic: Country{
		Name: "Lao People's Democratic Republic",
		Iso:  iso3166.Countries.LaoPeoplesDemocraticRepublic,
	},
	Latvia: Country{
		Name: "Latvia",
		Iso:  iso3166.Countries.Latvia,
	},
	Lebanon: Country{
		Name: "Lebanon",
		Iso:  iso3166.Countries.Lebanon,
	},
	Lesotho: Country{
		Name: "Lesotho",
		Iso:  iso3166.Countries.Lesotho,
	},
	Liberia: Country{
		Name: "Liberia",
		Iso:  iso3166.Countries.Liberia,
	},
	Libya: Country{
		Name: "Libya",
		Iso:  iso3166.Countries.Libya,
	},
	Liechtenstein: Country{
		Name: "Liechtenstein",
		Iso:  iso3166.Countries.Liechtenstein,
	},
	Lithuania: Country{
		Name: "Lithuania",
		Iso:  iso3166.Countries.Lithuania,
	},
	Luxembourg: Country{
		Name: "Luxembourg",
		Iso:  iso3166.Countries.Luxembourg,
	},
	Macao: Country{
		Name: "Macao",
		Iso:  iso3166.Countries.Macao,
	},
	Madagascar: Country{
		Name: "Madagascar",
		Iso:  iso3166.Countries.Madagascar,
	},
	Malawi: Country{
		Name: "Malawi",
		Iso:  iso3166.Countries.Malawi,
	},
	Malaysia: Country{
		Name: "Malaysia",
		Iso:  iso3166.Countries.Malaysia,
	},
	Maldives: Country{
		Name: "Maldives",
		Iso:  iso3166.Countries.Maldives,
	},
	Mali: Country{
		Name: "Mali",
		Iso:  iso3166.Countries.Mali,
	},
	Malta: Country{
		Name: "Malta",
		Iso:  iso3166.Countries.Malta,
	},
	MarshallIslands: Country{
		Name: "Marshall Islands",
		Iso:  iso3166.Countries.MarshallIslands,
	},
	Martinique: Country{
		Name: "Martinique",
		Iso:  iso3166.Countries.Martinique,
	},
	Mauritania: Country{
		Name: "Mauritania",
		Iso:  iso3166.Countries.Mauritania,
	},
	Mauritius: Country{
		Name: "Mauritius",
		Iso:  iso3166.Countries.Mauritius,
	},
	Mayotte: Country{
		Name: "Mayotte",
		Iso:  iso3166.Countries.Mayotte,
	},
	Mexico: Country{
		Name: "Mexico",
		Iso:  iso3166.Countries.Mexico,
	},
	Micronesia: Country{
		Name: "Micronesia, Federated States of",
		Iso:  iso3166.Countries.Micronesia,
	},
	Moldova: Country{
		Name: "Moldova, Republic of",
		Iso:  iso3166.Countries.Moldova,
	},
	Monaco: Country{
		Name: "Monaco",
		Iso:  iso3166.Countries.Monaco,
	},
	Mongolia: Country{
		Name: "Mongolia",
		Iso:  iso3166.Countries.Mongolia,
	},
	Montenegro: Country{
		Name: "Montenegro",
		Iso:  iso3166.Countries.Montenegro,
	},
	Montserrat: Country{
		Name: "Montserrat",
		Iso:  iso3166.Countries.Montserrat,
	},
	Morocco: Country{
		Name: "Morocco",
		Iso:  iso3166.Countries.Morocco,
	},
	Mozambique: Country{
		Name: "Mozambique",
		Iso:  iso3166.Countries.Mozambique,
	},
	Myanmar: Country{
		Name: "Myanmar",
		Iso:  iso3166.Countries.Myanmar,
	},
	Namibia: Country{
		Name: "Namibia",
		Iso:  iso3166.Countries.Namibia,
	},
	Nauru: Country{
		Name: "Nauru",
		Iso:  iso3166.Countries.Nauru,
	},
	Nepal: Country{
		Name: "Nepal",
		Iso:  iso3166.Countries.Nepal,
	},
	Netherlands: Country{
		Name: "Netherlands, Kingdom of the",
		Iso:  iso3166.Countries.Netherlands,
	},
	NewCaledonia: Country{
		Name: "New Caledonia",
		Iso:  iso3166.Countries.NewCaledonia,
	},
	NewZealand: Country{
		Name: "New Zealand",
		Iso:  iso3166.Countries.NewZealand,
	},
	Nicaragua: Country{
		Name: "Nicaragua",
		Iso:  iso3166.Countries.Nicaragua,
	},
	Niger: Country{
		Name: "Niger",
		Iso:  iso3166.Countries.Niger,
	},
	Nigeria: Country{
		Name: "Nigeria",
		Iso:  iso3166.Countries.Nigeria,
	},
	Niue: Country{
		Name: "Niue",
		Iso:  iso3166.Countries.Niue,
	},
	NorfolkIsland: Country{
		Name: "Norfolk Island",
		Iso:  iso3166.Countries.NorfolkIsland,
	},
	NorthMacedonia: Country{
		Name: "North Macedonia",
		Iso:  iso3166.Countries.NorthMacedonia,
	},
	NorthernMarianaIslands: Country{
		Name: "Northern Mariana Islands",
		Iso:  iso3166.Countries.NorthernMarianaIslands,
	},
	Norway: Country{
		Name: "Norway",
		Iso:  iso3166.Countries.Norway,
	},
	Oman: Country{
		Name: "Oman",
		Iso:  iso3166.Countries.Oman,
	},
	Pakistan: Country{
		Name: "Pakistan",
		Iso:  iso3166.Countries.Pakistan,
	},
	Palau: Country{
		Name: "Palau",
		Iso:  iso3166.Countries.Palau,
	},
	PalestineState: Country{
		Name: "Palestine, State of",
		Iso:  iso3166.Countries.PalestineState,
	},
	Panama: Country{
		Name: "Panama",
		Iso:  iso3166.Countries.Panama,
	},
	PapuaNewGuinea: Country{
		Name: "Papua New Guinea",
		Iso:  iso3166.Countries.PapuaNewGuinea,
	},
	Paraguay: Country{
		Name: "Paraguay",
		Iso:  iso3166.Countries.Paraguay,
	},
	Peru: Country{
		Name: "Peru",
		Iso:  iso3166.Countries.Peru,
	},
	Philippines: Country{
		Name: "Philippines",
		Iso:  iso3166.Countries.Philippines,
	},
	Pitcairn: Country{
		Name: "Pitcairn",
		Iso:  iso3166.Countries.Pitcairn,
	},
	Poland: Country{
		Name: "Poland",
		Iso:  iso3166.Countries.Poland,
	},
	Portugal: Country{
		Name: "Portugal",
		Iso:  iso3166.Countries.Portugal,
	},
	PuertoRico: Country{
		Name: "Puerto Rico",
		Iso:  iso3166.Countries.PuertoRico,
	},
	Qatar: Country{
		Name: "Qatar",
		Iso:  iso3166.Countries.Qatar,
	},
	Reunion: Country{
		Name: "Reunion",
		Iso:  iso3166.Countries.Reunion,
	},
	Romania: Country{
		Name: "Romania",
		Iso:  iso3166.Countries.Romania,
	},
	Russia: Country{
		Name: "Russian Federation",
		Iso:  iso3166.Countries.Russia,
	},
	Rwanda: Country{
		Name: "Rwanda",
		Iso:  iso3166.Countries.Rwanda,
	},
	SaintBarthelemy: Country{
		Name: "Saint Barthelemy",
		Iso:  iso3166.Countries.SaintBarthelemy,
	},
	SaintHelena: Country{
		Name: "Saint Helena, Ascension and Tristan da Cunha",
		Iso:  iso3166.Countries.SaintHelena,
	},
	SaintKittsAndNevis: Country{
		Name: "Saint Kitts and Nevis",
		Iso:  iso3166.Countries.SaintKittsAndNevis,
	},
	SaintLucia: Country{
		Name: "Saint Lucia",
		Iso:  iso3166.Countries.SaintLucia,
	},
	SaintMartin: Country{
		Name: "Saint Martin (French part)",
		Iso:  iso3166.Countries.SaintMartin,
	},
	SaintPierreAndMiquelon: Country{
		Name: "Saint Pierre and Miquelon",
		Iso:  iso3166.Countries.SaintPierreAndMiquelon,
	},
	SaintVincentAndTheGrenadines: Country{
		Name: "Saint Vincent and the Grenadines",
		Iso:  iso3166.Countries.SaintVincentAndTheGrenadines,
	},
	Samoa: Country{
		Name: "Samoa",
		Iso:  iso3166.Countries.Samoa,
	},
	SanMarino: Country{
		Name: "San Marino",
		Iso:  iso3166.Countries.SanMarino,
	},
	SaoTomeAndPrincipe: Country{
		Name: "Sao Tome and Principe",
		Iso:  iso3166.Countries.SaoTomeAndPrincipe,
	},
	SaudiArabia: Country{
		Name: "Saudi Arabia",
		Iso:  iso3166.Countries.SaudiArabia,
	},
	Senegal: Country{
		Name: "Senegal",
		Iso:  iso3166.Countries.Senegal,
	},
	Serbia: Country{
		Name: "Serbia",
		Iso:  iso3166.Countries.Serbia,
	},
	Seychelles: Country{
		Name: "Seychelles",
		Iso:  iso3166.Countries.Seychelles,
	},
	SierraLeone: Country{
		Name: "Sierra Leone",
		Iso:  iso3166.Countries.SierraLeone,
	},
	Singapore: Country{
		Name: "Singapore",
		Iso:  iso3166.Countries.Singapore,
	},
	SintMaarten: Country{
		Name: "Sint Maarten (Dutch part)",
		Iso:  iso3166.Countries.SintMaarten,
	},
	Slovakia: Country{
		Name: "Slovakia",
		Iso:  iso3166.Countries.Slovakia,
	},
	Slovenia: Country{
		Name: "Slovenia",
		Iso:  iso3166.Countries.Slovenia,
	},
	SolomonIslands: Country{
		Name: "Solomon Islands",
		Iso:  iso3166.Countries.SolomonIslands,
	},
	Somalia: Country{
		Name: "Somalia",
		Iso:  iso3166.Countries.Somalia,
	},
	SouthAfrica: Country{
		Name: "South Africa",
		Iso:  iso3166.Countries.SouthAfrica,
	},
	SouthGeorgiaAndTheSouthSandwichIslands: Country{
		Name: "South Georgia and the South Sandwich Islands",
		Iso:  iso3166.Countries.SouthGeorgiaAndTheSouthSandwichIslands,
	},
	SouthSudan: Country{
		Name: "South Sudan",
		Iso:  iso3166.Countries.SouthSudan,
	},
	Spain: Country{
		Name: "Spain",
		Iso:  iso3166.Countries.Spain,
	},
	SriLanka: Country{
		Name: "Sri Lanka",
		Iso:  iso3166.Countries.SriLanka,
	},
	Sudan: Country{
		Name: "Sudan",
		Iso:  iso3166.Countries.Sudan,
	},
	Suriname: Country{
		Name: "Suriname",
		Iso:  iso3166.Countries.Suriname,
	},
	SvalbardAndJanMayen: Country{
		Name: "Svalbard and Jan Mayen",
		Iso:  iso3166.Countries.SvalbardAndJanMayen,
	},
	Sweden: Country{
		Name: "Sweden",
		Iso:  iso3166.Countries.Sweden,
	},
	Switzerland: Country{
		Name: "Switzerland",
		Iso:  iso3166.Countries.Switzerland,
	},
	Syria: Country{
		Name: "Syrian Arab Republic",
		Iso:  iso3166.Countries.Syria,
	},
	Taiwan: Country{
		Name: "Taiwan, Province of China",
		Iso:  iso3166.Countries.Taiwan,
	},
	Tajikistan: Country{
		Name: "Tajikistan",
		Iso:  iso3166.Countries.Tajikistan,
	},
	Tanzania: Country{
		Name: "Tanzania, United Republic of",
		Iso:  iso3166.Countries.Tanzania,
	},
	Thailand: Country{
		Name: "Thailand",
		Iso:  iso3166.Countries.Thailand,
	},
	TimorLeste: Country{
		Name: "Timor-Leste",
		Iso:  iso3166.Countries.TimorLeste,
	},
	Togo: Country{
		Name: "Togo",
		Iso:  iso3166.Countries.Togo,
	},
	Tokelau: Country{
		Name: "Tokelau",
		Iso:  iso3166.Countries.Tokelau,
	},
	Tonga: Country{
		Name: "Tonga",
		Iso:  iso3166.Countries.Tonga,
	},
	TrinidadAndTobago: Country{
		Name: "Trinidad and Tobago",
		Iso:  iso3166.Countries.TrinidadAndTobago,
	},
	Tunisia: Country{
		Name: "Tunisia",
		Iso:  iso3166.Countries.Tunisia,
	},
	Turkey: Country{
		Name: "Turkiye",
		Iso:  iso3166.Countries.Turkey,
	},
	Turkmenistan: Country{
		Name: "Turkmenistan",
		Iso:  iso3166.Countries.Turkmenistan,
	},
	TurksAndCaicosIslands: Country{
		Name: "Turks and Caicos Islands",
		Iso:  iso3166.Countries.TurksAndCaicosIslands,
	},
	Tuvalu: Country{
		Name: "Tuvalu",
		Iso:  iso3166.Countries.Tuvalu,
	},
	Uganda: Country{
		Name: "Uganda",
		Iso:  iso3166.Countries.Uganda,
	},
	Ukraine: Country{
		Name: "Ukraine",
		Iso:  iso3166.Countries.Ukraine,
	},
	UnitedArabEmirates: Country{
		Name: "United Arab Emirates",
		Iso:  iso3166.Countries.UnitedArabEmirates,
	},
	UnitedKingdom: Country{
		Name: "United Kingdom of Great Britain and Northern Ireland",
		Iso:  iso3166.Countries.UnitedKingdom,
	},
	UnitedStates: Country{
		Name: "United States of America",
		Iso:  iso3166.Countries.UnitedStates,
	},
	UnitedStatesMinorOutlyingIslands: Country{
		Name: "United States Minor Outlying Islands",
		Iso:  iso3166.Countries.UnitedStatesMinorOutlyingIslands,
	},
	Uruguay: Country{
		Name: "Uruguay",
		Iso:  iso3166.Countries.Uruguay,
	},
	Uzbekistan: Country{
		Name: "Uzbekistan",
		Iso:  iso3166.Countries.Uzbekistan,
	},
	Vanuatu: Country{
		Name: "Vanuatu",
		Iso:  iso3166.Countries.Vanuatu,
	},
	Venezuela: Country{
		Name: "Venezuela, Bolivarian Republic of",
		Iso:  iso3166.Countries.Venezuela,
	},
	Vietnam: Country{
		Name: "Viet Nam",
		Iso:  iso3166.Countries.Vietnam,
	},
	BritishVirginIslands: Country{
		Name:    "Virgin Islands (British)",
		Aliases: []string{"BVI"},
		Iso:     iso3166.Countries.BritishVirginIslands,
	},
	UnitedStatesVirginIslands: Country{
		Name: "Virgin Islands (U.S.)",
		Iso:  iso3166.Countries.UnitedStatesVirginIslands,
	},
	WallisAndFutuna: Country{
		Name: "Wallis and Futuna",
		Iso:  iso3166.Countries.WallisAndFutuna,
	},
	WesternSahara: Country{
		Name: "Western Sahara",
		Iso:  iso3166.Countries.WesternSahara,
	},
	Yemen: Country{
		Name: "Yemen",
		Iso:  iso3166.Countries.Yemen,
	},
	Zambia: Country{
		Name: "Zambia",
		Iso:  iso3166.Countries.Zambia,
	},
	Zimbabwe: Country{
		Name: "Zimbabwe",
		Iso:  iso3166.Countries.Zimbabwe,
	},
}

func init() {
	Countries.all = []Country{
		Countries.Afghanistan,
		Countries.AlandIslands,
		Countries.Albania,
		Countries.Algeria,
		Countries.AmericanSamoa,
		Countries.Andorra,
		Countries.Angola,
		Countries.Anguilla,
		Countries.Antarctica,
		Countries.AntiguaAndBarbuda,
		Countries.Argentina,
		Countries.Armenia,
		Countries.Aruba,
		Countries.Australia,
		Countries.Austria,
		Countries.Azerbaijan,
		Countries.Bahamas,
		Countries.Bahrain,
		Countries.Bangladesh,
		Countries.Barbados,
		Countries.Belarus,
		Countries.Belgium,
		Countries.Belize,
		Countries.Benin,
		Countries.Bermuda,
		Countries.Bhutan,
		Countries.Bolivia,
		Countries.Bonaire,
		Countries.BosniaAndHerzegovina,
		Countries.Botswana,
		Countries.BouvetIsland,
		Countries.Brazil,
		Countries.BritishIndianOceanTerritory,
		Countries.BruneiDarussalam,
		Countries.Bulgaria,
		Countries.BurkinaFaso,
		Countries.Burundi,
		Countries.CaboVerde,
		Countries.Colombia,
		Countries.Cameroon,
		Countries.Canada,
		Countries.CaymanIslands,
		Countries.CentralAfricanRepublic,
		Countries.Chad,
		Countries.Chile,
		Countries.China,
		Countries.ChristmasIsland,
		Countries.CocosIslands,
		Countries.Columbia,
		Countries.Comoros,
		Countries.Congo,
		Countries.DemocraticRepublicOfTheCongo,
		Countries.CookIslands,
		Countries.CostaRica,
		Countries.CoteDivoire,
		Countries.Croatia,
		Countries.Cuba,
		Countries.Curacao,
		Countries.Cyprus,
		Countries.Czechia,
		Countries.Denmark,
		Countries.Djibouti,
		Countries.Dominica,
		Countries.DominicanRepublic,
		Countries.Ecuador,
		Countries.Egypt,
		Countries.ElSalvador,
		Countries.EquatorialGuinea,
		Countries.Eritrea,
		Countries.Estonia,
		Countries.Eswatini,
		Countries.Ethiopia,
		Countries.FalklandIslands,
		Countries.FaroeIslands,
		Countries.Fiji,
		Countries.Finland,
		Countries.France,
		Countries.FrenchGuiana,
		Countries.FrenchPolynesia,
		Countries.FrenchSouthernTerritories,
		Countries.Gabon,
		Countries.Gambia,
		Countries.Georgia,
		Countries.Germany,
		Countries.Ghana,
		Countries.Girbaltar,
		Countries.Greece,
		Countries.Greenland,
		Countries.Grenada,
		Countries.Guadeloupe,
		Countries.Guam,
		Countries.Guatemala,
		Countries.Guernsey,
		Countries.Guinea,
		Countries.GuineaBissau,
		Countries.Guyana,
		Countries.Haiti,
		Countries.HeardIslandAndMcDonaldIslands,
		Countries.HolySee,
		Countries.Honduras,
		Countries.HongKong,
		Countries.Hungary,
		Countries.Iceland,
		Countries.India,
		Countries.Indonesia,
		Countries.Iran,
		Countries.Iraq,
		Countries.Ireland,
		Countries.IsleOfMan,
		Countries.Israel,
		Countries.Italy,
		Countries.Jamaica,
		Countries.Japan,
		Countries.Jersey,
		Countries.Jordan,
		Countries.Kazakhstan,
		Countries.Kenya,
		Countries.Kiribati,
		Countries.DemocraticPeoplesRepublicOfKorea,
		Countries.RepublicOfKorea,
		Countries.Kuwait,
		Countries.Kyrgyzstan,
		Countries.LaoPeoplesDemocraticRepublic,
		Countries.Latvia,
		Countries.Lebanon,
		Countries.Lesotho,
		Countries.Liberia,
		Countries.Libya,
		Countries.Liechtenstein,
		Countries.Lithuania,
		Countries.Luxembourg,
		Countries.Macao,
		Countries.Madagascar,
		Countries.Malawi,
		Countries.Malaysia,
		Countries.Maldives,
		Countries.Mali,
		Countries.Malta,
		Countries.MarshallIslands,
		Countries.Martinique,
		Countries.Mauritania,
		Countries.Mauritius,
		Countries.Mayotte,
		Countries.Mexico,
		Countries.Micronesia,
		Countries.Moldova,
		Countries.Monaco,
		Countries.Mongolia,
		Countries.Montenegro,
		Countries.Montserrat,
		Countries.Morocco,
		Countries.Mozambique,
		Countries.Myanmar,
		Countries.Namibia,
		Countries.Nauru,
		Countries.Nepal,
		Countries.Netherlands,
		Countries.NewCaledonia,
		Countries.NewZealand,
		Countries.Nicaragua,
		Countries.Niger,
		Countries.Nigeria,
		Countries.Niue,
		Countries.NorfolkIsland,
		Countries.NorthMacedonia,
		Countries.NorthernMarianaIslands,
		Countries.Norway,
		Countries.Oman,
		Countries.Pakistan,
		Countries.Palau,
		Countries.PalestineState,
		Countries.Panama,
		Countries.PapuaNewGuinea,
		Countries.Paraguay,
		Countries.Peru,
		Countries.Philippines,
		Countries.Pitcairn,
		Countries.Poland,
		Countries.Portugal,
		Countries.PuertoRico,
		Countries.Qatar,
		Countries.Reunion,
		Countries.Romania,
		Countries.Russia,
		Countries.Rwanda,
		Countries.SaintBarthelemy,
		Countries.SaintHelena,
		Countries.SaintKittsAndNevis,
		Countries.SaintLucia,
		Countries.SaintMartin,
		Countries.SaintPierreAndMiquelon,
		Countries.SaintVincentAndTheGrenadines,
		Countries.Samoa,
		Countries.SanMarino,
		Countries.SaoTomeAndPrincipe,
		Countries.SaudiArabia,
		Countries.Senegal,
		Countries.Serbia,
		Countries.Seychelles,
		Countries.SierraLeone,
		Countries.Singapore,
		Countries.SintMaarten,
		Countries.Slovakia,
		Countries.Slovenia,
		Countries.SolomonIslands,
		Countries.Somalia,
		Countries.SouthAfrica,
		Countries.SouthGeorgiaAndTheSouthSandwichIslands,
		Countries.SouthSudan,
		Countries.Spain,
		Countries.SriLanka,
		Countries.Sudan,
		Countries.Suriname,
		Countries.SvalbardAndJanMayen,
		Countries.Sweden,
		Countries.Switzerland,
		Countries.Syria,
		Countries.Taiwan,
		Countries.Tajikistan,
		Countries.Tanzania,
		Countries.Thailand,
		Countries.TimorLeste,
		Countries.Togo,
		Countries.Tokelau,
		Countries.Tonga,
		Countries.TrinidadAndTobago,
		Countries.Tunisia,
		Countries.Turkey,
		Countries.Turkmenistan,
		Countries.TurksAndCaicosIslands,
		Countries.Tuvalu,
		Countries.Uganda,
		Countries.Ukraine,
		Countries.UnitedArabEmirates,
		Countries.UnitedKingdom,
		Countries.UnitedStates,
		Countries.UnitedStatesMinorOutlyingIslands,
		Countries.Uruguay,
		Countries.Uzbekistan,
		Countries.Vanuatu,
		Countries.Venezuela,
		Countries.Vietnam,
		Countries.BritishVirginIslands,
		Countries.UnitedStatesVirginIslands,
		Countries.WallisAndFutuna,
		Countries.WesternSahara,
		Countries.Yemen,
		Countries.Zambia,
		Countries.Zimbabwe,
	}

	Countries.names = make(map[string]*Country)

	for _, country := range Countries.all {
		name := stringer.ToLower(country.Name)
		Countries.names[name] = &country

		for _, alias := range country.Aliases {
			alias = stringer.ToLower(alias)
			if name != alias {
				Countries.names[alias] = &country
			}
		}
	}
}
