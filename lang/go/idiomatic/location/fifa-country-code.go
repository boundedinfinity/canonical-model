package location

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://en.wikipedia.org/wiki/List_of_FIFA_country_codes

type FifaCountryCode struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

func init() {
	// for _, line := range stringer.Split(fifaData, "\n") {
	// 	line = stringer.TrimSpace(line)

	// 	if stringer.IsEmpty(line) {
	// 		continue
	// 	}

	// 	fields := slicer.Map(func(s string) string {
	// 		return stringer.Trim(s, `"`)
	// 	}, stringer.Split(line, "|")...)

	// 	if len(fields) != 2 {
	// 		continue
	// 	}

	// 	country := FifaCountryCode{
	// 		Name: fields[0],
	// 		Code: fields[1],
	// 	}

	// 	FifaCountryCodes.codeToCountry[country.Code] = country
	// 	FifaCountryCodes.nameToCountry[country.Name] = country
	// 	FifaCountryCodes.Values = append(FifaCountryCodes.Values, country)
	// }
}

type fifaCountryCodes struct {
	Values        []FifaCountryCode
	nameToCountry map[string]FifaCountryCode
	codeToCountry map[string]FifaCountryCode
}

func (t fifaCountryCodes) Lookup(s string) (FifaCountryCode, bool) {
	var country FifaCountryCode
	var found bool

	country, found = t.ByName(s)

	if !found {
		country, found = t.ByCode(s)
	}

	return country, found
}

func (t fifaCountryCodes) ByName(s string) (FifaCountryCode, bool) {
	country, ok := t.codeToCountry[stringer.Lowercase(s)]
	return country, ok
}

func (t fifaCountryCodes) ByCode(s string) (FifaCountryCode, bool) {
	country, ok := t.codeToCountry[stringer.Lowercase(s)]
	return country, ok
}

var (
	FifaCountryCodes = fifaCountryCodes{}
	fifaData         = `
        Afghanistan|AFG
        Albania|ALB
        Algeria|ALG
        American Samoa|ASA
        Andorra|AND
        Angola|ANG
        Anguilla|AIA
        Antigua and Barbuda|ATG
        Argentina|ARG
        Armenia|ARM
        Aruba|ARU
        Australia|AUS
        Austria|AUT
        Azerbaijan|AZE
        Bahamas|BAH
        Bahrain|BHR
        Bangladesh|BAN
        Barbados|BRB
        Belarus|BLR
        Belgium|BEL
        Belize|BLZ
        Benin|BEN
        Bermuda|BER
        Bhutan|BHU
        Bolivia|BOL
        Bosnia and Herzegovina|BIH
        Botswana|BOT
        Brazil|BRA
        British Virgin Islands|VGB
        Brunei|BRU
        Bulgaria|BUL
        Burkina Faso|BFA
        Burundi|BDI
        Cambodia|CAM
        Cameroon|CMR
        Canada|CAN
        Cape Verde|CPV
        Cayman Islands|CAY
        Central African Republic|CTA
        Chad|CHA
        Chile|CHI
        China|CHN
        Chinese Taipei|TPE
        Colombia|COL
        Comoros|COM
        Congo|CGO
        Cook Islands|COK
        Costa Rica|CRC
        Croatia|CRO
        Cuba|CUB
        Curaçao|CUW
        Cyprus|CYP
        Czech Republic|CZE
        Denmark|DEN
        Djibouti|DJI
        Dominica|DMA
        Dominican Republic|DOM
        DR Congo|COD
        Ecuador|ECU
        Egypt|EGY
        El Salvador|SLV
        England|ENG
        Equatorial Guinea|EQG
        Eritrea|ERI
        Estonia|EST
        Eswatini|SWZ
        Ethiopia|ETH
        Faroe Islands|FRO
        Fiji|FIJ
        Finland|FIN
        France|FRA
        Gabon|GAB
        Gambia|GAM
        Georgia|GEO
        Germany|GER
        Ghana|GHA
        Gibraltar|GIB
        Greece|GRE
        Grenada|GRN
        Guam|GUM
        Guatemala|GUA
        Guinea|GUI
        Guinea-Bissau|GNB
        Guyana|GUY
        Haiti|HAI
        Honduras|HON
        Hong Kong|HKG
        Hungary|HUN
        Iceland|ISL
        India|IND
        Indonesia|IDN
        Iran|IRN
        Iraq|IRQ
        Israel|ISR
        Italy|ITA
        Ivory Coast|CIV
        Jamaica|JAM
        Japan|JPN
        Jordan|JOR
        Kazakhstan|KAZ
        Kenya|KEN
        Kosovo|KOS
        Kuwait|KUW
        Kyrgyzstan|KGZ
        Laos|LAO
        Latvia|LVA
        Lebanon|LBN
        Lesotho|LES
        Liberia|LBR
        Libya|LBY
        Liechtenstein|LIE
        Lithuania|LTU
        Luxembourg|LUX
        Macau|MAC
        Madagascar|MAD
        Malawi|MWI
        Malaysia|MAS
        Maldives|MDV
        Mali|MLI
        Malta|MLT
        Mauritania|MTN
        Mauritius|MRI
        Mexico|MEX
        Moldova|MDA
        Mongolia|MNG
        Montenegro|MNE
        Montserrat|MSR
        Morocco|MAR
        Mozambique|MOZ
        Myanmar|MYA
        Namibia|NAM
        Nepal|NEP
        Netherlands|NED
        New Caledonia|NCL
        New Zealand|NZL
        Nicaragua|NCA
        Niger|NIG
        Nigeria|NGA
        North Korea|PRK
        North Macedonia|MKD
        Northern Ireland|NIR
        Norway|NOR
        Oman|OMA
        Pakistan|PAK
        Palestine|PLE
        Panama|PAN
        Papua New Guinea|PNG
        Paraguay|PAR
        Peru|PER
        Philippines|PHI
        Poland|POL
        Portugal|POR
        Puerto Rico|PUR
        Qatar|QAT
        Republic of Ireland|IRL
        Romania|ROU
        Russia|RUS
        Rwanda|RWA
        Saint Kitts and Nevis|SKN
        Saint Lucia|LCA
        Saint Vincent and the Grenadines|VIN
        Samoa|SAM
        San Marino|SMR
        São Tomé and Príncipe|STP
        Saudi Arabia|KSA
        Scotland|SCO
        Senegal|SEN
        Serbia|SRB
        Seychelles|SEY
        Sierra Leone|SLE
        Singapore|SGP
        Slovakia|SVK
        Slovenia|SVN
        Solomon Islands|SOL
        Somalia|SOM
        South Africa|RSA
        South Korea|KOR
        South Sudan|SSD
        Spain|ESP
        Sri Lanka|SRI
        Sudan|SDN
        Suriname|SUR
        Sweden|SWE
        Switzerland|SUI
        Syria|SYR
        Tahiti|TAH
        Tajikistan|TJK
        Tanzania|TAN
        Thailand|THA
        Timor-Leste|TLS
        Togo|TOG
        Tonga|TGA
        Trinidad and Tobago|TRI
        Tunisia|TUN
        Turkey|TUR
        Turkmenistan|TKM
        Turks and Caicos Islands|TCA
        Uganda|UGA
        Ukraine|UKR
        United Arab Emirates|UAE
        United States|USA
        Uruguay|URU
        U.S. Virgin Islands|VIR
        Uzbekistan|UZB
        Vanuatu|VAN
        Venezuela|VEN
        Vietnam|VIE
        Wales|WAL
        Yemen|YEM
        Zambia|ZAM
        Zimbabwe|ZIM
    `
)
