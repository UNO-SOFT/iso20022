// Code generated by download. DO NOT EDIT.

package iso20022_pacs_010_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

// AccountIdentification4Choice
//
// Specifies the unique identification of an account as assigned by the account servicer.
type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Othr,omitempty"`
}

// AccountSchemeName1Choice
//
// Sets of elements to identify a name of the identification scheme.
type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// ActiveCurrencyAndAmount
//
// A number of monetary units specified in an active currency where the unit of currency is explicit and compliant with ISO 4217.
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

// AddressType3Choice
//
// Choice of formats for the type of address.
type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

// BranchAndFinancialInstitutionIdentification6
//
// Unique and unambiguous identification of a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 BrnchId,omitempty"`
}

// BranchData3
//
// Information that locates and identifies a specific branch of a financial institution.
type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PstlAdr,omitempty"`
}

// CashAccount40
//
// Provides the details to identify an account.
type CashAccount40 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id,omitempty"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prxy,omitempty"`
}

// CashAccountType2Choice
//
// Nature or use of the account.
type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// CategoryPurpose1Choice
//
// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

// ClearingSystemIdentification2Choice
//
// Choice of a clearing system identifier.
type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// ClearingSystemMemberIdentification2
//
// Unique identification, as assigned by a clearing system, to unambiguously identify a member of the clearing system.
type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// CreditTransferTransaction53
//
// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction53 struct {
	CdtId             Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtId"`
	BtchBookg         bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 BtchBookg,omitempty"`
	PmtTpInf          PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PmtTpInf,omitempty"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrBkSttlmDt,omitempty"`
	SttlmTmIndctn     SettlementDateTimeIndication1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SttlmTmIndctn,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstdAgt,omitempty"`
	IntrmyAgt1        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrmyAgt3Acct,omitempty"`
	CdtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtrAgt,omitempty"`
	CdtrAgtAcct       CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtrAgtAcct,omitempty"`
	Cdtr              BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cdtr"`
	CdtrAcct          CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtrAcct,omitempty"`
	UltmtCdtr         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []InstructionForCreditorAgent3               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstrForCdtrAgt,omitempty"`
	DrctDbtTxInf      []DirectDebitTransactionInformation27        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DrctDbtTxInf"`
	SplmtryData       []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SplmtryData,omitempty"`
}

// DirectDebitTransactionInformation27
//
// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation27 struct {
	PmtId           PaymentIdentification13                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PmtId"`
	PmtTpInf        PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PmtTpInf,omitempty"`
	IntrBkSttlmAmt  ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrBkSttlmAmt"`
	IntrBkSttlmDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 IntrBkSttlmDt,omitempty"`
	SttlmPrty       Priority3Code                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SttlmPrty,omitempty"`
	SttlmTmIndctn   SettlementDateTimeIndication1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SttlmTmIndctn,omitempty"`
	SttlmTmReq      SettlementTimeRequest2                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SttlmTmReq,omitempty"`
	UltmtDbtr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 UltmtDbtr,omitempty"`
	Dbtr            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Dbtr"`
	DbtrAcct        CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DbtrAcct,omitempty"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DbtrAgt,omitempty"`
	DbtrAgtAcct     CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DbtrAgtAcct,omitempty"`
	InstrForDbtrAgt Max210Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstrForDbtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Purp,omitempty"`
	RmtInf          RemittanceInformation2                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 RmtInf,omitempty"`
}

type Document struct {
	FIDrctDbt FinancialInstitutionDirectDebitV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 FIDrctDbt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 4 items long
type ExternalCategoryPurpose1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalCreditorAgentInstruction1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 35 items long
type ExternalLocalInstrument1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalPurpose1Code string

// May be no more than 4 items long
type ExternalServiceLevel1Code string

// FinancialIdentificationSchemeName1Choice
//
// Sets of elements to identify a name of the organisation identification scheme.
type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// FinancialInstitutionDirectDebitV05
//
// Scope:
// The FinancialInstitutionDirectDebit message is sent by an exchange or clearing house, or a financial institution, directly or through another agent, to the DebtorAgent. It is used to instruct the DebtorAgent to move funds from one or more debtor(s) account(s) to one or more creditor(s), where both debtor and creditor are financial institutions.
//
// Usage:
// The FinancialInstitutionDirectDebit message is exchanged between agents and can contain one or more financial institution direct debit instruction(s) for one or more creditor(s). The FinancialInstitutionDirectDebit message can be used in domestic and cross-border scenarios.
type FinancialInstitutionDirectDebitV05 struct {
	GrpHdr      GroupHeader92                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 GrpHdr"`
	CdtInstr    []CreditTransferTransaction53 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtInstr"`
	SplmtryData []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SplmtryData,omitempty"`
}

// FinancialInstitutionIdentification18
//
// Specifies the details to identify a financial institution.
type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Othr,omitempty"`
}

