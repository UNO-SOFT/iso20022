// Code generated by download. DO NOT EDIT.

package iso20022_auth_013_001_02

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

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type Agreement4 struct {
	AgrmtDtls  Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AgrmtDtls"`
	AgrmtId    Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AgrmtId,omitempty"`
	AgrmtDt    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AgrmtDt"`
	BaseCcy    ActiveCurrencyCode        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 BaseCcy"`
	AgrmtFrmwk AgreementFramework1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AgrmtFrmwk,omitempty"`
}

type AgreementFramework1Choice struct {
	AgrmtFrmwk AgreementFramework1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AgrmtFrmwk,omitempty"`
	PrtryId    GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PrtryId,omitempty"`
}

// May be one of FBAA, BBAA, DERV, ISDA, NONR
type AgreementFramework1Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of ACMB, CCPX, DCMB, FCMC, GCMB, SCMB
type CCPMemberType1Code string

type CashCollateral4 struct {
	AsstNb   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AsstNb,omitempty"`
	DpstAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 DpstAmt,omitempty"`
	DpstTp   DepositType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 DpstTp,omitempty"`
	BlckdAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 BlckdAmt,omitempty"`
	MtrtyDt  ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MtrtyDt,omitempty"`
	ValDt    ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValDt,omitempty"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XchgRate,omitempty"`
	CollVal  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollVal"`
	Hrcut    float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Hrcut,omitempty"`
}

type Collateral43 struct {
	AcctId    CollateralAccount3      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AcctId"`
	RptSummry Summary2                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptSummry"`
	CollValtn []CollateralValuation12 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollValtn,omitempty"`
}

type CollateralAccount3 struct {
	Id Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	Tp CollateralAccountIdentificationType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp,omitempty"`
	Nm Max70Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Nm,omitempty"`
}

type CollateralAccountIdentificationType3Choice struct {
	Tp    CollateralAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp,omitempty"`
	Prtry GenericIdentification36    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry,omitempty"`
}

// May be one of HOUS, CLIE, LIPR, MGIN, DFLT
type CollateralAccountType1Code string

type CollateralAmount1 struct {
	CollAmt        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollAmt"`
	RptdCcyAndAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptdCcyAndAmt"`
	MktValAmt      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MktValAmt"`
	AcrdIntrstAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AcrdIntrstAmt,omitempty"`
	FeesAndComssns ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 FeesAndComssns,omitempty"`
}

type CollateralAndExposureReportV04 struct {
	RptParams   ReportParameters6    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptParams"`
	Pgntn       Pagination1          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Pgntn,omitempty"`
	Oblgtn      Obligation6          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Oblgtn"`
	Agrmt       Agreement4           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Agrmt,omitempty"`
	CollRpt     []Collateral43       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollRpt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SplmtryData,omitempty"`
}

// May be one of APLD, EXCS
type CollateralAppliedExcess1Code string

// May be one of CDPA, CDPB
type CollateralDirection1Code string

type CollateralOwnership3 struct {
	Prtry  bool                         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry"`
	ClntNm PartyIdentification178Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ClntNm,omitempty"`
}

// May be one of COMO, CCCL, CEMC, CXCC, CFTD, CFTI, CTRC, CASH, LCRE, OTHR, SECU, CTCO, CCVR
type CollateralType8Code string

type CollateralValuation12 struct {
	CollId       Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollId,omitempty"`
	CollTp       CollateralType8Code            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollTp"`
	CollDrctn    CollateralDirection1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollDrctn,omitempty"`
	SttlmSts     SettlementStatus3Code          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SttlmSts"`
	ApldXcssInd  CollateralAppliedExcess1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ApldXcssInd,omitempty"`
	NbOfDaysAcrd float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 NbOfDaysAcrd,omitempty"`
	ValtnAmts    CollateralAmount1              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValtnAmts"`
	DayCntBsis   InterestComputationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 DayCntBsis,omitempty"`
	XchgRate     float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XchgRate,omitempty"`
	CcyHrcut     float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CcyHrcut,omitempty"`
	AdjstdRate   float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AdjstdRate,omitempty"`
	SctiesColl   SecuritiesCollateral9          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SctiesColl,omitempty"`
	CshColl      CashCollateral4                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CshColl,omitempty"`
	OthrColl     OtherCollateral8               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 OthrColl,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 DtTm,omitempty"`
}

