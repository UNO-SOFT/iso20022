// Code generated by download. DO NOT EDIT.

package iso20022_acmt_013_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountForAction1 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	Ccy ActiveCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Ccy"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Othr,omitempty"`
}

type AccountReportRequestV03 struct {
	Refs        References4                                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Refs"`
	Fr          OrganisationIdentification29                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Fr,omitempty"`
	AcctId      []AccountForAction1                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AcctId"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AcctSvcrId"`
	OrgId       OrganisationIdentification29                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 OrgId"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 DgtlSgntr,omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CtryOfBirth"`
}

type Document struct {
	AcctRptReq AccountReportRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AcctRptReq"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// May be no more than 4 items long
type ExternalPersonIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

// May be no more than 128 items long
type Max128Text string

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

// May be no more than 4 items long
type Max4Text string

// May be no more than 70 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CreDtTm"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PrvtId,omitempty"`
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Pty"`
	Sgntr SkipPayload            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Sgntr"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type References4 struct {
	MsgId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 MsgId"`
	PrcId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PrcId"`
	AttchdDocNm []Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AttchdDocNm,omitempty"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