// GenericAccountIdentification1
//
// Information related to a generic account identification.
type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Issr,omitempty"`
}

// GenericFinancialIdentification1
//
// Information related to an identification of a financial institution.
type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Issr,omitempty"`
}

// GenericIdentification30
//
// Information related to an identification, for example, party identification or account identification.
type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SchmeNm,omitempty"`
}

// GroupHeader92
//
// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader92 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CreDtTm"`
	NbOfTxs  Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 NbOfTxs"`
	CtrlSum  float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CtrlSum,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstdAgt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// ISODate
//
// A particular point in the progression of time in a calendar year expressed in the YYYY-MM-DD format. This representation is defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// ISODateTime
//
// A particular point in the progression of time defined by a mandatory date and a mandatory time component, expressed in either UTC time format (YYYY-MM-DDThh:mm:ss.sssZ), local time with UTC offset format (YYYY-MM-DDThh:mm:ss.sss+/-hh:mm), or local time format (YYYY-MM-DDThh:mm:ss.sss). These representations are defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
// Note on the time format:
// 1) beginning / end of calendar day
// 00:00:00 = the beginning of a calendar day
// 24:00:00 = the end of a calendar day
// 2) fractions of second in time format
// Decimal fractions of seconds may be included. In this case, the involved parties shall agree on the maximum number of digits that are allowed.
type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// ISOTime
//
// A particular point in the progression of time in a calendar day expressed in either UTC time format (hh:mm:ss.sssZ), local time with UTC offset format (hh:mm:ss.sss+/-hh:mm), or local time format (hh:mm:ss.sss). These representations are defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
// Note on the time format:
// 1) beginning / end of calendar day
// 00:00:00 = the beginning of a calendar day
// 24:00:00 = the end of a calendar day
// 2) fractions of second in time format
// Decimal fractions of seconds may be included. In this case, the involved parties shall agree on the maximum number of digits that are allowed.
type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// InstructionForCreditorAgent3
//
// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor's agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the creditor's agent.
type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	InstrInf Max140Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstrInf,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// LocalInstrument2Choice
//
// Set of elements that further identifies the type of local instruments being requested by the initiating party.
type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

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

// May be no more than 70 items long
type Max70Text string

// PaymentIdentification13
//
// Provides further means of referencing a payment transaction.
type PaymentIdentification13 struct {
	InstrId    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 EndToEndId"`
	TxId       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 UETR,omitempty"`
	ClrSysRef  Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 ClrSysRef,omitempty"`
}

// PaymentTypeInformation28
//
// Provides further details of the type of payment.
type PaymentTypeInformation28 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CtgyPurp,omitempty"`
}

// PostalAddress24
//
// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 AdrLine,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

// May be one of URGT, HIGH, NORM
type Priority3Code string

// ProxyAccountIdentification1
//
// Information related to a proxy  identification of the account.
type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Id"`
}

// ProxyAccountType1Choice
//
// Specifies the scheme used for the identification of an account alias.
type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// Purpose2Choice
//
// Specifies the underlying reason for the payment transaction.
// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// RemittanceInformation2
//
// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle.
type RemittanceInformation2 struct {
	Ustrd []Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Ustrd,omitempty"`
}

// ServiceLevel8Choice
//
// Specifies the service level of the transaction.
type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Cd,omitempty"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Prtry,omitempty"`
}

// SettlementDateTimeIndication1
//
// Information on the occurred settlement time(s) of the payment transaction.
type SettlementDateTimeIndication1 struct {
	DbtDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 DbtDtTm,omitempty"`
	CdtDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CdtDtTm,omitempty"`
}

// SettlementTimeRequest2
//
// Provides information on the requested settlement time(s) of the payment instruction.
type SettlementTimeRequest2 struct {
	CLSTm  ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 CLSTm,omitempty"`
	TillTm ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 TillTm,omitempty"`
	FrTm   ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 FrTm,omitempty"`
	RjctTm ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 RjctTm,omitempty"`
}

// SupplementaryData1
//
// Additional information that can not be captured in the structured fields and/or any other specific block.
type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.05 Envlp"`
}

// SupplementaryDataEnvelope1
//
// Technical component that contains the validated supplementary data information. This technical envelope allows to segregate the supplementary data information from any other information.
type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

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
