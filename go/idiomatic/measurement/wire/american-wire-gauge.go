package wire

// https://en.wikipedia.org/wiki/American_wire_gauge

type AmericanWireGauge struct {
	Name             string
	CommonName       string
	DiameterInInches float64
}

var Wires = wires{
	Awg0000: AmericanWireGauge{
		Name:             "0000",
		CommonName:       "4/0",
		DiameterInInches: 0.4600,
	},
	Awg000: AmericanWireGauge{
		Name:             "000",
		CommonName:       "3/0",
		DiameterInInches: 0.4096,
	},
	Awg00: AmericanWireGauge{
		Name:             "00",
		CommonName:       "2/0",
		DiameterInInches: 0.3648,
	},
	Awg0: AmericanWireGauge{
		Name:             "0",
		CommonName:       "1/0",
		DiameterInInches: 0.3249,
	},
	Awg1: AmericanWireGauge{
		Name:             "1",
		CommonName:       "1",
		DiameterInInches: 0.2893,
	},
	Awg2: AmericanWireGauge{
		Name:             "2",
		CommonName:       "2",
		DiameterInInches: 0.2576,
	},
	Awg3: AmericanWireGauge{
		Name:             "3",
		CommonName:       "3",
		DiameterInInches: 0.2294,
	},
	Awg4: AmericanWireGauge{
		Name:             "4",
		CommonName:       "4",
		DiameterInInches: 0.2043,
	},
	Awg5: AmericanWireGauge{
		Name:             "5",
		CommonName:       "5",
		DiameterInInches: 0.1819,
	},
	Awg6: AmericanWireGauge{
		Name:             "6",
		CommonName:       "6",
		DiameterInInches: 0.1620,
	},
	Awg7: AmericanWireGauge{
		Name:             "7",
		CommonName:       "7",
		DiameterInInches: 0.1443,
	},
	Awg8: AmericanWireGauge{
		Name:             "8",
		CommonName:       "8",
		DiameterInInches: 0.1285,
	},
	Awg9: AmericanWireGauge{
		Name:             "9",
		CommonName:       "9",
		DiameterInInches: 0.1144,
	},
	Awg10: AmericanWireGauge{
		Name:             "10",
		CommonName:       "10",
		DiameterInInches: 0.1019,
	},
	Awg11: AmericanWireGauge{
		Name:             "11",
		CommonName:       "11",
		DiameterInInches: 0.0907,
	},
	Awg12: AmericanWireGauge{
		Name:             "12",
		CommonName:       "12",
		DiameterInInches: 0.0808,
	},
	Awg13: AmericanWireGauge{
		Name:             "13",
		CommonName:       "13",
		DiameterInInches: 0.0720,
	},
	Awg14: AmericanWireGauge{
		Name:             "14",
		CommonName:       "14",
		DiameterInInches: 0.0641,
	},
	Awg15: AmericanWireGauge{
		Name:             "15",
		CommonName:       "15",
		DiameterInInches: 0.0571,
	},
	Awg16: AmericanWireGauge{
		Name:             "16",
		CommonName:       "16",
		DiameterInInches: 0.0508,
	},
	Awg17: AmericanWireGauge{
		Name:             "17",
		CommonName:       "17",
		DiameterInInches: 0.0453,
	},
	Awg18: AmericanWireGauge{
		Name:             "18",
		CommonName:       "18",
		DiameterInInches: 0.0403,
	},
	Awg19: AmericanWireGauge{
		Name:             "19",
		CommonName:       "19",
		DiameterInInches: 0.0359,
	},
	Awg20: AmericanWireGauge{
		Name:             "20",
		CommonName:       "20",
		DiameterInInches: 0.0320,
	},
	Awg21: AmericanWireGauge{
		Name:             "21",
		CommonName:       "21",
		DiameterInInches: 0.0285,
	},
	Awg22: AmericanWireGauge{
		Name:             "22",
		CommonName:       "22",
		DiameterInInches: 0.0253,
	},
	Awg23: AmericanWireGauge{
		Name:             "23",
		CommonName:       "23",
		DiameterInInches: 0.0226,
	},
	Awg24: AmericanWireGauge{
		Name:             "24",
		CommonName:       "24",
		DiameterInInches: 0.0201,
	},
	Awg25: AmericanWireGauge{
		Name:             "25",
		CommonName:       "25",
		DiameterInInches: 0.0179,
	},
	Awg26: AmericanWireGauge{
		Name:             "26",
		CommonName:       "26",
		DiameterInInches: 0.0159,
	},
	Awg27: AmericanWireGauge{
		Name:             "27",
		CommonName:       "27",
		DiameterInInches: 0.0142,
	},
	Awg28: AmericanWireGauge{
		Name:             "28",
		CommonName:       "28",
		DiameterInInches: 0.0126,
	},
	Awg29: AmericanWireGauge{
		Name:             "29",
		CommonName:       "29",
		DiameterInInches: 0.0113,
	},
	Awg30: AmericanWireGauge{
		Name:             "30",
		CommonName:       "30",
		DiameterInInches: 0.0100,
	},
	Awg31: AmericanWireGauge{
		Name:             "31",
		CommonName:       "31",
		DiameterInInches: 0.0089,
	},
	Awg32: AmericanWireGauge{
		Name:             "32",
		CommonName:       "32",
		DiameterInInches: 0.0080,
	},
	Awg33: AmericanWireGauge{
		Name:             "33",
		CommonName:       "33",
		DiameterInInches: 0.0071,
	},
	Awg34: AmericanWireGauge{
		Name:             "34",
		CommonName:       "34",
		DiameterInInches: 0.0063,
	},
	Awg35: AmericanWireGauge{
		Name:             "35",
		CommonName:       "35",
		DiameterInInches: 0.0056,
	},
	Awg36: AmericanWireGauge{
		Name:             "36",
		CommonName:       "36",
		DiameterInInches: 0.0050,
	},
	Awg37: AmericanWireGauge{
		Name:             "37",
		CommonName:       "37",
		DiameterInInches: 0.0045,
	},
	Awg38: AmericanWireGauge{
		Name:             "38",
		CommonName:       "38",
		DiameterInInches: 0.0040,
	},
	Awg39: AmericanWireGauge{
		Name:             "39",
		CommonName:       "39",
		DiameterInInches: 0.0035,
	},
	Awg40: AmericanWireGauge{
		Name:             "40",
		CommonName:       "40",
		DiameterInInches: 0.0031,
	},
}

