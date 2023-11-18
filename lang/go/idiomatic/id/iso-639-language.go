package id

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

// https://en.wikipedia.org/wiki/List_of_ISO_639|1_codes
// https://iso639|3.sil.org/code_tables/macrolanguage_mappings/data

type Iso639Language struct {
	Name string `json:"name,omitempty"`
	V1   string `json:"v1,omitempty"`
	V2t  string `json:"v2-t,omitempty"`
	V2b  string `json:"v2-b,omitempty"`
	V3   string `json:"v3,omitempty"`
}

func Iso639LanguageLookup(s string) (Iso639Language, bool) {
	l := stringer.ToLower(s)
	var lang Iso639Language
	var found bool

	lang, found = iso639NameToLang[l]

	if !found {
		lang, found = iso639V1ToLang[l]
	}

	if !found {
		lang, found = iso639V2tToLang[l]
	}

	if !found {
		lang, found = iso639V2bToLang[l]
	}

	if !found {
		lang, found = iso639V3ToName[l]
	}

	return lang, found
}

func init() {
	lines := stringer.Split(iso639Data, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = stringer.TrimSpace(lines[i])
	}

	for _, line := range lines {
		if stringer.IsEmpty(line) {
			continue
		}

		comps := stringer.Split(line, "|")

		lang := Iso639Language{
			Name: stringer.TrimSpace(comps[0]),
			V1:   stringer.TrimSpace(comps[1]),
			V2t:  stringer.TrimSpace(comps[2]),
			V2b:  stringer.TrimSpace(comps[3]),
			V3:   stringer.TrimSpace(comps[4]),
		}

		iso639NameToLang[lang.Name] = lang
		iso639V1ToLang[lang.V1] = lang
		iso639V2tToLang[lang.V2t] = lang
		iso639V2bToLang[lang.V2b] = lang
		iso639V3ToName[lang.V3] = lang
	}
}