type DateCode9Choice struct {
	Cd    DateType2Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry,omitempty"`
}

type DateFormat14Choice struct {
	Dt   ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Dt,omitempty"`
	DtCd DateCode9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 DtCd,omitempty"`
}

// May be one of OPEN
type DateType2Code string

// May be one of FITE, CALL
type DepositType1Code string

type Document struct {
	CollAndXpsrRpt CollateralAndExposureReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollAndXpsrRpt"`
}

// May be one of DAIL, INDA, ONDE
type EventFrequency6Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, SWPT, TBAS, TRBD, TRCP
type ExposureType5Code string

// May be one of CCIR, CRPR, EQUI, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, MGLD, OTCD, REPO, RVPO, SLOA, SBSC, SCRP, SLEB, SHSL, SCIR, SCIE, ESCL, SWPT, TBAS, ECRT, ECFR, EMLO, EMLI, EOIM, EOMI, TRCP, TRBD, BFWD, PAYM, CCPC, COMM, CRDS, CRTL, CRSP, EOMO
type ExposureType8Code string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Unit,omitempty"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 FaceAmt,omitempty"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AmtsdVal,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id,omitempty"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry,omitempty"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014, NARR
type InterestComputationMethod2Code string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress6 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Nm"`
	Adr PostalAddress2 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Adr"`
}

type Obligation6 struct {
	PtyA       PartyIdentification242       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PtyA"`
	SvcgPtyA   PartyIdentification178Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SvcgPtyA,omitempty"`
	PtyB       PartyIdentification242       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PtyB"`
	SvcgPtyB   PartyIdentification178Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SvcgPtyB,omitempty"`
	CollAcctId CollateralAccount3           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollAcctId,omitempty"`
	XpsrTp     ExposureType5Code            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XpsrTp,omitempty"`
	ValtnDt    DateAndDateTime2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValtnDt"`
}

type OtherCollateral8 struct {
	AsstNb       Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AsstNb,omitempty"`
	LttrOfCdtId  Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 LttrOfCdtId,omitempty"`
	LttrOfCdtAmt ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 LttrOfCdtAmt,omitempty"`
	GrntAmt      ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 GrntAmt,omitempty"`
	OthrTpOfColl OtherTypeOfCollateral2             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 OthrTpOfColl,omitempty"`
	CollOwnrsh   CollateralOwnership3               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollOwnrsh,omitempty"`
	IsseDt       DateFormat14Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 IsseDt,omitempty"`
	XpryDt       DateFormat14Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XpryDt,omitempty"`
	LtdCvrgInd   bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 LtdCvrgInd,omitempty"`
	Issr         PartyIdentification178Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Issr,omitempty"`
	BlckdQty     FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 BlckdQty,omitempty"`
	ValDt        ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValDt,omitempty"`
	XchgRate     float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XchgRate,omitempty"`
	MktVal       ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MktVal,omitempty"`
	Hrcut        float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Hrcut,omitempty"`
	CollVal      ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollVal"`
	SfkpgPlc     SafekeepingPlaceFormat29Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgPlc,omitempty"`
	SfkpgAcct    SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgAcct,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp"`
}

type OtherTypeOfCollateral2 struct {
	Desc Max140Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Desc"`
	Qty  FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Qty,omitempty"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 LastPgInd"`
}

type PartyIdentification178Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PrtryId,omitempty"`
	NmAndAdr NameAndAddress6         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 NmAndAdr,omitempty"`
}

type PartyIdentification242 struct {
	Id       PartyIdentification178Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	CCPMmbTp CCPMemberType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CCPMmbTp,omitempty"`
}

type PostalAddress2 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PstCdId"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Ctry"`
}

type Price7 struct {
	Tp  YieldedOrValueType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp"`
	Val PriceRateOrAmount3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Val"`
}

type PriceRateOrAmount3Choice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Rate,omitempty"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Amt,omitempty"`
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

type ReportParameters6 struct {
	RptId      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptId"`
	RptDtAndTm DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptDtAndTm"`
	Frqcy      EventFrequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Frqcy"`
	RptCcy     ActiveCurrencyCode     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RptCcy"`
	ClctnDt    ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ClctnDt,omitempty"`
}

