// Code generated by gen. DO NOT EDIT.

package iso20022_camt_017_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type CurrencyExchange20 struct {
	XchgRate float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 XchgRate"`
	QtdCcy   ActiveOrHistoricCurrencyCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 QtdCcy"`
	QtnDt    ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 QtnDt"`
	LwLmt    ExchangeRateOrPercentage1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 LwLmt,omitempty"`
	HghLmt   ExchangeRateOrPercentage1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 HghLmt,omitempty"`
}

type CurrencyExchangeReport4 struct {
	CcyRef       CurrencySourceTarget1            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CcyRef"`
	CcyXchgOrErr ExchangeRateReportOrError4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CcyXchgOrErr"`
}

type CurrencySourceTarget1 struct {
	SrcCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 SrcCcy"`
	TrgtCcy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 TrgtCcy"`
}

type Document struct {
	RtrCcyXchgRate ReturnCurrencyExchangeRateV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 RtrCcyXchgRate"`
}

type ErrorHandling1Choice struct {
	Cd    ErrorHandling1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Cd,omitempty"`
	Prtry Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Prtry,omitempty"`
}

// May be one of X020, X030, X050
type ErrorHandling1Code string

type ErrorHandling3 struct {
	Err  ErrorHandling1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Desc,omitempty"`
}

type ExchangeRateOrPercentage1Choice struct {
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Rate,omitempty"`
	Pctg float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Pctg,omitempty"`
}

type ExchangeRateReportOrError3Choice struct {
	CcyXchgRpt []CurrencyExchangeReport4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CcyXchgRpt,omitempty"`
	OprlErr    []ErrorHandling3          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 OprlErr,omitempty"`
}

type ExchangeRateReportOrError4Choice struct {
	BizErr  []ErrorHandling3   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 BizErr,omitempty"`
	CcyXchg CurrencyExchange20 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CcyXchg,omitempty"`
}

// May be no more than 4 items long
type ExternalEnquiryRequestType1Code string

// May be no more than 4 items long
type ExternalPaymentControlRequestType1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Issr,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

type MessageHeader7 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CreDtTm,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 ReqTp,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 QryNm,omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 CreDtTm,omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 PmtCtrl,omitempty"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Enqry,omitempty"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Prtry,omitempty"`
}

type ReturnCurrencyExchangeRateV05 struct {
	MsgHdr      MessageHeader7                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 MsgHdr"`
	RptOrErr    ExchangeRateReportOrError3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 RptOrErr"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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