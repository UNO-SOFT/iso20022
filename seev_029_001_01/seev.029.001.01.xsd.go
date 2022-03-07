// Code generated by download. DO NOT EDIT.

package iso20022_seev_029_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgentCADeactivationCancellationRequestV01 struct {
	Id                   DocumentIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Id"`
	AgtCADeactvtnInstrId DocumentIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 AgtCADeactvtnInstrId"`
	CorpActnGnlInf       CorporateActionInformation1             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 CorpActnGnlInf"`
	DeactvtnInstrDtls    CorporateActionDeactivationInstruction1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 DeactvtnInstrDtls,omitempty"`
}

type AlternateSecurityIdentification3 struct {
	Id         Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 PrtryIdSrc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateActionDeactivationInstruction1 struct {
	DeactvtnDtAndTm ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 DeactvtnDtAndTm"`
	OptnDtls        []CorporateActionOption2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 OptnDtls,omitempty"`
}

// May be one of GENL, DISN, REOR
type CorporateActionEventProcessingType1Code string

type CorporateActionEventProcessingType1FormatChoice struct {
	Cd    CorporateActionEventProcessingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Cd,omitempty"`
	Prtry GenericIdentification13                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Prtry,omitempty"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, INCR, INTR, LIQU, MCAL, MRGR, ODLT, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, OTHR
type CorporateActionEventType2Code string

type CorporateActionEventType2FormatChoice struct {
	Cd    CorporateActionEventType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Cd,omitempty"`
	Prtry GenericIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Prtry,omitempty"`
}

type CorporateActionInformation1 struct {
	AgtId             PartyIdentification2Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 AgtId"`
	IssrCorpActnId    Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 IssrCorpActnId,omitempty"`
	CorpActnPrcgId    Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 CorpActnPrcgId,omitempty"`
	EvtTp             CorporateActionEventType2FormatChoice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 EvtTp"`
	MndtryVlntryEvtTp CorporateActionMandatoryVoluntary1FormatChoice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 MndtryVlntryEvtTp"`
	EvtPrcgTp         CorporateActionEventProcessingType1FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 EvtPrcgTp,omitempty"`
	UndrlygScty       FinancialInstrumentDescription3                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 UndrlygScty"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary1FormatChoice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Cd,omitempty"`
	Prtry GenericIdentification13                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Prtry,omitempty"`
}

type CorporateActionOption1FormatChoice struct {
	Cd    CorporateActionOptionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Cd,omitempty"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Prtry,omitempty"`
}

type CorporateActionOption2 struct {
	OptnTp CorporateActionOption1FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 OptnTp"`
	OptnNb Exact3NumericText                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 OptnNb"`
}

// May be one of BSPL, BUYA, CASE, CASH, CEXC, CTEN, CONN, CONY, EXER, LAPS, MPUT, NOAC, OFFR, OVER, SECU, SLLE, SPLI, NOQU, OTHR, QINV
type CorporateActionOptionType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	AgtCADeactvtnCxlReq AgentCADeactivationCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 AgtCADeactvtnCxlReq"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 CreDtTm,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type FinancialInstrumentDescription3 struct {
	SctyId     SecurityIdentification7    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 SctyId"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 PlcOfListg,omitempty"`
	SfkpgPlc   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 SfkpgPlc,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Issr"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

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

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 BICOrBEI,omitempty"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 NmAndAdr,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Ctry"`
}

type SecurityIdentification7 struct {
	ISIN   ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 ISIN,omitempty"`
	OthrId AlternateSecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 OthrId,omitempty"`
	Desc   Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Desc,omitempty"`
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
