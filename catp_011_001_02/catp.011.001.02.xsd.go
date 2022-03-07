// Code generated by download. DO NOT EDIT.

package iso20022_catp_011_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, SNDM, RPTC
type ATMCommand4Code string

type ATMCommand7 struct {
	Tp        ATMCommand4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Tp"`
	Urgcy     TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Urgcy"`
	DtTm      ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 DtTm,omitempty"`
	CmdId     ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CmdId,omitempty"`
	CmdParams ATMCommandParameters1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CmdParams,omitempty"`
}

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Prcr,omitempty"`
}

type ATMCommandParameters1Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMReqrdGblSts,omitempty"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 XpctdMsgFctn,omitempty"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ReqrdCfgtnParam,omitempty"`
}

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Vrsn"`
}

type ATMContext17 struct {
	SsnRef Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SsnRef,omitempty"`
	Svc    ATMService21 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Svc"`
}

type ATMCustomer5 struct {
	Prfl         ATMCustomerProfile2              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Prfl,omitempty"`
	AuthntcnRslt []TransactionVerificationResult5 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AuthntcnRslt,omitempty"`
}

type ATMCustomerProfile2 struct {
	PrflRef Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PrflRef,omitempty"`
	CstmrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CstmrId,omitempty"`
}

// May be one of CDIS, DPRN, JRNL, JPRN, RPRN, RWDR
type ATMDevice1Code string

type ATMEnvironment12 struct {
	Acqrr          Acquirer7                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Acqrr,omitempty"`
	ATMMgr         Acquirer8                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMMgr,omitempty"`
	HstgNtty       TerminalHosting1         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 HstgNtty,omitempty"`
	ATM            AutomatedTellerMachine2  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATM"`
	Cstmr          ATMCustomer5             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Cstmr"`
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData19          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PlainCardData,omitempty"`
}

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 HstSvcCd,omitempty"`
}

type ATMPINManagementResponse2 struct {
	Envt  ATMEnvironment12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Envt"`
	Cntxt ATMContext17     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Cntxt"`
	Tx    ATMTransaction22 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Tx"`
}

type ATMPINManagementResponseV02 struct {
	Hdr                  Header31                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Hdr"`
	PrtctdATMPINMgmtRspn ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PrtctdATMPINMgmtRspn,omitempty"`
	ATMPINMgmtRspn       ATMPINManagementResponse2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMPINMgmtRspn,omitempty"`
	SctyTrlr             ContentInformationType15  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SctyTrlr,omitempty"`
}

type ATMService21 struct {
	SvcRef     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SvcRef,omitempty"`
	ATMSvcCd   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMSvcCd,omitempty"`
	HstSvcCd   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 HstSvcCd,omitempty"`
	SvcTp      ATMServiceType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SvcTp"`
	SvcVarntId []Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SvcVarntId,omitempty"`
}

// May be one of PINC, PINR, PINU
type ATMServiceType5Code string

// May be one of INSV, OUTS
type ATMStatus1Code string

type ATMTransaction22 struct {
	TxId        TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TxId"`
	RcncltnId   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RcncltnId,omitempty"`
	CmpltnReqrd bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CmpltnReqrd,omitempty"`
	TxRspn      ResponseType7          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TxRspn"`
	Actn        []Action7              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Actn,omitempty"`
	ICCRltdData Max10000Binary         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ICCRltdData,omitempty"`
	Cmd         []ATMCommand7          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Cmd,omitempty"`
}

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Brnch,omitempty"`
}

type Acquirer8 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Id"`
	ApplVrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ApplVrsn,omitempty"`
}

type Action7 struct {
	ActnTp     ActionType6Code       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ActnTp"`
	MsgToPres  ActionMessage4        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MsgToPres,omitempty"`
	ReqToPrfrm MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ReqToPrfrm,omitempty"`
}

type ActionMessage4 struct {
	Frmt         OutputFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Frmt,omitempty"`
	Msg          Max20000Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Msg,omitempty"`
	Ref          Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Ref,omitempty"`
	Dvc          ATMDevice1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Dvc,omitempty"`
	MsgCnttSgntr Max35Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MsgCnttSgntr,omitempty"`
}

// May be one of DCCQ, FEES, HAMT, LAMT, BUSY, CPTR, DISP, CPNS, RQST, PINL, PINR, TRCK
type ActionType6Code string

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of TOKA, BIOM, MOBL, OTHR, FPIN, NPIN, PSWD, SCRT, SCNL
type AuthenticationMethod7Code string

type AutomatedTellerMachine2 struct {
	Id      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Id"`
	AddtlId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AddtlId,omitempty"`
	SeqNb   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SeqNb,omitempty"`
	BaseCcy ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 BaseCcy,omitempty"`
	Lctn    PostalAddress17    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Lctn,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMPINMgmtRspn ATMPINManagementResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ATMPINMgmtRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptdCntt,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 GeogcCordints,omitempty"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 UTMCordints,omitempty"`
}

type Header31 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 104 items long
type Max104Text string

// May be no more than 10 items long
type Max10Text string

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 20000 items long
type Max20000Text string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 35 items long
type Max35Text string

// May be no more than 37 items long
type Max37Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 500 items long
type Max500Text string

// May be no more than 6 items long
type Max6Text string

// May be no more than 70 items long
type Max70Text string

// May be no more than 76 items long
type Max76Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

// May be one of MREF, SREF, TEXT, HTML
type OutputFormat2Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PlainCardData19 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PAN,omitempty"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 XpryDt,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Trck3,omitempty"`
}

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KEK,omitempty"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyIdr,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AttrVal"`
}

// May be one of APPR, DECL, PART
type Response4Code string

type ResponseType7 struct {
	Rspn         Response4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Rspn"`
	RspnRsn      ResultDetail4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RspnRsn,omitempty"`
	AddtlRspnInf Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AddtlRspnInf,omitempty"`
}

// May be one of ACTF, ACQS, AMLV, AMTA, AUTH, BANK, CRDR, CRDF, ACTC, CTVG, DBER, FEES, TXNL, AMTD, NMBD, CRDX, FDCL, FMTR, TXNG, FNDI, ACPI, AMTI, ADDI, BRHI, CHDI, CRDI, CTFV, AMTO, PINV, TKKO, SGNI, TKID, TXNV, DATI, ISSP, ISSF, ISSO, ISST, ISSU, KEYS, LBLA, CRDL, MACR, MACK, ICCM, PINN, CRDA, LBLU, PINA, NPRA, OFFL, ONLP, NPRC, TXNM, OTHR, BALO, SEQO, PINC, PIND, PINS, PINX, PINE, QMAX, RECD, CRDT, SECV, SRVU, SFWE, SPCC, CRDS, SRCH, CNTC, FRDS, SYSP, SYSM, TRMI, ACTT, TTLV, TXNU, TXND, ORGF, UNBO, UNBP, UNBC, CMKY, CRDU, SVSU, VNDR, VNDF, AMTW, NMBW, CRDW, MEDI, SRVI
type ResultDetail4Code string

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TracDtTmOut"`
}

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 TxRef"`
}

type TransactionVerificationResult5 struct {
	Mtd         AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Mtd"`
	VrfctnNtty  AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 VrfctnNtty,omitempty"`
	Rslt        Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Rslt,omitempty"`
	AddtlRslt   Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AddtlRslt,omitempty"`
	AuthntcnTkn Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 AuthntcnTkn,omitempty"`
}

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 UTMNrthwrd"`
}

// May be one of FAIL, MISS, NOVF, PART, SUCC, ERRR
type Verification1Code string

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