// Code generated by download. DO NOT EDIT.

package iso20022_camt_059_001_07

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstlAdr,omitempty"`
}

type CashAccount40 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id,omitempty"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type Document struct {
	NtfctnToRcvStsRpt NotificationToReceiveStatusReportV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 NtfctnToRcvStsRpt"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Nb     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nb,omitempty"`
	RltdDt ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	Desc Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdOrPrtry"`
	Issr      Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalDiscountAmountType1Code string

// May be no more than 4 items long
type ExternalDocumentLineType1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalGarnishmentType1Code string

// May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// May be no more than 4 items long
type ExternalPersonIdentification1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalPurpose1Code string

// May be no more than 4 items long
type ExternalTaxAmountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Othr,omitempty"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RefNb,omitempty"`
	Dt                ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdOrPrtry"`
	Issr      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type GroupHeader84 struct {
	MsgId   Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 MsgId"`
	CreDtTm ISODateTime   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CreDtTm"`
	MsgRcpt Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 MsgRcpt,omitempty"`
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

type ISOYear time.Time

func (t *ISOYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISOYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be no more than 105 items long
type Max105Text string

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

type NameAndAddress16 struct {
	Nm  Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm"`
	Adr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Adr"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

// May be one of RCBD, RCVD, NRCD
type NotificationStatus3Code string

type NotificationToReceiveStatusReportV07 struct {
	GrpHdr            GroupHeader84          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 GrpHdr"`
	OrgnlNtfctnAndSts OriginalNotification13 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlNtfctnAndSts"`
	SplmtryData       []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SplmtryData,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type OriginalItemAndStatus7 struct {
	OrgnlItmId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlItmId"`
	OrgnlEndToEndId Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       UUIDv4Identifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlUETR,omitempty"`
	Amt             ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt"`
	XpctdValDt      ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 XpctdValDt,omitempty"`
	ItmSts          NotificationStatus3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ItmSts"`
	AddtlStsInf     Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AddtlStsInf,omitempty"`
	OrgnlItmRef     OriginalItemReference6            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlItmRef,omitempty"`
}

type OriginalItemReference6 struct {
	Acct       CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Acct,omitempty"`
	AcctOwnr   Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AcctOwnr,omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AcctSvcr,omitempty"`
	RltdAcct   CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RltdAcct,omitempty"`
	Dbtr       Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dbtr,omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DbtrAgt,omitempty"`
	IntrmyAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 IntrmyAgt,omitempty"`
	Purp       Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Purp,omitempty"`
	RltdRmtInf RemittanceLocation7                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RltdRmtInf,omitempty"`
	RmtInf     RemittanceInformation21                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtInf,omitempty"`
}

type OriginalNotification13 struct {
	OrgnlMsgId     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlMsgId"`
	OrgnlCreDtTm   ISODateTime                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlCreDtTm,omitempty"`
	OrgnlNtfctnId  Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlNtfctnId"`
	NtfctnSts      NotificationStatus3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 NtfctnSts,omitempty"`
	AddtlStsInf    Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AddtlStsInf,omitempty"`
	OrgnlNtfctnRef []OriginalNotificationReference11 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlNtfctnRef,omitempty"`
}

type OriginalNotificationReference11 struct {
	Acct           CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Acct,omitempty"`
	AcctOwnr       Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AcctOwnr,omitempty"`
	AcctSvcr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AcctSvcr,omitempty"`
	RltdAcct       CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RltdAcct,omitempty"`
	TtlAmt         ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TtlAmt,omitempty"`
	XpctdValDt     ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 XpctdValDt,omitempty"`
	Dbtr           Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dbtr,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DbtrAgt,omitempty"`
	IntrmyAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 IntrmyAgt,omitempty"`
	OrgnlItmAndSts []OriginalItemAndStatus7                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgnlItmAndSts"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PrvtId,omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Pty,omitempty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Agt,omitempty"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Nb       Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nb,omitempty"`
	RltdDt   ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Issr,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtdAmt,omitempty"`
}

type RemittanceInformation21 struct {
	Ustrd []Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation17 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Mtd"`
	ElctrncAdr Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PstlAdr,omitempty"`
}

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

type StructuredRemittanceInformation17 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Invcee,omitempty"`
	TaxRmt      TaxData1                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 GrnshmtRmt,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount3 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails3               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cd,omitempty"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prtry,omitempty"`
}

type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Titl,omitempty"`
	Nm   Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Nm,omitempty"`
}

type TaxData1 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 SeqNb,omitempty"`
	Rcrd            []TaxRecord3                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxId,omitempty"`
	RegnId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RegnId,omitempty"`
	TaxTp  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxId,omitempty"`
	RegnId  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Authstn,omitempty"`
}

type TaxPeriod3 struct {
	Yr     ISOYear              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FrToDt,omitempty"`
}

type TaxRecord3 struct {
	Tp       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Tp,omitempty"`
	Ctgy     Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 FrmsCd,omitempty"`
	Prd      TaxPeriod3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prd,omitempty"`
	TaxAmt   TaxAmount3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 AddtlInf,omitempty"`
}

type TaxRecordDetails3 struct {
	Prd TaxPeriod3                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.07 Amt"`
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

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

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006")
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
