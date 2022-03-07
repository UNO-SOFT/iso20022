// Code generated by gen. DO NOT EDIT.

package iso20022_admi_006_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	RctAck ReceiptAcknowledgementV01 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 RctAck"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 SchmeNm,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 70 items long
type Max70Text string

type MessageHeader10 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 CreDtTm,omitempty"`
	QryNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 QryNm,omitempty"`
}

type MessageReference1 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Ref"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 MsgNm,omitempty"`
	RefIssr PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 RefIssr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Adr,omitempty"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 NmAndAdr,omitempty"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 LEI,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Ctry"`
}

type ReceiptAcknowledgementReport2 struct {
	RltdRef MessageReference1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 RltdRef"`
	ReqHdlg RequestHandling2  `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 ReqHdlg"`
}

type ReceiptAcknowledgementV01 struct {
	MsgId       MessageHeader10                 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 MsgId"`
	Rpt         []ReceiptAcknowledgementReport2 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Rpt"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 SplmtryData,omitempty"`
}

type RequestHandling2 struct {
	StsCd   Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 StsCd"`
	StsDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 StsDtTm,omitempty"`
	Desc    Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Envlp"`
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
