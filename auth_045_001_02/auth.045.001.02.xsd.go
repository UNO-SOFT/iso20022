// Code generated by download. DO NOT EDIT.

package iso20022_auth_045_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	FinInstrmRptgNonEqtyTradgActvtyRslt FinancialInstrumentReportingNonEquityTradingActivityResultV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FinInstrmRptgNonEqtyTradgActvtyRslt"`
}

type FinancialInstrumentReportingNonEquityTradingActivityResultV02 struct {
	RptHdr               SecuritiesMarketReportHeader1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 RptHdr"`
	NonEqtyTrnsprncyData []TransparencyDataReport18    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 NonEqtyTrnsprncyData"`
	SplmtryData          []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 SplmtryData,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 50 items long
type Max50Text string

// May be one of SFPS, SDRV, DERV, EMAL, BOND, ETCS, ETNS
type NonEquityInstrumentReportingClassification1Code string

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 ToDt"`
}

type Period4Choice struct {
	Dt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Dt,omitempty"`
	FrDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FrDt,omitempty"`
	ToDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 ToDt,omitempty"`
	FrDtToDt Period2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FrDtToDt,omitempty"`
}

type SecuritiesMarketReportHeader1 struct {
	RptgNtty     TradingVenueIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 RptgNtty"`
	RptgPrd      Period4Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 RptgPrd"`
	SubmissnDtTm ISODateTime                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 SubmissnDtTm,omitempty"`
}

type StatisticsTransparency2 struct {
	TtlNbOfTxsExctd  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 TtlNbOfTxsExctd"`
	TtlVolOfTxsExctd float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 TtlVolOfTxsExctd"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TonsOrCurrency2Choice struct {
	Nb  float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Nb,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Amt,omitempty"`
}

// May be one of APPA, CTPS
type TradingVenue2Code string

type TradingVenueIdentification1Choice struct {
	MktIdCd          MICIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 MktIdCd,omitempty"`
	NtlCmptntAuthrty CountryCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 NtlCmptntAuthrty,omitempty"`
	Othr             TradingVenueIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Othr,omitempty"`
}

type TradingVenueIdentification2 struct {
	Id Max50Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Id"`
	Tp TradingVenue2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Tp"`
}

type TransparencyDataReport18 struct {
	TechRcrdId                  Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 TechRcrdId,omitempty"`
	Id                          ISINOct2015Identifier                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Id"`
	FinInstrmClssfctn           NonEquityInstrumentReportingClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FinInstrmClssfctn,omitempty"`
	FullNm                      Max350Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 FullNm,omitempty"`
	TradgVn                     MICIdentifier                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 TradgVn,omitempty"`
	RptgPrd                     Period4Choice                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 RptgPrd,omitempty"`
	Lqdty                       bool                                            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Lqdty,omitempty"`
	PreTradLrgInScaleThrshld    TonsOrCurrency2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 PreTradLrgInScaleThrshld,omitempty"`
	PstTradLrgInScaleThrshld    TonsOrCurrency2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 PstTradLrgInScaleThrshld,omitempty"`
	PreTradInstrmSzSpcfcThrshld TonsOrCurrency2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 PreTradInstrmSzSpcfcThrshld,omitempty"`
	PstTradInstrmSzSpcfcThrshld TonsOrCurrency2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 PstTradInstrmSzSpcfcThrshld,omitempty"`
	Sttstcs                     StatisticsTransparency2                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.045.001.02 Sttstcs,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
