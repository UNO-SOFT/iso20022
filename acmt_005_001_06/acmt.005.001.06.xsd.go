// Code generated by download. DO NOT EDIT.

package iso20022_acmt_005_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type CashAccount40 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Id,omitempty"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MmbId"`
}

type CommunicationAddress10 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 PstlAdr"`
	PhneNb   PhoneNumber              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 PhneNb"`
	FaxNb    PhoneNumber              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 FaxNb,omitempty"`
	EmailAdr Max2048Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 EmailAdr,omitempty"`
}

type ContactIdentificationAndAddress2 struct {
	Nm     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Nm,omitempty"`
	Role   PaymentRole1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Role"`
	ComAdr CommunicationAddress10 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 ComAdr"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	RtrMmb ReturnMemberV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 RtrMmb"`
}

type ErrorHandling1Choice struct {
	Cd    ErrorHandling1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

// May be one of X020, X030, X050
type ErrorHandling1Code string

type ErrorHandling3 struct {
	Err  ErrorHandling1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Desc,omitempty"`
}

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalEnquiryRequestType1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalPaymentControlRequestType1Code string

// May be no more than 4 items long
type ExternalPaymentRole1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalSystemMemberType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LongPostalAddress1Choice struct {
	Ustrd Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Ustrd,omitempty"`
	Strd  StructuredLongPostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Strd,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 70 items long
type Max70Text string

type Member7 struct {
	Nm      Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Nm,omitempty"`
	RtrAdr  []MemberIdentification3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 RtrAdr,omitempty"`
	Acct    []CashAccount40                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Acct,omitempty"`
	Tp      SystemMemberType1Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Tp,omitempty"`
	Sts     SystemMemberStatus1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Sts,omitempty"`
	CtctRef []ContactIdentificationAndAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 CtctRef,omitempty"`
	ComAdr  CommunicationAddress10             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 ComAdr,omitempty"`
}

type MemberIdentification3Choice struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 ClrSysMmbId,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Othr,omitempty"`
}

type MemberReport6 struct {
	MmbId    MemberIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MmbId"`
	MmbOrErr MemberReportOrError8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MmbOrErr"`
}

type MemberReportOrError7Choice struct {
	Rpt     []MemberReport6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Rpt,omitempty"`
	OprlErr []ErrorHandling3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 OprlErr,omitempty"`
}

type MemberReportOrError8Choice struct {
	Mmb    Member7        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Mmb,omitempty"`
	BizErr ErrorHandling3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 BizErr,omitempty"`
}

// May be one of ENBL, DSBL, DLTD, JOIN
type MemberStatus1Code string

type MessageHeader7 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 CreDtTm,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 ReqTp,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 QryNm,omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 CreDtTm,omitempty"`
}

type PaymentRole1Choice struct {
	Cd    ExternalPaymentRole1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 PmtCtrl,omitempty"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Enqry,omitempty"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type ReturnMemberV05 struct {
	MsgHdr      MessageHeader7             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 MsgHdr"`
	RptOrErr    MemberReportOrError7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 RptOrErr"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 SplmtryData,omitempty"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 BldgNm,omitempty"`
	StrtNm     Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 StrtNm,omitempty"`
	StrtBldgId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 StrtBldgId,omitempty"`
	Flr        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Flr,omitempty"`
	TwnNm      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 TwnNm"`
	DstrctNm   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 DstrctNm,omitempty"`
	RgnId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 RgnId,omitempty"`
	Stat       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Stat,omitempty"`
	CtyId      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 CtyId,omitempty"`
	Ctry       CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Ctry"`
	PstCdId    Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 PstCdId"`
	POB        Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 POB,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemMemberStatus1Choice struct {
	Cd    MemberStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
}

type SystemMemberType1Choice struct {
	Cd    ExternalSystemMemberType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.05 Prtry,omitempty"`
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