// Code generated by download. DO NOT EDIT.

package iso20022_caam_009_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ATMCassette2 struct {
	PhysId    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PhysId,omitempty"`
	LogclId   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 LogclId"`
	SrlNb     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SrlNb,omitempty"`
	Tp        ATMCassetteType1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	SubTp     []ATMNoteType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SubTp,omitempty"`
	MdiaTp    ATMMediaType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MdiaTp,omitempty"`
	MdiaCntrs []ATMCassetteCounters3 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MdiaCntrs,omitempty"`
}

type ATMCassetteCounters3 struct {
	UnitVal  float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 UnitVal,omitempty"`
	Ccy      ActiveCurrencyCode     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ccy,omitempty"`
	MdiaCtgy ATMMediaType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MdiaCtgy,omitempty"`
	CurNb    float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CurNb"`
	CurAmt   float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CurAmt,omitempty"`
	FlowTtls []ATMCassetteCounters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 FlowTtls,omitempty"`
}

type ATMCassetteCounters4 struct {
	Tp        ATMCounterType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	AddedNb   float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AddedNb,omitempty"`
	RmvdNb    float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RmvdNb,omitempty"`
	DspnsdNb  float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 DspnsdNb,omitempty"`
	DpstdNb   float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 DpstdNb,omitempty"`
	RcycldNb  float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RcycldNb,omitempty"`
	RtrctdNb  float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RtrctdNb,omitempty"`
	RjctdNb   float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RjctdNb,omitempty"`
	PresntdNb float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PresntdNb,omitempty"`
}

// May be one of DPST, DISP, RCYC, RJCT, RPLT, RTRC
type ATMCassetteType1Code string

// May be one of ABAL, CCNT, RPTC
type ATMCommand5Code string

type ATMCommand8 struct {
	Tp          ATMCommand5Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	ReqrdDtTm   ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ReqrdDtTm,omitempty"`
	PrcdDtTm    ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PrcdDtTm"`
	CmdId       ATMCommandIdentification1           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CmdId,omitempty"`
	Rslt        TerminalManagementActionResult2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Rslt"`
	AddtlErrInf Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AddtlErrInf,omitempty"`
}

type ATMCommand9 struct {
	Tp    ATMCommand5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	CmdId ATMCommandIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CmdId,omitempty"`
}

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Prcr,omitempty"`
}

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Vrsn"`
}

// May be one of INQU, CTXN, CTOF, BDAY, PRTN, OPER
type ATMCounterType1Code string

// May be one of BDAY, INQU, CTOF, OPER
type ATMCounterType2Code string

type ATMEnvironment10 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Acqrr,omitempty"`
	ATMMgrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMMgrId,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine8 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATM"`
}

type ATMEquipment1 struct {
	Manfctr    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Manfctr,omitempty"`
	Mdl        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Mdl,omitempty"`
	SrlNb      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SrlNb,omitempty"`
	ApplPrvdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ApplPrvdr,omitempty"`
	ApplNm     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ApplNm,omitempty"`
	ApplVrsn   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ApplVrsn,omitempty"`
	ApprvlNb   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ApprvlNb,omitempty"`
	CfgtnParam []ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CfgtnParam,omitempty"`
}

// May be one of CARD, COIN, CMDT, CPNS, NOTE, STMP, UDTM
type ATMMediaType1Code string

// May be one of CNTR, FITN, FITU, SPCT, UNFT, UNRG
type ATMMediaType3Code string

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 HstSvcCd,omitempty"`
}

// May be one of ALLT, CNTR, IDVD, SCNT, UNFT
type ATMNoteType1Code string

// May be one of ADJU, INSR, LOAD, REMV, UNLD
type ATMOperation1Code string

type ATMReconciliationAdvice2 struct {
	Envt     ATMEnvironment10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Envt"`
	CmdRslt  []ATMCommand8    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CmdRslt,omitempty"`
	CmdCntxt ATMCommand9      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CmdCntxt,omitempty"`
	Tx       ATMTransaction25 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tx"`
}

type ATMReconciliationAdviceV02 struct {
	Hdr                  Header32                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Hdr"`
	PrtctdATMRcncltnAdvc ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PrtctdATMRcncltnAdvc,omitempty"`
	ATMRcncltnAdvc       ATMReconciliationAdvice2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMRcncltnAdvc,omitempty"`
	SctyTrlr             ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SctyTrlr,omitempty"`
}

type ATMTotals1 struct {
	MdiaTp   ATMMediaType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MdiaTp,omitempty"`
	Ccy      ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ccy,omitempty"`
	ATMBal   float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMBal,omitempty"`
	ATMCur   float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMCur,omitempty"`
	ATMBalNb float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMBalNb,omitempty"`
	ATMCurNb float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMCurNb,omitempty"`
}

type ATMTotals3 struct {
	Id      Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Id"`
	AddtlId Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AddtlId,omitempty"`
	Prd     ATMCounterType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Prd"`
	Ccy     ActiveCurrencyCode  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ccy,omitempty"`
	Cnt     float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Cnt"`
	Amt     float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Amt,omitempty"`
}

type ATMTransaction25 struct {
	TpOfOpr    ATMOperation1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TpOfOpr,omitempty"`
	TxId       TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TxId"`
	RcncltnId  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RcncltnId"`
	ATMTtls    []ATMTotals1           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMTtls,omitempty"`
	Csstt      []ATMCassette2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Csstt,omitempty"`
	TxTtls     []ATMTotals3           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TxTtls,omitempty"`
	RtndCard   float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RtndCard,omitempty"`
	AddtlTxInf Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AddtlTxInf,omitempty"`
}

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Brnch,omitempty"`
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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MAC"`
}

type AutomatedTellerMachine8 struct {
	Id       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Id"`
	AddtlId  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AddtlId,omitempty"`
	SeqNb    Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SeqNb,omitempty"`
	BaseCcy  ActiveCurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 BaseCcy"`
	Lctn     PostalAddress17             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Lctn,omitempty"`
	LctnCtgy TransactionEnvironment2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 LctnCtgy,omitempty"`
	Eqpmnt   ATMEquipment1               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Eqpmnt,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMRcncltnAdvc ATMReconciliationAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ATMRcncltnAdvc"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptdCntt,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 GeogcCordints,omitempty"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 UTMCordints,omitempty"`
}

type Header32 struct {
	MsgFctn        ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MsgFctn"`
	PrtcolVrsn     Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PrtcolVrsn"`
	XchgId         Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 XchgId"`
	ReTrnsmssnCntr float64             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CreDtTm"`
	InitgPty       Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 InitgPty"`
	RcptPty        Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RcptPty,omitempty"`
	PrcStat        Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PrcStat,omitempty"`
	Tracblt        []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptdKey"`
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
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KEK,omitempty"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyIdr,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 AttrVal"`
}

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Id,omitempty"`
}

// May be one of CNTE, FMTE, HRDW, NSUP, SECR, SUCC, SYNE, TIMO, UKRF
type TerminalManagementActionResult2Code string

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TracDtTmOut"`
}

// May be one of PRIV, PUBL
type TransactionEnvironment2Code string

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 TxRef"`
}

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 UTMNrthwrd"`
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