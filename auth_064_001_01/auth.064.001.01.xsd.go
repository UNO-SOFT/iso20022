// Code generated by download. DO NOT EDIT.

package iso20022_auth_064_001_01

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AvailableFinancialResourcesAmount1 struct {
	TtlInitlMrgn        ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 TtlInitlMrgn"`
	TtlPrfnddDfltFnd    ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 TtlPrfnddDfltFnd"`
	CCPSkinInTheGame    []ReportingAssetBreakdown1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 CCPSkinInTheGame"`
	OthrDfltFndCntrbtn  ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 OthrDfltFndCntrbtn"`
	UfnddMmbCmmtmnt     ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 UfnddMmbCmmtmnt"`
	UfnddThrdPtyCmmtmnt ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 UfnddThrdPtyCmmtmnt"`
}

type CCPAvailableFinancialResourcesReportV01 struct {
	AvlblFinRsrcsAmt AvailableFinancialResourcesAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 AvlblFinRsrcsAmt"`
	OthrPrfnddRsrcs  ReportingAssetBreakdown1           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 OthrPrfnddRsrcs,omitempty"`
	SplmtryData      []SupplementaryData1               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 SplmtryData,omitempty"`
}

type Document struct {
	CCPAvlblFinRsrcsRpt CCPAvailableFinancialResourcesReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 CCPAvlblFinRsrcsRpt"`
}

// May be no more than 350 items long
type Max350Text string

// May be one of BOND, CASH, OTHR, EQUI
type ProductType6Code string

type ReportingAssetBreakdown1 struct {
	RptgAsstTp ProductType6Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 RptgAsstTp"`
	Id         Max350Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 Id,omitempty"`
	Amt        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 Amt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.064.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