type ReturnExcessCash1 struct {
	RtrXcssCshTp ReturnExcessCash1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RtrXcssCshTp"`
	CshCollCcy   ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CshCollCcy"`
}

type ReturnExcessCash1Choice struct {
	Cd    ReturnExcessCash1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry,omitempty"`
}

// May be one of RTND, RTDN, SSPD
type ReturnExcessCash1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE
type SafekeepingPlace3Code string

type SafekeepingPlaceFormat29Choice struct {
	Id      SafekeepingPlaceTypeAndText8           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id,omitempty"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Ctry,omitempty"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TpAndId,omitempty"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Prtry,omitempty"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
}

type SafekeepingPlaceTypeAndText8 struct {
	SfkpgPlcTp SafekeepingPlace3Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Nm,omitempty"`
}

type SecuritiesCollateral9 struct {
	AsstNb     Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AsstNb,omitempty"`
	SctyId     SecurityIdentification19           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SctyId"`
	MtrtyDt    DateAndDateTime2Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MtrtyDt,omitempty"`
	CollOwnrsh CollateralOwnership3               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollOwnrsh,omitempty"`
	LtdCvrgInd bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 LtdCvrgInd,omitempty"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Qty"`
	BlckdQty   FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 BlckdQty,omitempty"`
	Pric       Price7                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Pric,omitempty"`
	MktVal     ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MktVal,omitempty"`
	Hrcut      float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Hrcut,omitempty"`
	CollVal    ActiveCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollVal,omitempty"`
	ValDt      ISODate                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValDt,omitempty"`
	SfkpgAcct  SecuritiesAccount19                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgAcct,omitempty"`
	SfkpgPlc   SafekeepingPlaceFormat29Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SfkpgPlc"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Desc,omitempty"`
}

// May be one of ASTL, AAUT, ACCF, ARCF, MTCH, PSTL, RJCT, STLD, STCR, SPLT, NMAT
type SettlementStatus3Code string

// May be one of SHOR, LONG
type ShortLong1Code string

type Summary2 struct {
	XpsdAmtPtyA     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XpsdAmtPtyA,omitempty"`
	XpsdAmtPtyB     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XpsdAmtPtyB,omitempty"`
	XpsrTp          ExposureType8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 XpsrTp"`
	TtlValOfColl    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TtlValOfColl"`
	NetXcssDfcit    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 NetXcssDfcit,omitempty"`
	NetXcssDfcitInd ShortLong1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 NetXcssDfcitInd,omitempty"`
	ValtnDtTm       ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValtnDtTm"`
	ReqdSttlmDt     ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ReqdSttlmDt,omitempty"`
	SummryDtls      SummaryAmounts2         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 SummryDtls,omitempty"`
}

type SummaryAmounts2 struct {
	ThrshldAmt           ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ThrshldAmt,omitempty"`
	ThrshldTp            ThresholdType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ThrshldTp,omitempty"`
	PreHrcutCollVal      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PreHrcutCollVal,omitempty"`
	AdjstdXpsr           ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 AdjstdXpsr,omitempty"`
	CollReqrd            ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 CollReqrd,omitempty"`
	RtrXcssCshAndCollCcy []ReturnExcessCash1     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RtrXcssCshAndCollCcy,omitempty"`
	MinTrfAmt            ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 MinTrfAmt,omitempty"`
	RndgAmt              ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 RndgAmt,omitempty"`
	PrvsXpsrVal          ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PrvsXpsrVal,omitempty"`
	PrvsCollVal          ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PrvsCollVal,omitempty"`
	TtlPdgIncmgColl      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TtlPdgIncmgColl,omitempty"`
	TtlPdgOutgngColl     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TtlPdgOutgngColl,omitempty"`
	TtlAcrdIntrstAmt     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TtlAcrdIntrstAmt,omitempty"`
	TtlFees              ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 TtlFees,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of SECU, UNSE
type ThresholdType1Code string

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 Yldd,omitempty"`
	ValTp PriceValueType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.04 ValTp,omitempty"`
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
