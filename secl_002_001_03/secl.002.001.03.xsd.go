// Code generated by download. DO NOT EDIT.

package iso20022_secl_002_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification26 struct {
	Prtry SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternatePartyIdentification4 struct {
	IdTp    IdentificationType6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 IdTp"`
	Ctry    CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry"`
	AltrnId Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AltrnId"`
}

type AlternatePartyIdentification5 struct {
	IdTp    IdentificationType40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AltrnId"`
}

type AmountAndDirection21 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CdtDbtInd,omitempty"`
}

type AmountAndDirection27 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms17            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 FXDtls,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type Clearing4 struct {
	SttlmNetgElgblCd NettingEligible1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SttlmNetgElgblCd"`
	ClrSgmt          PartyIdentification35Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ClrSgmt,omitempty"`
	GrntedTrad       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 GrntedTrad,omitempty"`
	NonGrntedTrad    NonGuaranteedTrade3         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NonGrntedTrad,omitempty"`
}

// May be one of HOUS, CLIE, LIPR
type ClearingAccountType1Code string

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DtTm,omitempty"`
}

type DateCode3Choice struct {
	Cd    DateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

type DateFormat15Choice struct {
	Dt   ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Dt,omitempty"`
	DtCd DateCode3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DtCd,omitempty"`
}

// May be one of UKWN
type DateType1Code string

type DeliveringPartiesAndAccount11 struct {
	Dpstry         PartyIdentification34Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Dpstry"`
	Pty1           PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Pty1"`
	Pty2           PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Pty2,omitempty"`
	SctiesSttlmSys Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SctiesSttlmSys,omitempty"`
}

type Document struct {
	TradLegNtfctnCxl TradeLegNotificationCancellationV03 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradLegNtfctnCxl"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Unit,omitempty"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 FaceAmt,omitempty"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AmtsdVal,omitempty"`
}

type ForeignExchangeTerms17 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 RsltgAmt"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SchmeNm,omitempty"`
}

type GenericIdentification29 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SchmeNm,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SchmeNm,omitempty"`
}

type GenericIdentification40 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SchmeNm,omitempty"`
}

type GenericIdentification58 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id,omitempty"`
	Tp GenericIdentification40 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

type IdentificationType40Choice struct {
	Cd    TypeOfIdentification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry GenericIdentification29   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

type IdentificationType6Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MarketIdentification1Choice struct {
	MktIdrCd MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 MktIdrCd,omitempty"`
	Desc     Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Desc,omitempty"`
}

type MarketIdentification84 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id,omitempty"`
	Tp MarketType8Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp"`
}

type MarketIdentification85 struct {
	Id MarketIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id,omitempty"`
	Tp MarketType9Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp"`
}

// May be one of PRIM, SECM, OTCO, VARI, EXCH
type MarketType2Code string

// May be one of OTCO, EXCH
type MarketType5Code string

type MarketType8Choice struct {
	Cd    MarketType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

type MarketType9Choice struct {
	Cd    MarketType5Code         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

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

// May be no more than 70 items long
type Max70Text string

type NameAndAddress13 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm"`
	Adr PostalAddress8 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Adr,omitempty"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Adr,omitempty"`
}

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Adr"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

// May be one of GROS, NETT, AGFS
type NettingEligible1Code string

type NonGuaranteedTrade3 struct {
	TradCtrPtyMmbId    PartyIdentification35Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradCtrPtyMmbId"`
	TradCtrPtyClrMmbId PartyIdentification35Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradCtrPtyClrMmbId"`
	DlvrgPties         DeliveringPartiesAndAccount11 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DlvrgPties,omitempty"`
	RcvgPties          ReceivingPartiesAndAccount11  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 RcvgPties,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp"`
}

type PartyIdentification33Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AnyBIC,omitempty"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrtryId,omitempty"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NmAndAdr,omitempty"`
}

type PartyIdentification34Choice struct {
	BIC      AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 BIC,omitempty"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NmAndAdr,omitempty"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry,omitempty"`
}

type PartyIdentification35Choice struct {
	BIC     AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 BIC,omitempty"`
	PrtryId GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrtryId,omitempty"`
}

type PartyIdentification83Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AnyBIC,omitempty"`
	PrtryId  GenericIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrtryId,omitempty"`
	NmAndAdr NameAndAddress13        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NmAndAdr,omitempty"`
}

type PartyIdentificationAndAccount100 struct {
	Id        PartyIdentification83Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	AltrnId   AlternatePartyIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AltrnId,omitempty"`
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SfkpgAcct,omitempty"`
	PrcgId    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrcgId,omitempty"`
	AddtlInf  PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AddtlInf,omitempty"`
}

type PartyIdentificationAndAccount102 struct {
	PtyId    PartyIdentification33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PtyId"`
	AcctId   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AcctId,omitempty"`
	PrcgId   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrcgId,omitempty"`
	PrcgDt   DateAndDateTimeChoice       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PrcgDt,omitempty"`
	SubAcct  SubAccount4                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SubAcct,omitempty"`
	CtctPrsn ContactIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CtctPrsn,omitempty"`
}

