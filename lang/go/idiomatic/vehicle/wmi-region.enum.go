package vehicle

//go:generate enumer -path=./wmi-region.enum.go

type WmiRegion string

type wmiRegions struct {
	Unassigned         WmiRegion `enum:"unassigned"`
	SouthAfrica        WmiRegion `enum:"south-africa"`
	CoteDIvoire        WmiRegion `enum:"cote-d-ivoire"`
	Angola             WmiRegion `enum:"angola"`
	Kenya              WmiRegion `enum:"kenya"`
	Tanzania           WmiRegion `enum:"tanzania"`
	Uganda             WmiRegion `enum:"uganda"`
	Benin              WmiRegion `enum:"benin"`
	Madagascar         WmiRegion `enum:"madagascar"`
	Tunisia            WmiRegion `enum:"tunisia"`
	Egypt              WmiRegion `enum:"egypt"`
	Morocco            WmiRegion `enum:"morocco"`
	Zambia             WmiRegion `enum:"zambia"`
	Ethiopia           WmiRegion `enum:"ethiopia"`
	Mozambique         WmiRegion `enum:"mozambique"`
	Ghana              WmiRegion `enum:"ghana"`
	Nigeria            WmiRegion `enum:"nigeria"`
	SriLanka           WmiRegion `enum:"sri-lanka"`
	Israel             WmiRegion `enum:"israel"`
	SouthKorea         WmiRegion `enum:"south-korea"`
	Jordan             WmiRegion `enum:"jordan"`
	China              WmiRegion `enum:"china"`
	India              WmiRegion `enum:"india"`
	Indonesia          WmiRegion `enum:"indonesia"`
	Thailand           WmiRegion `enum:"thailand"`
	Myanmar            WmiRegion `enum:"myanmar"`
	Mongolia           WmiRegion `enum:"mongolia"`
	Kazakhstan         WmiRegion `enum:"kazakhstan"`
	Iran               WmiRegion `enum:"iran"`
	Pakistan           WmiRegion `enum:"pakistan"`
	Turkey             WmiRegion `enum:"turkey"`
	Uzbekistan         WmiRegion `enum:"uzbekistan"`
	Philippines        WmiRegion `enum:"philippines"`
	Singapore          WmiRegion `enum:"singapore"`
	Malaysia           WmiRegion `enum:"malaysia"`
	Bangladesh         WmiRegion `enum:"bangladesh"`
	UnitedArabEmirates WmiRegion `enum:"united-arab-emirates"`
	Taiwan             WmiRegion `enum:"taiwan"`
	SaudiArabia        WmiRegion `enum:"saudi-arabia"`
	Vietnam            WmiRegion `enum:"vietnam"`
	Hongkong           WmiRegion `enum:"hongkong"`
	UnitedKingdom      WmiRegion `enum:"united-kingdom"`
	Germany            WmiRegion `enum:"germany"`
}
