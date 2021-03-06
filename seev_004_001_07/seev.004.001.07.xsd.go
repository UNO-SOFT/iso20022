// Code generated by download. DO NOT EDIT.

package iso20022_seev_004_001_07

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type AttendanceCard3 struct {
	AttndncCardLbllg Max105Text         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AttndncCardLbllg,omitempty"`
	DlvryMtd         DeliveryPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 DlvryMtd"`
	OthrAdr          NameAndAddress9    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 OthrAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth2 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 CityOfBirth,omitempty"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 CtryOfBirth,omitempty"`
}

// May be one of EMAL, EMPL, INDI, ENTR, OADR
type DeliveryPlace3Code string

type Document struct {
	MtgInstr MeetingInstructionV07 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 MtgInstr"`
}

type DocumentIdentification32 struct {
	Id    DocumentIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	DocNb DocumentNumber5Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 DocNb,omitempty"`
	LkgTp ProcessingPosition7Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LkgTp,omitempty"`
}

type DocumentIdentification3Choice struct {
	AcctSvcrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AcctSvcrDocId,omitempty"`
	AcctOwnrDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AcctOwnrDocId,omitempty"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ShrtNb,omitempty"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LngNb,omitempty"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrtryNb,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity18Choice struct {
	Unit    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Unit,omitempty"`
	FaceAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 FaceAmt,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Issr"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
}

