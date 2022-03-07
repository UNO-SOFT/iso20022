// Code generated by download. DO NOT EDIT.

package iso20022_seev_024_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgentCAInformationStatusAdviceV01 struct {
	Id               DocumentIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	AgtCAInfAdvcId   DocumentIdentification8                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AgtCAInfAdvcId"`
	CorpActnAddtlInf CorporateActionAdditionalInformation1   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CorpActnAddtlInf,omitempty"`
	InfStsDtls       CorporateActionInformationStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 InfStsDtls"`
}

type AlternateSecurityIdentification3 struct {
	Id         Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 PrtryIdSrc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

type BeneficialOwner1 struct {
	BnfclOwnrId    PartyIdentification2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 BnfclOwnrId"`
	AddtlId        GenericIdentification16                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AddtlId,omitempty"`
	Ntlty          CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Ntlty,omitempty"`
	DmclCtry       CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 DmclCtry,omitempty"`
	NonDmclCtry    CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 NonDmclCtry,omitempty"`
	CertfctnInd    bool                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CertfctnInd,omitempty"`
	CertfctnTp     BeneficiaryCertificationType1FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CertfctnTp,omitempty"`
	DclrtnDtls     Max350Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 DclrtnDtls,omitempty"`
	SctyId         SecurityIdentification7                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 SctyId,omitempty"`
	ElctdSctiesQty UnitOrFaceAmount1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 ElctdSctiesQty"`
}

// May be one of ACCI, DOMI, FULL, QIBB, TRBD, NCOM
type BeneficiaryCertificationType1Code string

type BeneficiaryCertificationType1FormatChoice struct {
	Cd    BeneficiaryCertificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Cd,omitempty"`
	Prtry GenericIdentification13           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Prtry,omitempty"`
}

type CashAccountIdentification1Choice struct {
	IBAN     IBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 IBAN,omitempty"`
	BBAN     BBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 BBAN,omitempty"`
	UPIC     UPICIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 UPIC,omitempty"`
	DmstAcct SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 DmstAcct,omitempty"`
}

type CorporateActionAdditionalInformation1 struct {
	BnfclOwnrDtls []BeneficialOwner1                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 BnfclOwnrDtls,omitempty"`
	RegnDtls      Max350Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 RegnDtls,omitempty"`
	RcvrId        PartyIdentification2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 RcvrId,omitempty"`
	CertfctnInd   bool                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CertfctnInd,omitempty"`
	CertfctnTp    BeneficiaryCertificationType1FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CertfctnTp,omitempty"`
	DlvryDtls     []ProceedsDelivery1                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 DlvryDtls,omitempty"`
	AddtlInstr    Max350Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AddtlInstr,omitempty"`
}

type CorporateActionInformationProcessingStatus1 struct {
	Sts      ProcessedStatus5FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AddtlInf,omitempty"`
}

type CorporateActionInformationRejectedStatus1 struct {
	Rsn      []RejectionReason15FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Rsn"`
	AddtlInf Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AddtlInf,omitempty"`
}

type CorporateActionInformationStatus1Choice struct {
	PrcdSts  CorporateActionInformationProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 PrcdSts,omitempty"`
	RjctdSts CorporateActionInformationRejectedStatus1   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 RjctdSts,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	AgtCAInfStsAdvc AgentCAInformationStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AgtCAInfStsAdvc"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CreDtTm,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Issr"`
}

type GenericIdentification16 struct {
	Id   Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
	IdTp PersonIdentificationType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 IdTp"`
	Issr Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Issr,omitempty"`
}

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

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
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 BICOrBEI,omitempty"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 NmAndAdr,omitempty"`
}

type PersonIdentificationType3Choice struct {
	Cd    PersonIdentificationType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Cd,omitempty"`
	Prtry GenericIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Prtry,omitempty"`
}

// May be one of ARNU, CCPT, EMID, DRLC, FINN, TXID
type PersonIdentificationType3Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Ctry"`
}

type ProceedsDelivery1 struct {
	SctiesAcctId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 SctiesAcctId,omitempty"`
	CshAcctId    CashAccountIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 CshAcctId,omitempty"`
	AcctOwnrId   PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AcctOwnrId,omitempty"`
	AcctSvcrId   PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 AcctSvcrId,omitempty"`
}

// May be one of RECE, PACK
type ProcessedStatus5Code string

type ProcessedStatus5FormatChoice struct {
	Cd    ProcessedStatus5Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Prtry,omitempty"`
}

// May be one of FAIL
type RejectionReason15Code string

type RejectionReason15FormatChoice struct {
	Cd    RejectionReason15Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Prtry,omitempty"`
}

type SecurityIdentification7 struct {
	ISIN   ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 ISIN,omitempty"`
	OthrId AlternateSecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 OthrId,omitempty"`
	Desc   Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Desc,omitempty"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Id"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

type UnitOrFaceAmount1Choice struct {
	Unit    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Unit,omitempty"`
	FaceAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 FaceAmt,omitempty"`
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
