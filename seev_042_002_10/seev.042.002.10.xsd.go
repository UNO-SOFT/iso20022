// Code generated by download. DO NOT EDIT.

package iso20022_seev_042_002_10

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification57 struct {
	SfkpgAcct         RestrictedFINXMax35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SfkpgAcct"`
	AcctOwnr          PartyIdentification136Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AcctOwnr,omitempty"`
	SfkpgPlc          SafekeepingPlaceFormat32Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SfkpgPlc,omitempty"`
	CorpActnEvtAndBal []CorporateActionEventAndBalance21 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CorpActnEvtAndBal,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AmountPrice5 struct {
	AmtPricTp AmountPriceType1Code                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AmtPricTp"`
	PricVal   RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PricVal"`
}

// May be one of ACTU, DISC, PLOT, PREM
type AmountPriceType1Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type BalanceFormat7Choice struct {
	Bal         SignedQuantityFormat8 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Bal,omitempty"`
	ElgblBal    SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ElgblBal,omitempty"`
	NotElgblBal SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NotElgblBal,omitempty"`
}

type CancelledReason11Choice struct {
	Cd    CancelledStatusReason6Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type CancelledStatus15Choice struct {
	NoSpcfdRsn NoReasonCode              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NoSpcfdRsn,omitempty"`
	Rsn        []CancelledStatusReason14 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Rsn,omitempty"`
}

type CancelledStatusReason14 struct {
	RsnCd       CancelledReason11Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RsnCd"`
	AddtlRsnInf RestrictedFINXMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AddtlRsnInf,omitempty"`
}

// May be one of CANI, CANO, CANS, CSUB, OTHR
type CancelledStatusReason6Code string

type CorporateActionBalance45 struct {
	TtlElgblBal      Quantity22Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlElgblBal"`
	UinstdBal        BalanceFormat7Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 UinstdBal"`
	TtlInstdBalDtls  InstructedBalance15   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlInstdBalDtls"`
	BlckdBal         SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 BlckdBal,omitempty"`
	BrrwdBal         SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 BrrwdBal,omitempty"`
	CollInBal        SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CollInBal,omitempty"`
	CollOutBal       SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CollOutBal,omitempty"`
	OnLnBal          SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OnLnBal,omitempty"`
	OutForRegnBal    SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OutForRegnBal,omitempty"`
	SttlmPosBal      SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SttlmPosBal,omitempty"`
	StrtPosBal       SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StrtPosBal,omitempty"`
	TradDtPosBal     SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TradDtPosBal,omitempty"`
	InTrnsShipmntBal SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InTrnsShipmntBal,omitempty"`
	RegdBal          SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RegdBal,omitempty"`
	OblgtdBal        SignedQuantityFormat9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OblgtdBal,omitempty"`
	PdgDlvryBal      []PendingBalance6     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PdgDlvryBal,omitempty"`
	PdgRctBal        []PendingBalance6     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PdgRctBal,omitempty"`
}

type CorporateActionEventAndBalance21 struct {
	GnlInf      EventInformation16       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 GnlInf"`
	UndrlygScty SecurityIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 UndrlygScty"`
	Bal         CorporateActionBalance45 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Bal,omitempty"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SplmtryData,omitempty"`
}

type CorporateActionEventDeadlines4 struct {
	EarlyRspnDdln  DateFormat49Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 EarlyRspnDdln,omitempty"`
	RspnDdln       DateFormat54Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RspnDdln,omitempty"`
	MktDdln        DateFormat49Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 MktDdln,omitempty"`
	PrtctDdln      DateFormat49Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtctDdln,omitempty"`
	CoverPrtctDdln DateFormat49Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CoverPrtctDdln,omitempty"`
}

type CorporateActionEventType103Choice struct {
	Cd    CorporateActionEventType34Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, ACCU
type CorporateActionEventType34Code string

type CorporateActionInstructionStatementReport002V10 struct {
	Pgntn           Pagination1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Pgntn"`
	StmtGnlDtls     Statement75               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StmtGnlDtls"`
	AcctAndStmtDtls []AccountIdentification57 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AcctAndStmtDtls"`
	SplmtryData     []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SplmtryData,omitempty"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary4Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

// May be one of ABST, BSPL, BUYA, CASE, CASH, CEXC, CONN, CONY, CTEN, EXER, LAPS, MPUT, NOAC, NOQU, OFFR, OTHR, OVER, QINV, SECU, SLLE, PRUN
type CorporateActionOption11Code string

type CorporateActionOption36Choice struct {
	Cd    CorporateActionOption11Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

// May be one of MASE, SAME
type CorporateActionStatementReportingType1Code string

// May be one of MISS, ALLL, BALO, BALI
type CorporateActionStatementType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtTm,omitempty"`
}

type DateCode22Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type DateCode26Choice struct {
	Cd    DateType7Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type DateCodeAndTimeFormat4 struct {
	DtCd DateCode26Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtCd"`
	Tm   ISOTime          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Tm"`
}

type DateFormat49Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Dt,omitempty"`
	DtCd DateCode22Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtCd,omitempty"`
}

type DateFormat54Choice struct {
	Dt        DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Dt,omitempty"`
	DtCdAndTm DateCodeAndTimeFormat4 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtCdAndTm,omitempty"`
	DtCd      DateCode22Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtCd,omitempty"`
}

type DateOrDateTimePeriod1Choice struct {
	Dt   DatePeriod2     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Dt,omitempty"`
	DtTm DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DtTm,omitempty"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ToDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ToDtTm"`
}

// May be one of ONGO
type DateType7Code string

// May be one of UKWN, ONGO
type DateType8Code string

type DefaultProcessingOrStandingInstruction1Choice struct {
	DfltOptnInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DfltOptnInd,omitempty"`
	StgInstrInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StgInstrInd,omitempty"`
}

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	CorpActnInstrStmtRpt CorporateActionInstructionStatementReport002V10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CorpActnInstrStmtRpt"`
}

// May be one of YEAR, ADHO, MNTH, DAIL, INDA, WEEK
type EventFrequency4Code string

type EventInformation16 struct {
	CorpActnEvtId      RestrictedFINXMax16Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CorpActnEvtId"`
	OffclCorpActnEvtId RestrictedFINXMax16Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OffclCorpActnEvtId,omitempty"`
	EvtTp              CorporateActionEventType103Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 MndtryVlntryEvtTp"`
	LastNtfctnId       NotificationIdentification6              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 LastNtfctnId,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity15Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Unit,omitempty"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 FaceAmt,omitempty"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AmtsdVal,omitempty"`
}

type Frequency26Choice struct {
	Cd    EventFrequency4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SchmeNm,omitempty"`
}

type GenericIdentification84 struct {
	Id      RestrictedFINXMax34Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id"`
	Issr    Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Issr"`
	SchmeNm Max4AlphaNumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SchmeNm,omitempty"`
}

type GenericIdentification85 struct {
	Tp GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Tp"`
	Id RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type IdentificationSource4Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry RestrictedFINExact2Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type InstructedBalance15 struct {
	TtlInstdBal       BalanceFormat7Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlInstdBal"`
	TtlAccptdInstrBal SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlAccptdInstrBal,omitempty"`
	TtlCancInstrBal   SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlCancInstrBal,omitempty"`
	TtlPdgInstrBal    SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlPdgInstrBal,omitempty"`
	TtlRjctdInstrBal  SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlRjctdInstrBal,omitempty"`
	TtlPrtctInstrBal  SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TtlPrtctInstrBal,omitempty"`
	OptnDtls          []InstructedCorporateActionOption16 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnDtls,omitempty"`
}

type InstructedCorporateActionOption16 struct {
	OptnNb             Exact3NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnNb,omitempty"`
	OptnTp             CorporateActionOption36Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnTp"`
	InstdBal           BalanceFormat7Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstdBal"`
	DfltActn           DefaultProcessingOrStandingInstruction1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 DfltActn,omitempty"`
	OptnAccptdInstdBal SignedQuantityFormat9                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnAccptdInstdBal,omitempty"`
	OptnCancInstrBal   SignedQuantityFormat9                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnCancInstrBal,omitempty"`
	OptnPdgInstrBal    SignedQuantityFormat9                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnPdgInstrBal,omitempty"`
	OptnRjctdInstrBal  SignedQuantityFormat9                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnRjctdInstrBal,omitempty"`
	OptnPrtctInstrBal  SignedQuantityFormat9                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnPrtctInstrBal,omitempty"`
	EvtDdlns           CorporateActionEventDeadlines4                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 EvtDdlns"`
	OptnInstrDtls      []OptionInstructionDetails6                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OptnInstrDtls,omitempty"`
}

type InstructionProcessingStatus43Choice struct {
	Accptd             NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Accptd,omitempty"`
	Canc               CancelledStatus15Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Canc,omitempty"`
	AccptdForFrthrPrcg NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AccptdForFrthrPrcg,omitempty"`
	Rjctd              RejectedStatus39Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Rjctd,omitempty"`
	Pdg                NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Pdg,omitempty"`
	PdgCxl             PendingCancellationStatus9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PdgCxl,omitempty"`
	Cvrd               NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cvrd,omitempty"`
	Ucvrd              NoSpecifiedReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Ucvrd,omitempty"`
}

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be one of NORE
type NoReasonCode string

type NoSpecifiedReason1 struct {
	NoSpcfdRsn NoReasonCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NoSpcfdRsn"`
}

type NotificationIdentification6 struct {
	Id      RestrictedFINXMax16Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id"`
	CreDtTm DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CreDtTm,omitempty"`
}

type OptionInstructionDetails6 struct {
	InstrId      RestrictedFINMax15Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrId"`
	InstrSeqNb   Max3NumericText                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrSeqNb,omitempty"`
	PrtctInd     ProtectTransactionType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtctInd,omitempty"`
	InstrQty     FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrQty"`
	InstrDt      ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrDt"`
	PrtctDt      ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtctDt,omitempty"`
	CoverPrtctDt ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CoverPrtctDt,omitempty"`
	BidPric      PriceFormat57Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 BidPric,omitempty"`
	CondlQty     FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CondlQty,omitempty"`
	CstmrRef     RestrictedFINMax30Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 CstmrRef,omitempty"`
	InstrNrrtv   RestrictedFINXMax350Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrNrrtv,omitempty"`
	InstrSts     InstructionProcessingStatus43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrSts"`
}

type OriginalAndCurrentQuantities7 struct {
	ShrtLngPos ShortLong1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ShrtLngPos"`
	FaceAmt    float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 FaceAmt"`
	AmtsdVal   float64        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AmtsdVal"`
}

type OtherIdentification2 struct {
	Id  RestrictedFINXMax31Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Sfx,omitempty"`
	Tp  IdentificationSource4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Tp"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 LastPgInd"`
}

type PartyIdentification136Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AnyBIC,omitempty"`
	PrtryId GenericIdentification84 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtryId,omitempty"`
}

type PendingBalance6 struct {
	Bal    SignedQuantityFormat9               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Bal"`
	PdgTxs []SettlementTypeAndIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PdgTxs,omitempty"`
}

// May be one of ADEA, DQUA, DQCS, LATE, OTHR
type PendingCancellationReason5Code string

type PendingCancellationReason6Choice struct {
	Cd    PendingCancellationReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type PendingCancellationStatus9Choice struct {
	NotSpcfdRsn NoReasonCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NotSpcfdRsn,omitempty"`
	Rsn         []PendingCancellationStatusReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Rsn,omitempty"`
}

type PendingCancellationStatusReason9 struct {
	RsnCd       PendingCancellationReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RsnCd"`
	AddtlRsnInf RestrictedFINMax210Text          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AddtlRsnInf,omitempty"`
}

type PercentagePrice1 struct {
	PctgPricTp PriceRateType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PctgPricTp"`
	PricVal    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PricVal"`
}

type PriceFormat57Choice struct {
	PctgPric     PercentagePrice1     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PctgPric,omitempty"`
	AmtPric      AmountPrice5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AmtPric,omitempty"`
	NotSpcfdPric PriceValueType10Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NotSpcfdPric,omitempty"`
}

// May be one of DISC, PREM, PRCT, YIEL
type PriceRateType3Code string

// May be one of UKWN
type PriceValueType10Code string

type ProprietaryQuantity10 struct {
	ShrtLngPos ShortLong1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ShrtLngPos,omitempty"`
	Qty        float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Qty"`
	QtyTp      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 QtyTp"`
	Issr       Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Issr"`
	SchmeNm    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SchmeNm,omitempty"`
}

type ProprietaryQuantity9 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 QtyTp"`
	Issr    Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SchmeNm,omitempty"`
}

// May be one of PROT, COVP, COVR
type ProtectTransactionType2Code string

type Quantity21Choice struct {
	Qty      FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Qty,omitempty"`
	PrtryQty ProprietaryQuantity9                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtryQty,omitempty"`
}

type Quantity22Choice struct {
	QtyChc   Quantity23Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 QtyChc,omitempty"`
	PrtryQty ProprietaryQuantity10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PrtryQty,omitempty"`
}

type Quantity23Choice struct {
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities7 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OrgnlAndCurFaceAmt,omitempty"`
	SgndQty            SignedQuantityFormat9         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SgndQty,omitempty"`
}

type RejectedReason39Choice struct {
	Cd    RejectionReason57Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type RejectedStatus39Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NoSpcfdRsn,omitempty"`
	Rsn        []RejectedStatusReason37 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Rsn,omitempty"`
}

type RejectedStatusReason37 struct {
	RsnCd       RejectedReason39Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RsnCd"`
	AddtlRsnInf RestrictedFINMax210Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 AddtlRsnInf,omitempty"`
}

// May be one of ADEA, CERT, INVA, OPTY, ULNK, DSEC, LACK, LATE, NMTY, FULL, CANC, INTV, OPNM, OTHR, DQUA, REFT, SAFE, EVNM, DQCS, DQCC, DQAM, IRDQ, DQBV, DQBI, DCAN, DPRG, INIR, SHAR, BSTR, CTCT, DUPL, PROI, PROT, PRON, TRTI
type RejectionReason57Code string

type RestrictedFINActiveCurrencyAnd13DecimalAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern XX|TS
type RestrictedFINExact2Text string

// May be no more than 15 items long
type RestrictedFINMax15Text string

// May be no more than 210 items long
type RestrictedFINMax210Text string

// Must match the pattern ([^/]+/)+([^/]+)|([^/]*)
type RestrictedFINMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,140}
type RestrictedFINXMax140Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax16Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,210}
type RestrictedFINXMax210Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax30Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,31}
type RestrictedFINXMax31Text string

// Must match the pattern ([0-9a-zA-Z\-\?:\(\)\.,'\+ ]([0-9a-zA-Z\-\?:\(\)\.,'\+ ]*(/[0-9a-zA-Z\-\?:\(\)\.,'\+ ])?)*)
type RestrictedFINXMax34Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.\n\r,'\+ ]{1,350}
type RestrictedFINXMax350Text string

// Must match the pattern [0-9a-zA-Z/\-\?:\(\)\.,'\+ ]{1,35}
type RestrictedFINXMax35Text string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat32Choice struct {
	Id      SafekeepingPlaceTypeAndText9           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id,omitempty"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Ctry,omitempty"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TpAndId,omitempty"`
	Prtry   GenericIdentification85                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id"`
}

type SafekeepingPlaceTypeAndText9 struct {
	SfkpgPlcTp SafekeepingPlace2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SfkpgPlcTp"`
	Id         RestrictedFINXMax30Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Id,omitempty"`
}

type SecurityIdentification20 struct {
	ISIN   ISINOct2015Identifier    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ISIN,omitempty"`
	OthrId []OtherIdentification2   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 OthrId,omitempty"`
	Desc   RestrictedFINXMax140Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Desc,omitempty"`
}

type SettlementTypeAndIdentification26 struct {
	Pmt     DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Pmt"`
	TxId    RestrictedFINXMax16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 TxId"`
	SttlmDt DateAndDateTime2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 SttlmDt,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat8 struct {
	ShrtLngPos ShortLong1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ShrtLngPos"`
	QtyChc     Quantity21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 QtyChc"`
}

type SignedQuantityFormat9 struct {
	ShrtLngPos ShortLong1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Qty"`
}

type Statement75 struct {
	StmtTp        CorporateActionStatementType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StmtTp"`
	RptgTp        CorporateActionStatementReportingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RptgTp"`
	StmtId        RestrictedFINXMax16Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StmtId"`
	InstrAggtnPrd DatePeriod2                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 InstrAggtnPrd,omitempty"`
	RptNb         Max5NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 RptNb,omitempty"`
	StmtDtTm      DateAndDateTime2Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 StmtDtTm"`
	Frqcy         Frequency26Choice                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Frqcy"`
	UpdTp         UpdateType16Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 UpdTp"`
	ActvtyInd     bool                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 ActvtyInd"`
	NtfctnDdlnPrd DateOrDateTimePeriod1Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 NtfctnDdlnPrd,omitempty"`
}

// May be one of COMP, DELT
type StatementUpdateType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UpdateType16Choice struct {
	Cd    StatementUpdateType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Cd,omitempty"`
	Prtry GenericIdentification47  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.10 Prtry,omitempty"`
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
