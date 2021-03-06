// Code generated by download. DO NOT EDIT.

package iso20022_camt_079_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type AcknowledgedAcceptedStatus21Choice struct {
	NoSpcfdRsn NoReasonCode             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NoSpcfdRsn,omitempty"`
	Rsn        []AcknowledgementReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
}

type AcknowledgementReason12Choice struct {
	Cd    AcknowledgementReason5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

// May be one of ADEA, SMPG, OTHR, CDCY, CDRG, CDRE, NSTP, RQWV, LATE
type AcknowledgementReason5Code string

type AcknowledgementReason9 struct {
	Cd          AcknowledgementReason12Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd"`
	AddtlRsnInf Max210Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AmtWthtCcy,omitempty"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AmtWthCcy,omitempty"`
}

type AmountAndDirection5 struct {
	Amt    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Amt"`
	CdtDbt CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CdtDbt,omitempty"`
}

type AmountAndQuantityBreakdown1 struct {
	LotNb       GenericIdentification37            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LotNb,omitempty"`
	LotAmt      AmountAndDirection5                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LotAmt,omitempty"`
	LotQty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LotQty,omitempty"`
	CshSubBalTp GenericIdentification30            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshSubBalTp,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PstlAdr,omitempty"`
}

type CancellationReason19Choice struct {
	Cd    CancelledStatusReason13Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type CancellationReason9 struct {
	Cd          CancellationReason19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd"`
	AddtlRsnInf Max210Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type CancellationStatus14Choice struct {
	NoSpcfdRsn NoReasonCode          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NoSpcfdRsn,omitempty"`
	Rsn        []CancellationReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
}

// May be one of CANI, CANS, CSUB, CXLR, CANT, CANZ, CORP, SCEX, OTHR, CTHP
type CancelledStatusReason13Code string

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type CashBalanceType3Choice struct {
	Cd    ExternalBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type CashSubBalanceTypeAndQuantityBreakdown3 struct {
	Tp        CashBalanceType3Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Tp"`
	QtyBrkdwn []AmountAndQuantityBreakdown1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 QtyBrkdwn,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MmbId"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 DtTm,omitempty"`
}

type Document struct {
	IntraBalMvmntQryRspn IntraBalanceMovementQueryResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 IntraBalMvmntQryRspn"`
}

type DocumentIdentification51 struct {
	Id       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	CreDtTm  DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CreDtTm,omitempty"`
	CpyDplct CopyDuplicate1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CpyDplct,omitempty"`
	MsgOrgtr PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MsgRcpt,omitempty"`
}

type DocumentNumber5Choice struct {
	ShrtNb  Exact3NumericText                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 ShrtNb,omitempty"`
	LngNb   ISO20022MessageIdentificationText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LngNb,omitempty"`
	PrtryNb GenericIdentification36           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrtryNb,omitempty"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Desc,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalBalanceType1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalSystemErrorHandling1Code string

// May be one of AWMO, BYIY, CLAT, ADEA, CANR, CAIS, OBJT, AWSH, PHSE, STCD, DOCY, MLAT, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type FailingReason3Code string

type FailingReason7 struct {
	Cd          FailingReason7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd"`
	AddtlRsnInf Max210Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type FailingReason7Choice struct {
	Cd    FailingReason3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type FailingStatus9Choice struct {
	NoSpcfdRsn NoReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NoSpcfdRsn,omitempty"`
	Rsn        []FailingReason7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Othr,omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Unit,omitempty"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 FaceAmt,omitempty"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AmtsdVal,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SchmeNm,omitempty"`
}

type GenericIdentification37 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type IntraBalanceMovement5 struct {
	CshAcct           CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcct,omitempty"`
	CshAcctOwnr       SystemPartyIdentification8                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcctOwnr,omitempty"`
	CshAcctSvcr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcctSvcr,omitempty"`
	StsAndRsn         IntraBalanceStatusAndReason2                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 StsAndRsn,omitempty"`
	AcctOwnrTxId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AcctOwnrTxId"`
	AcctSvcrTxId      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MktInfrstrctrTxId,omitempty"`
	PrcrTxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrcrTxId,omitempty"`
	PoolId            Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PoolId,omitempty"`
	CorpActnEvtId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CorpActnEvtId,omitempty"`
	MvmntDtls         IntraBalanceMovement6                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MvmntDtls,omitempty"`
}

type IntraBalanceMovement6 struct {
	BalFr              CashSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BalFr"`
	BalTo              CashSubBalanceTypeAndQuantityBreakdown3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BalTo"`
	SttlmAmt           Amount2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SttlmAmt"`
	SttldAmt           Amount2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SttldAmt,omitempty"`
	PrevslySttldAmt    Amount2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrevslySttldAmt,omitempty"`
	RmngSttlmAmt       Amount2Choice                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RmngSttlmAmt,omitempty"`
	IntnddSttlmDt      DateAndDateTime2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 IntnddSttlmDt"`
	FctvSttlmDt        DateAndDateTime2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 FctvSttlmDt,omitempty"`
	StsDt              ISODateTime                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 StsDt,omitempty"`
	CshSubBalId        GenericIdentification37                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshSubBalId,omitempty"`
	Lnkgs              []Linkages57                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Lnkgs,omitempty"`
	Prty               PriorityNumeric4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prty,omitempty"`
	MsgOrgtr           SystemPartyIdentification8              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MsgOrgtr,omitempty"`
	CreDtTm            ISODateTime                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CreDtTm"`
	InstrPrcgAddtlDtls Max350Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 InstrPrcgAddtlDtls,omitempty"`
	SplmtryData        []SupplementaryData1                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SplmtryData,omitempty"`
}

type IntraBalanceMovementQueryResponseV01 struct {
	Id         DocumentIdentification51              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id,omitempty"`
	Pgntn      Pagination1                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Pgntn"`
	RptGnlDtls MovementReport1                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RptGnlDtls"`
	RptOrErr   IntraBalanceOrOperationalError7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RptOrErr,omitempty"`
}

type IntraBalanceMovements3 struct {
	CshAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcct,omitempty"`
	CshAcctOwnr SystemPartyIdentification8                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcctOwnr,omitempty"`
	CshAcctSvcr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CshAcctSvcr,omitempty"`
	StsAndRsn   IntraBalanceStatusAndReason2                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 StsAndRsn,omitempty"`
	Mvmnt       []IntraBalanceMovement5                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Mvmnt"`
}

type IntraBalanceOrOperationalError7Choice struct {
	Mvmnts  []IntraBalanceMovements3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Mvmnts,omitempty"`
	OprlErr []ErrorHandling5         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 OprlErr,omitempty"`
}

type IntraBalanceStatusAndReason2 struct {
	PrcgSts  []ProcessingStatus67Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrcgSts,omitempty"`
	SttlmSts []SettlementStatus16Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SttlmSts,omitempty"`
	Sttld    ProprietaryReason4         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Sttld,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type Linkages57 struct {
	PrcgPos ProcessingPosition7Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrcgPos,omitempty"`
	MsgNb   DocumentNumber5Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MsgNb,omitempty"`
	Ref     References34Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Ref"`
	RefOwnr PartyIdentification127Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RefOwnr,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 210 items long
type Max210Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

type MovementReport1 struct {
	QryRef    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 QryRef,omitempty"`
	RptId     Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RptId,omitempty"`
	QryTp     MovementResponseType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 QryTp"`
	ActvtyInd bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 ActvtyInd"`
}

// May be one of FULL, STTS
type MovementResponseType1Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Adr,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LastPgInd"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NmAndAdr,omitempty"`
}

type PartyIdentification127Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AnyBIC,omitempty"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrtryId,omitempty"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 LEI,omitempty"`
}

// May be one of AWMO, ADEA, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, PART, NMAS, NOFX, CMON, YCOL, COLL, DEPO, FLIM, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, IAAD, OTHR, PHCK, BENO, BOTH, CLHT, DENO, DISA, DKNY, FROZ, LAAW, LATE, LIQU, PRCY, REGT, SETS, CERT, PRSY, INBC
type PendingReason10Code string

type PendingReason14 struct {
	Cd          PendingReason26Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd"`
	AddtlRsnInf Max210Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type PendingReason26Choice struct {
	Cd    PendingReason10Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type PendingStatus36Choice struct {
	NoSpcfdRsn NoReasonCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NoSpcfdRsn,omitempty"`
	Rsn        []PendingReason14 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Ctry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AdrLine,omitempty"`
}

type PriorityNumeric4Choice struct {
	Nmrc  Exact4NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Nmrc,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

type ProcessingPosition7Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type ProcessingStatus67Choice struct {
	Rjctd      RejectionOrRepairStatus38Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rjctd,omitempty"`
	Rpr        RejectionOrRepairStatus38Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rpr,omitempty"`
	Canc       CancellationStatus14Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Canc,omitempty"`
	AckdAccptd AcknowledgedAcceptedStatus21Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AckdAccptd,omitempty"`
	Prtry      ProprietaryStatusAndReason6        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PrtryRsn,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type References34Choice struct {
	SctiesSttlmTxId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 SctiesSttlmTxId,omitempty"`
	IntraPosMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 IntraPosMvmntId,omitempty"`
	IntraBalMvmntId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 IntraBalMvmntId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 MktInfrstrctrTxId,omitempty"`
	PoolId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PoolId,omitempty"`
	OthrTxId          Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 OthrTxId,omitempty"`
}

type RejectionAndRepairReason32Choice struct {
	Cd    RejectionReason33Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type RejectionOrRepairReason32 struct {
	Cd          []RejectionAndRepairReason32Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Cd,omitempty"`
	AddtlRsnInf Max210Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 AddtlRsnInf,omitempty"`
}

type RejectionOrRepairStatus38Choice struct {
	NoSpcfdRsn NoReasonCode                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 NoSpcfdRsn,omitempty"`
	Rsn        []RejectionOrRepairReason32 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Rsn,omitempty"`
}

// May be one of CASH, ADEA, DMON, NCRR, LATE, INVL, INVB, INVN, VALR, MONY, CAEV, DDAT, REFE, OTHR, DQUA, DSEC, MINO, MUNO
type RejectionReason33Code string

type SettlementStatus16Choice struct {
	Pdg   PendingStatus36Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Pdg,omitempty"`
	Flng  FailingStatus9Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Flng,omitempty"`
	Prtry ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Prtry,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemPartyIdentification8 struct {
	Id           PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 Id"`
	RspnsblPtyId PartyIdentification136 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.079.001.01 RspnsblPtyId,omitempty"`
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
