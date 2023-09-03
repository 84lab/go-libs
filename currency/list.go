package currency

type ISO string

const (
	AFA ISO = "AFA"
	ALL ISO = "ALL"
	DZD ISO = "DZD"
	ADP ISO = "ADP"
	AZM ISO = "AZM"
	ARS ISO = "ARS"
	AUD ISO = "AUD"
	BSD ISO = "BSD"
	BHD ISO = "BHD"
	BDT ISO = "BDT"
	AMD ISO = "AMD"
	BBD ISO = "BBD"
	BMD ISO = "BMD"
	BTN ISO = "BTN"
	BOB ISO = "BOB"
	BWP ISO = "BWP"
	BZD ISO = "BZD"
	SBD ISO = "SBD"
	BND ISO = "BND"
	BGL ISO = "BGL"
	MMK ISO = "MMK"
	BIF ISO = "BIF"
	KHR ISO = "KHR"
	CAD ISO = "CAD"
	CVE ISO = "CVE"
	KYD ISO = "KYD"
	LKR ISO = "LKR"
	CLP ISO = "CLP"
	CNY ISO = "CNY"
	COP ISO = "COP"
	KMF ISO = "KMF"
	CRC ISO = "CRC"
	HRK ISO = "HRK"
	CUP ISO = "CUP"
	CYP ISO = "CYP"
	CZK ISO = "CZK"
	DKK ISO = "DKK"
	DOP ISO = "DOP"
	ECS ISO = "ECS"
	SVC ISO = "SVC"
	ETB ISO = "ETB"
	ERN ISO = "ERN"
	EEK ISO = "EEK"
	FKP ISO = "FKP"
	FJD ISO = "FJD"
	DJF ISO = "DJF"
	GMD ISO = "GMD"
	GHC ISO = "GHC"
	GIP ISO = "GIP"
	GTQ ISO = "GTQ"
	GNF ISO = "GNF"
	GYD ISO = "GYD"
	HTG ISO = "HTG"
	HNL ISO = "HNL"
	HKD ISO = "HKD"
	HUF ISO = "HUF"
	ISK ISO = "ISK"
	INR ISO = "INR"
	IDR ISO = "IDR"
	IRR ISO = "IRR"
	IQD ISO = "IQD"
	ILS ISO = "ILS"
	JMD ISO = "JMD"
	JPY ISO = "JPY"
	KZT ISO = "KZT"
	JOD ISO = "JOD"
	KES ISO = "KES"
	KPW ISO = "KPW"
	KRW ISO = "KRW"
	KWD ISO = "KWD"
	KGS ISO = "KGS"
	LAK ISO = "LAK"
	LBP ISO = "LBP"
	LSL ISO = "LSL"
	LVL ISO = "LVL"
	LRD ISO = "LRD"
	LYD ISO = "LYD"
	LTL ISO = "LTL"
	MOP ISO = "MOP"
	MGF ISO = "MGF"
	MWK ISO = "MWK"
	MYR ISO = "MYR"
	MVR ISO = "MVR"
	MTL ISO = "MTL"
	MRO ISO = "MRO"
	MUR ISO = "MUR"
	MXN ISO = "MXN"
	MNT ISO = "MNT"
	MDL ISO = "MDL"
	MAD ISO = "MAD"
	MZM ISO = "MZM"
	OMR ISO = "OMR"
	NAD ISO = "NAD"
	NPR ISO = "NPR"
	ANG ISO = "ANG"
	AWG ISO = "AWG"
	VUV ISO = "VUV"
	NZD ISO = "NZD"
	NIO ISO = "NIO"
	NGN ISO = "NGN"
	NOK ISO = "NOK"
	PKR ISO = "PKR"
	PAB ISO = "PAB"
	PGK ISO = "PGK"
	PYG ISO = "PYG"
	PEN ISO = "PEN"
	PHP ISO = "PHP"
	GWP ISO = "GWP"
	TPE ISO = "TPE"
	QAR ISO = "QAR"
	ROL ISO = "ROL"
	RUB ISO = "RUB"
	RWF ISO = "RWF"
	SHP ISO = "SHP"
	STD ISO = "STD"
	SAR ISO = "SAR"
	SCR ISO = "SCR"
	SLL ISO = "SLL"
	SGD ISO = "SGD"
	SKK ISO = "SKK"
	VND ISO = "VND"
	SIT ISO = "SIT"
	SOS ISO = "SOS"
	ZAR ISO = "ZAR"
	ZWD ISO = "ZWD"
	SDD ISO = "SDD"
	SRG ISO = "SRG"
	SZL ISO = "SZL"
	SEK ISO = "SEK"
	CHF ISO = "CHF"
	SYP ISO = "SYP"
	THB ISO = "THB"
	TOP ISO = "TOP"
	TTD ISO = "TTD"
	AED ISO = "AED"
	TND ISO = "TND"
	TRL ISO = "TRL"
	TMM ISO = "TMM"
	UGX ISO = "UGX"
	MKD ISO = "MKD"
	EGP ISO = "EGP"
	GBP ISO = "GBP"
	TZS ISO = "TZS"
	USD ISO = "USD"
	UYU ISO = "UYU"
	UZS ISO = "UZS"
	VEB ISO = "VEB"
	YER ISO = "YER"
	YUM ISO = "YUM"
	ZMK ISO = "ZMK"
	TWD ISO = "TWD"
	TRY ISO = "TRY"
	XAF ISO = "XAF"
	XCD ISO = "XCD"
	XOF ISO = "XOF"
	XPF ISO = "XPF"
	TJS ISO = "TJS"
	AOA ISO = "AOA"
	BYR ISO = "BYR"
	BGN ISO = "BGN"
	CDF ISO = "CDF"
	BAM ISO = "BAM"
	EUR ISO = "EUR"
	MXV ISO = "MXV"
	UAH ISO = "UAH"
	GEL ISO = "GEL"
	ECV ISO = "ECV"
	BOV ISO = "BOV"
	PLN ISO = "PLN"
	BRL ISO = "BRL"
	CLF ISO = "CLF"
)

