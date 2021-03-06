// Code generated by download. DO NOT EDIT.

package iso20022_camt_046_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	GetRsvatn GetReservationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 GetRsvatn"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalEnquiryRequestType1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 3 items long
type ExternalMarketInfrastructure1Code string

// May be no more than 4 items long
type ExternalPaymentControlRequestType1Code string

// May be no more than 4 items long
type ExternalReservationType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SchmeNm,omitempty"`
}

type GetReservationV06 struct {
	MsgHdr       MessageHeader9       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 MsgHdr"`
	RsvatnQryDef ReservationQuery4    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 RsvatnQryDef,omitempty"`
	SplmtryData  []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SplmtryData,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 70 items long
type Max70Text string

type MessageHeader9 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 CreDtTm,omitempty"`
	ReqTp   RequestType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 ReqTp,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 AdrLine,omitempty"`
}

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PmtCtrl,omitempty"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Enqry,omitempty"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

type ReservationCriteria4Choice struct {
	QryNm   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 QryNm,omitempty"`
	NewCrit ReservationCriteria5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 NewCrit,omitempty"`
}

type ReservationCriteria5 struct {
	NewQryNm Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 NewQryNm,omitempty"`
	SchCrit  []ReservationSearchCriteria4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SchCrit,omitempty"`
	RtrCrit  ReservationReturnCriteria1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 RtrCrit,omitempty"`
}

type ReservationQuery4 struct {
	QryTp      QueryType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 QryTp,omitempty"`
	RsvatnCrit ReservationCriteria4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 RsvatnCrit,omitempty"`
}

type ReservationReturnCriteria1 struct {
	StartDtTmInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 StartDtTmInd,omitempty"`
	StsInd       bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 StsInd,omitempty"`
}

type ReservationSearchCriteria4 struct {
	SysId        SystemIdentification2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 SysId,omitempty"`
	DfltRsvatnTp []ReservationType2Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 DfltRsvatnTp,omitempty"`
	CurRsvatnTp  []ReservationType2Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 CurRsvatnTp,omitempty"`
	AcctOwnr     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 AcctOwnr,omitempty"`
	AcctId       AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 AcctId,omitempty"`
}

type ReservationType2Choice struct {
	Cd    ExternalReservationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Prtry,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 MktInfrstrctrId,omitempty"`
	Ctry            CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.06 Ctry,omitempty"`
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
