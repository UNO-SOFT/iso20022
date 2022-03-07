// Code generated by download. DO NOT EDIT.

package iso20022_reda_015_001_01

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

type DateAndDateTimeSearch4Choice struct {
	DtTm DateTimeSearch2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 DtTm,omitempty"`
	Dt   DatePeriodSearch1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Dt,omitempty"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrDt,omitempty"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ToDt,omitempty"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrToDt,omitempty"`
	EQDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 EQDt,omitempty"`
	NEQDt  ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 NEQDt,omitempty"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ToDtTm"`
}

type DateTimeSearch2Choice struct {
	FrDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrDtTm,omitempty"`
	ToDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ToDtTm,omitempty"`
	FrToDtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 FrToDtTm,omitempty"`
	EQDtTm   ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 EQDtTm,omitempty"`
	NEQDtTm  ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 NEQDtTm,omitempty"`
}

type Document struct {
	PtyQry PartyQueryV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PtyQry"`
}

// May be no more than 4 items long
type ExternalSystemPartyType1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Issr,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 SchmeNm,omitempty"`
}

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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be one of LOCK, ULCK
type LockStatus1Code string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 70 items long
type Max70Text string

type MessageHeader2 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 CreDtTm,omitempty"`
	ReqTp   RequestType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ReqTp,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Adr,omitempty"`
}

type PartyDataReturnCriteria2 struct {
	OpngDt       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 OpngDt,omitempty"`
	ClsgDt       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ClsgDt,omitempty"`
	Tp           bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Tp,omitempty"`
	PtyId        bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PtyId,omitempty"`
	RspnsblPtyId bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RspnsblPtyId,omitempty"`
	RstrctnId    bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RstrctnId,omitempty"`
	RstrctdOnDt  bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RstrctdOnDt,omitempty"`
	Nm           bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Nm,omitempty"`
	ShrtNm       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ShrtNm,omitempty"`
	Adr          bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Adr,omitempty"`
	TechAdr      bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 TechAdr,omitempty"`
	CtctDtls     bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 CtctDtls,omitempty"`
	ResTp        bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ResTp,omitempty"`
	LckSts       bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 LckSts,omitempty"`
	MktSpcfcAttr bool `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 MktSpcfcAttr,omitempty"`
}

type PartyDataSearchCriteria2 struct {
	OpngDt        DatePeriodSearch1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 OpngDt,omitempty"`
	ClsgDt        DatePeriodSearch1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ClsgDt,omitempty"`
	Tp            SystemPartyType1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Tp,omitempty"`
	RspnsblPtyId  PartyIdentification136       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RspnsblPtyId,omitempty"`
	PtyId         PartyIdentification136       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PtyId,omitempty"`
	RstrctnId     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RstrctnId,omitempty"`
	RstrctnIsseDt DateAndDateTimeSearch4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RstrctnIsseDt,omitempty"`
	ResTp         ResidenceType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 ResTp,omitempty"`
	LckSts        PartyLockStatus1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 LckSts,omitempty"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 NmAndAdr,omitempty"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 LEI,omitempty"`
}

type PartyLockStatus1 struct {
	VldFr  ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 VldFr,omitempty"`
	Sts    LockStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Sts"`
	LckRsn []Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 LckRsn,omitempty"`
}

type PartyQueryV01 struct {
	MsgHdr      MessageHeader2           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 MsgHdr,omitempty"`
	SchCrit     PartyDataSearchCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 SchCrit"`
	RtrCrit     PartyDataReturnCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 RtrCrit,omitempty"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 SplmtryData,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Ctry"`
}

// May be one of RT01, RT02, RT03, RT04, RT05
type RequestType1Code string

type RequestType2Choice struct {
	PmtCtrl RequestType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PmtCtrl,omitempty"`
	Enqry   RequestType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Enqry,omitempty"`
	Prtry   GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Prtry,omitempty"`
}

// May be one of RT11, RT12, RT13, RT14, RT15
type RequestType2Code string

// May be one of DMST, FRGN, MXED
type ResidenceType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyType1Choice struct {
	Cd    ExternalSystemPartyType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.015.001.01 Prtry,omitempty"`
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
