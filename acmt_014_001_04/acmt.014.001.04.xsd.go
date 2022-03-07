// Code generated by download. DO NOT EDIT.

package iso20022_acmt_014_001_04

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AccountContract3 struct {
	TrgtGoLiveDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TrgtGoLiveDt,omitempty"`
	TrgtClsgDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TrgtClsgDt,omitempty"`
	GoLiveDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 GoLiveDt,omitempty"`
	ClsgDt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ClsgDt,omitempty"`
	UrgcyFlg     bool    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 UrgcyFlg,omitempty"`
	RmvlInd      bool    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 RmvlInd,omitempty"`
}

type AccountForAction1 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	Ccy ActiveCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Ccy"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Othr,omitempty"`
}

type AccountReport28 struct {
	Acct             CustomerAccount5                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Acct"`
	UndrlygMstrAgrmt ContractDocument1                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 UndrlygMstrAgrmt,omitempty"`
	CtrctDts         AccountContract3                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtrctDts,omitempty"`
	Mndt             []OperationMandate4                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Mndt,omitempty"`
	Grp              []Group4                                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Grp,omitempty"`
	RefAcct          CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 RefAcct,omitempty"`
	BalTrfAcct       AccountForAction1                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BalTrfAcct,omitempty"`
	TrfAcctSvcrId    BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TrfAcctSvcrId,omitempty"`
}

