// Code generated by gen. DO NOT EDIT.

package iso20022_caam_006_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, SNDM, RPTC
type ATMCommand4Code string

type ATMCommand7 struct {
	Tp        ATMCommand4Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Tp"`
	Urgcy     TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Urgcy"`
	DtTm      ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 DtTm,omitempty"`
	CmdId     ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CmdId,omitempty"`
	CmdParams ATMCommandParameters1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CmdParams,omitempty"`
}

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Prcr,omitempty"`
}

type ATMCommandParameters1Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATMReqrdGblSts,omitempty"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 XpctdMsgFctn,omitempty"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ReqrdCfgtnParam,omitempty"`
}

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Vrsn"`
}

// May be one of ALRM, BRCD, CAMR, CRDD, CRDR, CSHD, CSHI, CSHR, CHCK, CDIS, DPST, DPRN, DOOR, INPM, JRNL, JPRN, SNSR, PSBK, PINR, RPRN, SCAN, RWDR
type ATMDevice2Code string

type ATMDiagnosticResponse2 struct {
	Envt ATMEnvironment6 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Envt"`
	Cmd  []ATMCommand7   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Cmd,omitempty"`
}

type ATMDiagnosticResponseV02 struct {
	Hdr                 Header31                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Hdr"`
	PrtctdATMDgnstcRspn ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 PrtctdATMDgnstcRspn,omitempty"`
	ATMDgnstcRspn       ATMDiagnosticResponse2   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATMDgnstcRspn,omitempty"`
	SctyTrlr            ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SctyTrlr,omitempty"`
}

type ATMEnvironment6 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Acqrr,omitempty"`
	ATMMgrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATMMgrId,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine5 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATM"`
}

type ATMEquipment1 struct {
	Manfctr    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Manfctr,omitempty"`
	Mdl        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Mdl,omitempty"`
	SrlNb      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SrlNb,omitempty"`
	ApplPrvdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ApplPrvdr,omitempty"`
	ApplNm     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ApplNm,omitempty"`
	ApplVrsn   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ApplVrsn,omitempty"`
	ApprvlNb   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ApprvlNb,omitempty"`
	CfgtnParam []ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CfgtnParam,omitempty"`
}

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 HstSvcCd,omitempty"`
}

// May be one of INSV, OUTS
type ATMStatus1Code string

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Brnch,omitempty"`
}

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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 MAC"`
}

type AutomatedTellerMachine5 struct {
	Id          Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Id"`
	AddtlId     Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AddtlId,omitempty"`
	SeqNb       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SeqNb,omitempty"`
	BaseCcy     ActiveCurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 BaseCcy"`
	Lctn        PostalAddress17             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Lctn,omitempty"`
	LctnCtgy    TransactionEnvironment2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 LctnCtgy,omitempty"`
	Eqpmnt      ATMEquipment1               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Eqpmnt,omitempty"`
	OutOfSvcDvc []ATMDevice2Code            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 OutOfSvcDvc,omitempty"`
	MsgPrtcn    MessageProtection1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 MsgPrtcn,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMDgnstcRspn ATMDiagnosticResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ATMDgnstcRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptdCntt,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 GeogcCordints,omitempty"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 UTMCordints,omitempty"`
}

type Header31 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 35 items long
type Max35Text string

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

// May be no more than 6 items long
type Max6Text string

// May be no more than 70 items long
type Max70Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// May be one of EVLP, MACB, MACM, UNPR
type MessageProtection1Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KEK,omitempty"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyIdr,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 AttrVal"`
}

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 TracDtTmOut"`
}

// May be one of PRIV, PUBL
type TransactionEnvironment2Code string

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.02 UTMNrthwrd"`
}

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