type Info struct {
	ISO   ISO
	Label string
}

//nolint:funlen
func GetCurrencyMap() map[ISO]Info {
	return map[ISO]Info{
		// AFA: "Afghani",
		// ALL: "Leck",
		// DZD: "Algerian Dinar",
		// ADP: "Andorran Peseta",
		// AZM: "Azerbaijanian Manat",
		// ARS: "Argentine Peso",
		// AUD: "Australian Dollar",
		// BSD: "Bahamian Dollar",
		// BHD: "Bahraini Dinar",
		// BDT: "Taka",
		// AMD: "Armenian Dram",
		// BBD: "Barbados Dollar",
		// BMD: "Bermudian Dollar",
		// BTN: "Ngultrum",
		// BOB: "Boliviano",
		// BWP: "Pula",
		// BZD: "Belize Dollar",
		// SBD: "Solomon Islands Dollar",
		// BND: "Brunei Dollar",
		// BGL: "Lev",
		// MMK: "Kyat",
		// BIF: "Burundi Franc",
		// KHR: "Riel",
		// CAD: "Canadian Dollar",
		// CVE: "Cape Verde Escudo",
		// KYD: "Cayman Islands Dollar",
		// LKR: "Sri Lanka Rupee",
		// CLP: "Chilean Peso",
		// CNY: "Yuan Renminbi",
		// COP: "Colombian Peso",
		// KMF: "Comoro Franc",
		// CRC: "Costa Rican Colon",
		// HRK: "Croatian kuna",
		// CUP: "Cuban Peso",
		// CYP: "Cyprus Pound",
		// CZK: "Czech Koruna",
		// DKK: "Danish Krone",
		// DOP: "Dominican Peso",
		// ECS: "Sucre",
		// SVC: "El Salvador Colon",
		// ETB: "Ethiopian Birr",
		// ERN: "Nakfa",
		// EEK: "Kroon",
		// FKP: "Falkland Islands Pound",
		// FJD: "Fiji Dollar",
		// DJF: "Djibouti Franc",
		// GMD: "Dalasi",
		// GHC: "Cedi",
		// GIP: "Gibraltar Pound",
		// GTQ: "Quetzal",
		// GNF: "Guinea Franc",
		// GYD: "Guyana Dollar",
		// HTG: "Gourde",
		// HNL: "Lempira",
		// HKD: "Hong Kong Dollar",
		// HUF: "Forint",
		// ISK: "Iceland Krona",
		// INR: "Indian Rupee",
		// IDR: "Rupiah",
		// IRR: "Iranian Rial",
		// IQD: "Iraqi Dinar",
		// ILS: "New Israeli Sheqel",
		// JMD: "Jamaican Dollar",
		// JPY: "Yen",
		// KZT: "Tenge",
		// JOD: "Jordanian Dinar",
		// KES: "Kenyan Shilling",
		// KPW: "North Korean Won",
		KRW: {
			ISO:   KRW,
			Label: "South Korean Won",
		},
		// KWD: "Kuwaiti Dinar",
		// KGS: "Som",
		// LAK: "Kip",
		// LBP: "Lebanese Pound",
		// LSL: "Loti",
		// LVL: "Latvian Lats",
		// LRD: "Liberian Dollar",
		// LYD: "Lybian Dinar",
		// LTL: "Lithuanian Litus",
		// MOP: "Pataca",
		// MGF: "Malagasy Franc",
		// MWK: "Kwacha",
		// MYR: "Malaysian Ringgit",
		// MVR: "Rufiyaa",
		// MTL: "Maltese Lira",
		// MRO: "Ouguiya",
		// MUR: "Mauritius Rupee",
		// MXN: "Mexican Peso",
		// MNT: "Tugrik",
		// MDL: "Moldovan Leu",
		// MAD: "Moroccan Dirham",
		// MZM: "Metical",
		// OMR: "Rial Omani",
		// NAD: "Namibia Dollar",
		// NPR: "Nepalese Rupee",
		// ANG: "Netherlands Antillan Guilder",
		// AWG: "Aruban Guilder",
		// VUV: "Vatu",
		// NZD: "New Zealand Dollar",
		// NIO: "Cordoba Oro",
		// NGN: "Naira",
		// NOK: "Norwegian Krone",
		// PKR: "Pakistan Rupee",
		// PAB: "Balboa",
		// PGK: "Kina",
		// PYG: "Guarani",
		// PEN: "Nuevo Sol",
		// PHP: "Philippine Peso",
		// GWP: "Guinea-Bissau Peso",
		// TPE: "Timor Escudo",
		// QAR: "Qatari Rial",
		// ROL: "Leu",
		RUB: {
			ISO:   RUB,
			Label: "Russian Ruble",
		},
		// RWF: "Rwanda Franc",
		// SHP: "Saint Helena Pound",
		// STD: "Dobra",
		// SAR: "Saudi Riyal",
		// SCR: "Seychelles Rupee",
		// SLL: "Leone",
		// SGD: "Singapore Dollar",
		// SKK: "Slovak Koruna",
		// VND: "Dong",
		// SIT: "Tolar",
		// SOS: "Somali Shilling",
		// ZAR: "Rand",
		// ZWD: "Zimbabwe Dollar",
		// SDD: "Sudanese Dinar",
		// SRG: "Suriname Guilder",
		// SZL: "Lilangeni",
		// SEK: "Swedish Krona",
		// CHF: "Swiss Franc",
		// SYP: "Syrian Pound",
		// THB: "Baht",
		// TOP: "Pa´anga",
		// TTD: "Trinidad and Tobago Dollar",
		// AED: "UAE Dirham",
		// TND: "Tunisian Dinar",
		// TRL: "Turkish Lira",
		// TMM: "Manat",
		// UGX: "Uganda Shilling",
		// MKD: "Denar",
		// EGP: "Egyptian Pound",
		// GBP: "Pound Sterling",
		// TZS: "Tanzanian Shilling",
		USD: {
			ISO:   USD,
			Label: "US Dollar",
		},
		// UYU: "Peso Uruguayo",
		// UZS: "Uzbekistan Sum",
		// VEB: "Bolivar",
		// YER: "Yemeni Rial",
		// YUM: "Yugoslavian Dinar",
		// ZMK: "Kwacha",
		// TWD: "New Taiwan Dollar",
		// TRY: "New Turkish Lira",
		// XAF: "CFA Franc BEAC",
		// XCD: "East Caribbean Dollar",
		// XOF: "CFA Franc BCEAO",
		// XPF: "CFP Franc",
		// TJS: "Somoni",
		// AOA: "Kwanza",
		// BYR: "Belarussian Ruble",
		// BGN: "Bulgarian Lev",
		// CDF: "Franc Congolais",
		// BAM: "Convertible Marks",
		EUR: {
			ISO:   EUR,
			Label: "Euro",
		},
		// MXV: "Mexican Unidad de Inversion (UDI)",
		// UAH: "Hryvnia",
		// GEL: "Lari",
		// ECV: "Unidad de Valor Constante (UVC)",
		// BOV: "Mvdol",
		// PLN: "Zloty",
		// BRL: "Brazilian Real",
		// CLF: "Unidades de fomento",
	}
}