var (
	iso639NameToLang = map[string]Iso639Language{}
	iso639V1ToLang   = map[string]Iso639Language{}
	iso639V2tToLang  = map[string]Iso639Language{}
	iso639V2bToLang  = map[string]Iso639Language{}
	iso639V3ToName   = map[string]Iso639Language{}
	iso639Data       = `
        Abkhazian|ab|abk|abk|abk
        Afar|aa|aar|aar|aar
        Afrikaans|af|afr|afr|afr
        Akan|ak|aka|aka|aka
        Albanian|sq|sqi|alb|sqi
        Amharic|am|amh|amh|amh
        Arabic|ar|ara|ara|ara
        Aragonese|an|arg|arg|arg
        Armenian|hy|hye|arm|hye
        Assamese|as|asm|asm|asm
        Avaric|av|ava|ava|ava
        Avestan|ae|ave|ave|ave
        Aymara|ay|aym|aym|aym
        Azerbaijani|az|aze|aze|aze
        Bambara|bm|bam|bam|bam
        Bashkir|ba|bak|bak|bak
        Basque|eu|eus|baq|eus
        Belarusian|be|bel|bel|bel
        Bengali|bn|ben|ben|ben
        Bislama|bi|bis|bis|bis
        Bosnian|bs|bos|bos|bos
        Breton|br|bre|bre|bre
        Bulgarian|bg|bul|bul|bul
        Burmese|my|mya|bur|mya
        Catalan, Valencian|ca|cat|cat|cat
        Chamorro|ch|cha|cha|cha
        Chechen|ce|che|che|che
        Chichewa, Chewa, Nyanja|ny|nya|nya|nya
        Chinese|zh|zho|chi|zho
        Church Slavonic, Old Slavonic, Old Church Slavonic|cu|chu|chu|chu
        Chuvash|cv|chv|chv|chv
        Cornish|kw|cor|cor|cor
        Corsican|co|cos|cos|cos
        Cree|cr|cre|cre|cre
        Croatian|hr|hrv|hrv|hrv
        Czech|cs|ces|cze|ces
        Danish|da|dan|dan|dan
        Divehi, Dhivehi, Maldivian|dv|div|div|div
        Dutch, Flemish|nl|nld|dut|nld
        Dzongkha|dz|dzo|dzo|dzo
        English|en|eng|eng|eng
        Esperanto|eo|epo|epo|epo
        Estonian|et|est|est|est
        Ewe|ee|ewe|ewe|ewe
        Faroese|fo|fao|fao|fao
        Fijian|fj|fij|fij|fij
        Finnish|fi|fin|fin|fin
        French|fr|fra|fre|fra
        Western Frisian|fy|fry|fry|fry
        Fulah|ff|ful|ful|ful
        Gaelic, Scottish Gaelic|gd|gla|gla|gla
        Galician|gl|glg|glg|glg
        Ganda|lg|lug|lug|lug
        Georgian|ka|kat|geo|kat
        German|de|deu|ger|deu
        Greek, Modern|el|ell|gre|ell
        Kalaallisut, Greenlandic|kl|kal|kal|kal
        Guarani|gn|grn|grn|grn
        Gujarati|gu|guj|guj|guj
        Haitian, Haitian Creole|ht|hat|hat|hat
        Hausa|ha|hau|hau|hau
        Hebrew|he|heb|heb|heb
        Herero|hz|her|her|her
        Hindi|hi|hin|hin|hin
        Hiri Motu|ho|hmo|hmo|hmo
        Hungarian|hu|hun|hun|hun
        Icelandic|is|isl|ice|isl
        Ido|io|ido|ido|ido
        Igbo|ig|ibo|ibo|ibo
        Indonesian|id|ind|ind|ind
        Interlingua|ia|ina|ina|ina
        Interlingue, Occidental|ie|ile|ile|ile
        Inuktitut|iu|iku|iku|iku
        Inupiaq|ik|ipk|ipk|ipk
        Irish|ga|gle|gle|gle
        Italian|it|ita|ita|ita
        Japanese|ja|jpn|jpn|jpn
        Javanese|jv|jav|jav|jav
        Kannada|kn|kan|kan|kan
        Kanuri|kr|kau|kau|kau
        Kashmiri|ks|kas|kas|kas
        Kazakh|kk|kaz|kaz|kaz
        Central Khmer|km|khm|khm|khm
        Kikuyu, Gikuyu|ki|kik|kik|kik
        Kinyarwanda|rw|kin|kin|kin
        Kirghiz, Kyrgyz|ky|kir|kir|kir
        Komi|kv|kom|kom|kom
        Kongo|kg|kon|kon|kon
        Korean|ko|kor|kor|kor
        Kuanyama, Kwanyama|kj|kua|kua|kua
        Kurdish|ku|kur|kur|kur
        Lao|lo|lao|lao|lao
        Latin|la|lat|lat|lat
        Latvian|lv|lav|lav|lav
        Limburgan, Limburger, Limburgish|li|lim|lim|lim
        Lingala|ln|lin|lin|lin
        Lithuanian|lt|lit|lit|lit
        Luba|Katanga|lu|lub|lub|lub
        Luxembourgish, Letzeburgesch|lb|ltz|ltz|ltz
        Macedonian|mk|mkd|mac|mkd
        Malagasy|mg|mlg|mlg|mlg
        Malay|ms|msa|may|msa
        Malayalam|ml|mal|mal|mal
        Maltese|mt|mlt|mlt|mlt
        Manx|gv|glv|glv|glv
        Maori|mi|mri|mao|mri
        Marathi|mr|mar|mar|mar
        Marshallese|mh|mah|mah|mah
        Mongolian|mn|mon|mon|mon
        Nauru|na|nau|nau|nau
        Navajo, Navaho|nv|nav|nav|nav
        North Ndebele|nd|nde|nde|nde
        South Ndebele|nr|nbl|nbl|nbl
        Ndonga|ng|ndo|ndo|ndo
        Nepali|ne|nep|nep|nep
        Norwegian|no|nor|nor|nor
        Norwegian Bokmål|nb|nob|nob|nob
        Norwegian Nynorsk|nn|nno|nno|nno
        Sichuan Yi, Nuosu|ii|iii|iii|iii
        Occitan|oc|oci|oci|oci
        Ojibwa|oj|oji|oji|oji
        Oriya|or|ori|ori|ori
        Oromo|om|orm|orm|orm
        Ossetian, Ossetic|os|oss|oss|oss
        Pali|pi|pli|pli|pli
        Pashto, Pushto|ps|pus|pus|pus
        Persian|fa|fas|per|fas
        Polish|pl|pol|pol|pol
        Portuguese|pt|por|por|por
        Punjabi, Panjabi|pa|pan|pan|pan
        Quechua|qu|que|que|que
        Romanian, Moldavian, Moldovan|ro|ron|rum|ron
        Romansh|rm|roh|roh|roh
        Rundi|rn|run|run|run
        Russian|ru|rus|rus|rus
        Northern Sami|se|sme|sme|sme
        Samoan|sm|smo|smo|smo
        Sango|sg|sag|sag|sag
        Sanskrit|sa|san|san|san
        Sardinian|sc|srd|srd|srd
        Serbian|sr|srp|srp|srp
        Shona|sn|sna|sna|sna
        Sindhi|sd|snd|snd|snd
        Sinhala, Sinhalese|si|sin|sin|sin
        Slovak|sk|slk|slo|slk
        Slovenian|sl|slv|slv|slv
        Somali|so|som|som|som
        Southern Sotho|st|sot|sot|sot
        Spanish, Castilian|es|spa|spa|spa
        Sundanese|su|sun|sun|sun
        Swahili|sw|swa|swa|swa
        Swati|ss|ssw|ssw|ssw
        Swedish|sv|swe|swe|swe
        Tagalog|tl|tgl|tgl|tgl
        Tahitian|ty|tah|tah|tah
        Tajik|tg|tgk|tgk|tgk
        Tamil|ta|tam|tam|tam
        Tatar|tt|tat|tat|tat
        Telugu|te|tel|tel|tel
        Thai|th|tha|tha|tha
        Tibetan|bo|bod|tib|bod
        Tigrinya|ti|tir|tir|tir
        Tonga|to|ton|ton|ton
        Tsonga|ts|tso|tso|tso
        Tswana|tn|tsn|tsn|tsn
        Turkish|tr|tur|tur|tur
        Turkmen|tk|tuk|tuk|tuk
        Twi|tw|twi|twi|twi
        Uighur, Uyghur|ug|uig|uig|uig
        Ukrainian|uk|ukr|ukr|ukr
        Urdu|ur|urd|urd|urd
        Uzbek|uz|uzb|uzb|uzb
        Venda|ve|ven|ven|ven
        Vietnamese|vi|vie|vie|vie
        Volapük|vo|vol|vol|vol
        Walloon|wa|wln|wln|wln
        Welsh|cy|cym|wel|cym
        Wolof|wo|wol|wol|wol
        Xhosa|xh|xho|xho|xho
        Yiddish|yi|yid|yid|yid
        Yoruba|yo|yor|yor|yor
        Zhuang, Chuang|za|zha|zha|zha
        Zulu|zu|zul|zul|zul
    `
)