type AccountReportV04 struct {
	Refs        References5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Refs"`
	Fr          OrganisationIdentification29                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Fr,omitempty"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AcctSvcrId"`
	Org         Organisation33                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Org"`
	Rpt         []AccountReport28                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Rpt,omitempty"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 DgtlSgntr,omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

// May be one of ENAB, DISA, DELE, FORM
type AccountStatus3Code string

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
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Authorisation2 struct {
	MaxAmtByTx          FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MaxAmtByTx,omitempty"`
	MaxAmtByPrd         []MaximumAmountByPeriod1      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MaxAmtByPrd,omitempty"`
	MaxAmtByBlkSubmissn FixedAmountOrUnlimited1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MaxAmtByBlkSubmissn,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Domn,omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SubFmlyCd"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstlAdr,omitempty"`
}

type CashAccount40 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id,omitempty"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type Channel2Choice struct {
	Cd    CommunicationMethod3Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MmbId"`
}

type CodeOrProprietary1Choice struct {
	Cd    Max4Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type CommunicationFormat1Choice struct {
	Cd    ExternalCommunicationFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type CommunicationMethod2Choice struct {
	Cd    CommunicationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

// May be one of EMAL, FAXI, FILE, ONLI, POST
type CommunicationMethod2Code string

// May be one of EMAL, FAXI, POST, PHON, FILE, ONLI
type CommunicationMethod3Code string

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PrefrdMtd,omitempty"`
}

type ContractDocument1 struct {
	Ref      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Ref"`
	SgnOffDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SgnOffDt,omitempty"`
	Vrsn     Max6Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Vrsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CustomerAccount5 struct {
	Id               []AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	Nm               Max70Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	Sts              AccountStatus3Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Sts,omitempty"`
	Tp               CashAccountType2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Tp,omitempty"`
	Ccy              ActiveCurrencyCode             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Ccy"`
	MnthlyPmtVal     float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MnthlyPmtVal,omitempty"`
	MnthlyRcvdVal    float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MnthlyRcvdVal,omitempty"`
	MnthlyTxNb       Max5NumericText                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MnthlyTxNb,omitempty"`
	AvrgBal          float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AvrgBal,omitempty"`
	AcctPurp         Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AcctPurp,omitempty"`
	FlrNtfctnAmt     float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 FlrNtfctnAmt,omitempty"`
	ClngNtfctnAmt    float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ClngNtfctnAmt,omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndForm1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 StmtFrqcyAndFrmt,omitempty"`
	ClsgDt           ISODate                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ClsgDt,omitempty"`
	Rstrctn          []Restriction1                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Rstrctn,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtryOfBirth"`
}

type Document struct {
	AcctRpt AccountReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AcctRpt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalBankTransactionDomain1Code string

// May be no more than 4 items long
type ExternalBankTransactionFamily1Code string

// May be no more than 4 items long
type ExternalBankTransactionSubFamily1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalCommunicationFormat1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// May be no more than 4 items long
type ExternalPersonIdentification1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Othr,omitempty"`
}

type FixedAmountOrUnlimited1Choice struct {
	Amt    ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Amt,omitempty"`
	NotLtd Unlimited9Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 NotLtd,omitempty"`
}

// May be one of YEAR, DAIL, MNTH, QURT, MIAN, TEND, MOVE, WEEK, INDA
type Frequency7Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr,omitempty"`
}

type Group4 struct {
	GrpId Max4AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 GrpId"`
	Pty   []PartyAndCertificate4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Pty"`
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

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 128 items long
type Max128Text string

// May be no more than 140 items long
type Max140Text string

// Must match the pattern [\+]{0,1}[0-9]{1,15}
type Max15PlusSignedNumericText string

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

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 4 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 6 items long
type Max6Text string

// May be no more than 70 items long
type Max70Text string

type MaximumAmountByPeriod1 struct {
	MaxAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MaxAmt"`
	NbOfDays Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 NbOfDays"`
}

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CreDtTm"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OperationMandate4 struct {
	Id           Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
	AplblChanl   []Channel2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AplblChanl"`
	ReqrdSgntrNb Max15PlusSignedNumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ReqrdSgntrNb"`
	SgntrOrdrInd bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SgntrOrdrInd"`
	MndtHldr     []PartyAndAuthorisation4        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MndtHldr,omitempty"`
	BkOpr        []BankTransactionCodeStructure4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BkOpr"`
	StartDt      ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 StartDt,omitempty"`
	EndDt        ISODate                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 EndDt,omitempty"`
}

type Organisation33 struct {
	FullLglNm    Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 FullLglNm"`
	TradgNm      Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TradgNm,omitempty"`
	CtryOfOpr    CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtryOfOpr"`
	RegnDt       ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 RegnDt,omitempty"`
	OprlAdr      PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 OprlAdr,omitempty"`
	BizAdr       PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BizAdr,omitempty"`
	LglAdr       PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 LglAdr"`
	BllgAdr      PostalAddress24              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BllgAdr,omitempty"`
	OrgId        OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 OrgId"`
	RprtvOffcr   []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 RprtvOffcr,omitempty"`
	TrsrMgr      PartyIdentification137       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TrsrMgr,omitempty"`
	MainMndtHldr []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MainMndtHldr,omitempty"`
	Sndr         []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Sndr,omitempty"`
	LglRprtv     []PartyIdentification137     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 LglRprtv,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PrvtId,omitempty"`
}

type PartyAndAuthorisation4 struct {
	PtyOrGrp  PartyOrGroup2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PtyOrGrp"`
	SgntrOrdr Max15PlusSignedNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SgntrOrdr,omitempty"`
	Authstn   Authorisation2             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Authstn"`
}

type PartyAndCertificate4 struct {
	Pty  PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Pty"`
	Cert Max10KBinary           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cert,omitempty"`
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Pty"`
	Sgntr SkipPayload            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Sgntr"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtctDtls,omitempty"`
}

type PartyIdentification137 struct {
	Nm        Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Nm,omitempty"`
	PstlAdr   PostalAddress24        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstlAdr,omitempty"`
	Id        PersonIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id,omitempty"`
	CtryOfRes CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtryOfRes,omitempty"`
	CtctDtls  Contact4               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtctDtls,omitempty"`
}

type PartyOrGroup2Choice struct {
	GrpId Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 GrpId,omitempty"`
	Pty   PartyAndCertificate4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Pty,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Issr,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Prtry,omitempty"`
}

type References5 struct {
	ReqTp       UseCases1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ReqTp"`
	MsgId       MessageIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 MsgId"`
	PrcId       MessageIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PrcId"`
	AckdMsgId   []MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AckdMsgId,omitempty"`
	Sts         Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Sts,omitempty"`
	AttchdDocNm []Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 AttchdDocNm,omitempty"`
}

type Restriction1 struct {
	RstrctnTp CodeOrProprietary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 RstrctnTp"`
	VldFr     ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 VldFr"`
	VldUntil  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 VldUntil,omitempty"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type StatementFrequencyAndForm1 struct {
	Frqcy    Frequency7Code             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Frqcy"`
	ComMtd   CommunicationMethod2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 ComMtd"`
	DlvryAdr Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 DlvryAdr"`
	Frmt     CommunicationFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Frmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// Must match the pattern UNLIMITED
type Unlimited9Text string

// May be one of OPEN, MNTN, CLSG, VIEW
type UseCases1Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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