type PartyIdentificationAndAccount31 struct {
	Id       PartyIdentification33Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	AltrnId  AlternatePartyIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AltrnId,omitempty"`
	AddtlInf PartyTextInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AddtlInf,omitempty"`
	ClrAcct  SecuritiesAccount18           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ClrAcct,omitempty"`
}

type PartyTextInformation1 struct {
	DclrtnDtls  Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DclrtnDtls,omitempty"`
	PtyCtctDtls Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PtyCtctDtls,omitempty"`
	RegnDtls    Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 RegnDtls,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry"`
}

type PostalAddress8 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry"`
}

type Price4 struct {
	Val PriceRateOrAmountChoice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Val"`
	Tp  PriceValueType7Code     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp,omitempty"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Rate,omitempty"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Amt,omitempty"`
}

// May be one of DISC, PREM, PARV, YIEL, SPRE, PEUN, ABSO, TEDP, TEDY, FICT, VACT, PRCT, ACTU
type PriceValueType7Code string

type ReceivingPartiesAndAccount11 struct {
	Dpstry         PartyIdentification34Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Dpstry"`
	Pty1           PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Pty1"`
	Pty2           PartyIdentificationAndAccount102 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Pty2,omitempty"`
	SctiesSttlmSys Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SctiesSttlmSys,omitempty"`
}

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat7Choice struct {
	Id      SafekeepingPlaceTypeAndText1             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id,omitempty"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Ctry,omitempty"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TpAndId,omitempty"`
	Prtry   GenericIdentification58                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Prtry,omitempty"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
}

type SafekeepingPlaceTypeAndText1 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id,omitempty"`
}

type SecuritiesAccount18 struct {
	Id Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Tp ClearingAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp"`
	Nm Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Desc,omitempty"`
}

type Settlement1 struct {
	SttlmAmt AmountAndDirection27        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SttlmAmt"`
	Dpstry   PartyIdentification34Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Dpstry,omitempty"`
}

// May be one of BUYI, SELL, TWOS, BUMI, SEPL, SESH, SSEX, CROS, CRSH, CSHE, DEFI, OPPO, UNDI
type Side1Code string

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
}

// May be one of REJT, PACK, PDNG
type Status5Code string

type SubAccount4 struct {
	Id    AccountIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Id"`
	Nm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Nm,omitempty"`
	Chrtc Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Chrtc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeLeg8 struct {
	TradLegId     Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradLegId"`
	TradId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradId,omitempty"`
	TradExctnId   Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradExctnId"`
	OrdrId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 OrdrId,omitempty"`
	AllcnId       Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AllcnId,omitempty"`
	Sts           Status5Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Sts,omitempty"`
	TradDt        ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradDt"`
	TxDtTm        ISODateTime                        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TxDtTm,omitempty"`
	SttlmDt       DateFormat15Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SttlmDt,omitempty"`
	FinInstrmId   SecurityIdentification14           `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 FinInstrmId"`
	TradgCcy      CurrencyCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradgCcy,omitempty"`
	BuySellInd    Side1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 BuySellInd"`
	TradQty       FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradQty"`
	DealPric      Price4                             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DealPric"`
	GrssAmt       AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 GrssAmt,omitempty"`
	AcrdIntrstAmt AmountAndDirection21               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 AcrdIntrstAmt,omitempty"`
	PlcOfTrad     MarketIdentification84             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PlcOfTrad"`
	PlcOfListg    MarketIdentification85             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 PlcOfListg,omitempty"`
	TradTp        TradeType1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradTp"`
	DerivRltdTrad bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DerivRltdTrad,omitempty"`
	Brkr          PartyIdentificationAndAccount100   `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Brkr,omitempty"`
	TradgPty      PartyIdentification35Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradgPty"`
	TradRegnOrgn  Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradRegnOrgn,omitempty"`
	TradgPtyAcct  SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradgPtyAcct,omitempty"`
	TradgCpcty    TradingCapacity5Code               `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradgCpcty"`
	TradPstngCd   TradePosting1Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradPstngCd,omitempty"`
	SfkpgPlc      SafekeepingPlaceFormat7Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SfkpgPlc,omitempty"`
	SfkpgAcct     SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SfkpgAcct,omitempty"`
}

type TradeLegNotificationCancellationV03 struct {
	ClrMmb      PartyIdentification35Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ClrMmb"`
	ClrAcct     SecuritiesAccount18             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ClrAcct"`
	DlvryAcct   SecuritiesAccount19             `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 DlvryAcct,omitempty"`
	NonClrMmb   PartyIdentificationAndAccount31 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 NonClrMmb,omitempty"`
	ClrDtls     Clearing4                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 ClrDtls,omitempty"`
	TradLegDtls TradeLeg8                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 TradLegDtls"`
	SttlmDtls   Settlement1                     `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SttlmDtls"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 SplmtryData,omitempty"`
}

// May be one of GROS, NETT
type TradePosting1Code string

// May be one of OOBK, OFBK, BKTR, COTR, GUTR, LKTR
type TradeType1Code string

// May be one of PRIN, RISP, AGEN
type TradingCapacity5Code string

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

// May be one of ARNU, CHTY, CORP, FIIN, TXID
type TypeOfIdentification2Code string

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