type wires struct {
	Awg0000 AmericanWireGauge
	Awg000  AmericanWireGauge
	Awg00   AmericanWireGauge
	Awg0    AmericanWireGauge
	Awg1    AmericanWireGauge
	Awg2    AmericanWireGauge
	Awg3    AmericanWireGauge
	Awg4    AmericanWireGauge
	Awg5    AmericanWireGauge
	Awg6    AmericanWireGauge
	Awg7    AmericanWireGauge
	Awg8    AmericanWireGauge
	Awg9    AmericanWireGauge
	Awg10   AmericanWireGauge
	Awg11   AmericanWireGauge
	Awg12   AmericanWireGauge
	Awg13   AmericanWireGauge
	Awg14   AmericanWireGauge
	Awg15   AmericanWireGauge
	Awg16   AmericanWireGauge
	Awg17   AmericanWireGauge
	Awg18   AmericanWireGauge
	Awg19   AmericanWireGauge
	Awg20   AmericanWireGauge
	Awg21   AmericanWireGauge
	Awg22   AmericanWireGauge
	Awg23   AmericanWireGauge
	Awg24   AmericanWireGauge
	Awg25   AmericanWireGauge
	Awg26   AmericanWireGauge
	Awg27   AmericanWireGauge
	Awg28   AmericanWireGauge
	Awg29   AmericanWireGauge
	Awg30   AmericanWireGauge
	Awg31   AmericanWireGauge
	Awg32   AmericanWireGauge
	Awg33   AmericanWireGauge
	Awg34   AmericanWireGauge
	Awg35   AmericanWireGauge
	Awg36   AmericanWireGauge
	Awg37   AmericanWireGauge
	Awg38   AmericanWireGauge
	Awg39   AmericanWireGauge
	Awg40   AmericanWireGauge
}
