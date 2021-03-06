// Code generated by gen. DO NOT EDIT.

package iso20022_reda_013_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Document struct {
	SctyDeltnReq SecurityDeletionRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 SctyDeltnReq"`
}

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISIN2021Identifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Prtry,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 CreDtTm,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Tp"`
}

type SecurityDeletionRequestV01 struct {
	MsgHdr      MessageHeader1           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 MsgHdr,omitempty"`
	FinInstrmId SecurityIdentification39 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 FinInstrmId"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 SplmtryData,omitempty"`
}

type SecurityIdentification39 struct {
	ISIN   ISIN2021Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.013.001.01 Envlp"`
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