type HoldingBalance10 struct {
	Bal      UnitOrFaceAmountOrCode2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Bal"`
	BalTp    SecuritiesEntryType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 BalTp,omitempty"`
	SfkpgPlc SafekeepingPlaceFormat28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SfkpgPlc,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

// Must match the pattern [a-z]{4}\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}
type ISO20022MessageIdentificationText string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type IdentificationType45Choice struct {
	Cd    TypeOfIdentification4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type IndividualPerson41 struct {
	Id              PartyIdentification232Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
	PrtcptnMtd      VotingParticipationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrtcptnMtd,omitempty"`
	EmplngPty       PartyIdentification129Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmplngPty,omitempty"`
	AttndncCardDtls AttendanceCard3                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AttndncCardDtls"`
}

type IndividualPerson42 struct {
	PrssgndPrxy     PartyIdentification232Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrssgndPrxy,omitempty"`
	EmplngPty       PartyIdentification129Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmplngPty,omitempty"`
	AttndncCardDtls AttendanceCard3              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AttndncCardDtls"`
}

type Instruction5 struct {
	SnglInstrId   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SnglInstrId"`
	ReqdExctnDt   ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ReqdExctnDt,omitempty"`
	VoteExctnConf bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteExctnConf"`
	AcctDtls      SafekeepingAccount12        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AcctDtls"`
	Prxy          Proxy10                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prxy,omitempty"`
	VoteDtls      VoteDetails5                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteDtls,omitempty"`
	MtgAttndee    []IndividualPerson41        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 MtgAttndee,omitempty"`
	SpcfcInstrReq SpecificInstructionRequest3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SpcfcInstrReq,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LongPostalAddress2Choice struct {
	Ustrd Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ustrd,omitempty"`
	Strd  PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Strd,omitempty"`
}

// May be no more than 105 items long
type Max105Text string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 256 items long
type Max256Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 50 items long
type Max50Text string

// May be no more than 70 items long
type Max70Text string

type MeetingInstructionV07 struct {
	MtgRef      MeetingReference10         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 MtgRef"`
	FinInstrmId SecurityIdentification19   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 FinInstrmId"`
	OthrDocId   []DocumentIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 OthrDocId,omitempty"`
	Instr       []Instruction5             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Instr"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SplmtryData,omitempty"`
}

type MeetingReference10 struct {
	MtgId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 MtgId"`
	IssrMtgId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 IssrMtgId,omitempty"`
	MtgDtAndTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 MtgDtAndTm"`
	Tp         MeetingType4Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Tp"`
	Clssfctn   MeetingTypeClassification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Clssfctn,omitempty"`
	Lctn       []PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Lctn,omitempty"`
	Issr       PartyIdentification129Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Issr,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL, BMET, CMET
type MeetingType4Code string

type MeetingTypeClassification2Choice struct {
	Cd    MeetingTypeClassification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

// May be one of AMET, CLAS, ISSU, OMET, VRHI
type MeetingTypeClassification2Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Adr,omitempty"`
}

type NameAndAddress9 struct {
	Nm  Max350Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Nm"`
	Adr LongPostalAddress2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Adr,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type NaturalPersonIdentification1 struct {
	Id   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	IdTp IdentificationType45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 IdTp,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Tp"`
}

type ParticipationMethod1Choice struct {
	Cd    VotingParticipationMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type PartyIdentification129Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmAndAdr,omitempty"`
	LEI      LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LEI,omitempty"`
}

type PartyIdentification198Choice struct {
	NtlRegnNb Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NtlRegnNb,omitempty"`
	LEI       LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LEI,omitempty"`
	AnyBIC    AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AnyBIC,omitempty"`
	ClntId    Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ClntId,omitempty"`
	PrtryId   GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrtryId,omitempty"`
}

type PartyIdentification221 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmailAdr,omitempty"`
	Id       PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
}

type PartyIdentification224 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmailAdr,omitempty"`
	Id       PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
}

type PartyIdentification231Choice struct {
	LglPrsn  PartyIdentification221   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LglPrsn,omitempty"`
	NtrlPrsn []PartyIdentification238 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NtrlPrsn,omitempty"`
}

type PartyIdentification232Choice struct {
	LglPrsn  PartyIdentification221 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LglPrsn,omitempty"`
	NtrlPrsn PartyIdentification238 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NtrlPrsn,omitempty"`
}

type PartyIdentification233Choice struct {
	LglPrsn  PartyIdentification224   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LglPrsn,omitempty"`
	NtrlPrsn []PartyIdentification240 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NtrlPrsn,omitempty"`
}

type PartyIdentification238 struct {
	NmAndAdr        PersonName3                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmAndAdr"`
	EmailAdr        Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmailAdr,omitempty"`
	Id              NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
	Ntlty           CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ntlty,omitempty"`
	DtAndPlcOfBirth DateAndPlaceOfBirth2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 DtAndPlcOfBirth,omitempty"`
}

type PartyIdentification240 struct {
	NmAndAdr        PersonName3                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmAndAdr"`
	EmailAdr        Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 EmailAdr,omitempty"`
	Id              NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
	Ntlty           CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ntlty,omitempty"`
	DtAndPlcOfBirth DateAndPlaceOfBirth2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 DtAndPlcOfBirth,omitempty"`
}

// May be one of GATR
type PartyRole3Code string

type PersonName2 struct {
	Nm  Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Nm"`
	Adr PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Adr,omitempty"`
}

type PersonName3 struct {
	NmPrfx NamePrefix2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NmPrfx,omitempty"`
	FrstNm Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 FrstNm"`
	Srnm   Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Srnm"`
	Adr    PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Adr,omitempty"`
}

type PledgeInformation1 struct {
	Pldgr        PartyIdentification232Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Pldgr"`
	ThrdPty      ThirdPartyIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ThrdPty,omitempty"`
	PldgTp       GenericIdentification36      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PldgTp"`
	RtrSctiesInd bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 RtrSctiesInd,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ctry"`
}

type PostalAddress26 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 BldgNb,omitempty"`
	PstBx       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PstBx,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ctry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type ProprietaryVote1 struct {
	Cd  GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd"`
	Qty QuantityOrCode1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Qty"`
}

type Proxy10 struct {
	PrxyTp   ProxyType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrxyTp"`
	PrsnDtls IndividualPerson42 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrsnDtls,omitempty"`
}

// May be one of CHRM, DISC, HLDR
type ProxyType2Code string

// May be one of QALL
type Quantity1Code string

type QuantityOrCode1Choice struct {
	Qty FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Qty,omitempty"`
	Cd  Quantity1Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
}

type SafekeepingAccount12 struct {
	AcctId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AcctId"`
	AcctOwnr  PartyIdentification231Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AcctOwnr,omitempty"`
	SubAcctId Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SubAcctId,omitempty"`
	InstdBal  []HoldingBalance10             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 InstdBal"`
	RghtsHldr []PartyIdentification233Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 RghtsHldr,omitempty"`
	PldgDtls  PledgeInformation1             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PldgDtls,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Ctry,omitempty"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 TpAndId,omitempty"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Id,omitempty"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Desc,omitempty"`
}

type SpecificInstructionRequest3 struct {
	PrtcptnMtd ParticipationMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PrtcptnMtd,omitempty"`
	SctiesRegn bool                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 SctiesRegn,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type ThirdPartyIdentification1 struct {
	Role      PartyRole3Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Role"`
	LglPrsnId PartyIdentification221 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 LglPrsnId,omitempty"`
}

// May be one of ARNU, CUST, CORP, DRLC, IDCD, NRIN, CCPT, SOCS, TXID
type TypeOfIdentification4Code string

type UnitOrFaceAmountOrCode2Choice struct {
	Unit    float64       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Unit,omitempty"`
	FaceAmt float64       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 FaceAmt,omitempty"`
	Cd      Quantity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Cd,omitempty"`
}

type Vote14 struct {
	IssrLabl          Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 IssrLabl"`
	ListgGrpRsltnLabl Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ListgGrpRsltnLabl,omitempty"`
	For               QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 For,omitempty"`
	Agnst             QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Agnst,omitempty"`
	Abstn             QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Abstn,omitempty"`
	Wthhld            QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Wthhld,omitempty"`
	WthMgmt           QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 WthMgmt,omitempty"`
	AgnstMgmt         QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 AgnstMgmt,omitempty"`
	Dscrtnry          QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Dscrtnry,omitempty"`
	OneYr             QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 OneYr,omitempty"`
	TwoYrs            QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 TwoYrs,omitempty"`
	ThreeYrs          QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 ThreeYrs,omitempty"`
	NoActn            QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 NoActn,omitempty"`
	Blnk              QuantityOrCode1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Blnk,omitempty"`
	Prtry             []ProprietaryVote1    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type Vote15 struct {
	IssrLabl Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 IssrLabl"`
	VoteOptn VoteInstructionType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteOptn"`
}

type Vote15Choice struct {
	VotePerAgndRsltn     Vote16Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VotePerAgndRsltn,omitempty"`
	VoteForAllAgndRsltns VoteInstructionType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteForAllAgndRsltns,omitempty"`
}

type Vote16Choice struct {
	VoteInstr    []Vote14 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteInstr,omitempty"`
	GblVoteInstr []Vote15 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 GblVoteInstr,omitempty"`
}

type VoteDetails5 struct {
	VoteInstrForAgndRsltn Vote15Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteInstrForAgndRsltn"`
	VoteInstrForMtgRsltn  VoteInstructionForMeetingResolution3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteInstrForMtgRsltn,omitempty"`
}

// May be one of ABST, CAGS, AMGT, BLNK, CHRM, DISC, CFOR, NOAC, ONEY, THRY, TWOY, WTHH, WMGT
type VoteInstruction6Code string

// May be one of ABST, CAGS, AMGT, BLNK, CFOR, NOAC, ONEY, THRY, TWOY, WTHH, WMGT
type VoteInstruction7Code string

type VoteInstructionForMeetingResolution3Choice struct {
	VoteIndctn VoteInstructionType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 VoteIndctn,omitempty"`
	Shrhldr    NameAndAddress9            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Shrhldr,omitempty"`
}

type VoteInstructionType1Choice struct {
	Tp    VoteInstruction6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Tp,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

type VoteInstructionType2Choice struct {
	Tp    VoteInstruction7Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Tp,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.07 Prtry,omitempty"`
}

// May be one of MAIL, PHYS, PHNV, PRXY, VIRT, EVOT
type VotingParticipationMethod1Code string

// May be one of PHNV
type VotingParticipationMethod2Code string

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
