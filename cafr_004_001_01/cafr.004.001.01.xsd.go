// Code generated by download. DO NOT EDIT.

package iso20022_cafr_004_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AdditionalData1 struct {
	Tp  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tp,omitempty"`
	Val Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Val,omitempty"`
}

type AdditionalFee1 struct {
	Tp         TypeOfAmount10Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tp"`
	OthrTp     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrTp,omitempty"`
	FeePrgm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 FeePrgm,omitempty"`
	FeeDscrptr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 FeeDscrptr,omitempty"`
	Amt        FeeAmount2         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Amt"`
	Labl       Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Labl,omitempty"`
}

type AdditionalInformation22 struct {
	Rcpt PartyType19Code      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Rcpt,omitempty"`
	Trgt []UserInterface8Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Trgt,omitempty"`
	Frmt OutputFormat4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Frmt,omitempty"`
	Tp   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tp,omitempty"`
	Val  Max20KText           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Val"`
}

type Address1 struct {
	AdrLine1       Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AdrLine1,omitempty"`
	AdrLine2       Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AdrLine2,omitempty"`
	StrtNm         Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 StrtNm,omitempty"`
	BldgNb         Max16Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 BldgNb,omitempty"`
	PstlCd         Max16Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PstlCd,omitempty"`
	TwnNm          Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TwnNm,omitempty"`
	CtrySubDvsnMnr Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CtrySubDvsnMnr,omitempty"`
	CtrySubDvsnMjr Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CtrySubDvsnMjr,omitempty"`
	Ctry           Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Ctry,omitempty"`
}

type BatchManagementInformation1 struct {
	ColltnId         Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 ColltnId,omitempty"`
	BtchId           Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 BtchId"`
	MsgSeqNb         Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MsgSeqNb,omitempty"`
	MsgChcksmInptVal Max140Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MsgChcksmInptVal,omitempty"`
}

type CardData3 struct {
	PAN           Max19NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PAN,omitempty"`
	PrtctdPANInd  bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PrtctdPANInd,omitempty"`
	CardSeqNb     Min2Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardSeqNb,omitempty"`
	FctvDt        Max10Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 FctvDt,omitempty"`
	XpryDt        Exact4NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 XpryDt,omitempty"`
	PmtAcctRef    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PmtAcctRef,omitempty"`
	PANRefIdr     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PANRefIdr,omitempty"`
	PANAcctRg     Max19NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PANAcctRg,omitempty"`
	CardCtryCd    ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardCtryCd,omitempty"`
	CardPdctTp    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardPdctTp,omitempty"`
	CardPdctSubTp Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardPdctSubTp,omitempty"`
	CardPrtflIdr  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardPrtflIdr,omitempty"`
	AddtlCardData Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlCardData,omitempty"`
}

type CardProgrammeMode1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Id"`
}

type Cardholder15 struct {
	CrdhldrNm CardholderName1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CrdhldrNm,omitempty"`
	Id        []Credentials1  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Id,omitempty"`
	Adr       Address1        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Adr,omitempty"`
	CtctInf   Contact1        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CtctInf,omitempty"`
	DtOfBirth ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 DtOfBirth,omitempty"`
}

type CardholderName1 struct {
	Nm         Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Nm,omitempty"`
	GvnNm      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 GvnNm,omitempty"`
	MddlInitls Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MddlInitls,omitempty"`
	LastNm     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 LastNm,omitempty"`
}

type Contact1 struct {
	Nm            Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Nm,omitempty"`
	HomePhneNb    PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 HomePhneNb,omitempty"`
	BizPhneNb     PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 BizPhneNb,omitempty"`
	MobPhneNb     PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MobPhneNb,omitempty"`
	OthrPhneNb    PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrPhneNb,omitempty"`
	PrsnlEmailAdr Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PrsnlEmailAdr,omitempty"`
	BizEmailAdr   Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 BizEmailAdr,omitempty"`
	OthrEmailAdr  Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrEmailAdr,omitempty"`
	Lang          ISO2ALanguageCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Lang,omitempty"`
}

type ContentInformationType20 struct {
	MACData MACData1          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MACData"`
	MAC     Max8HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MAC"`
}

type Context8 struct {
	TxCntxt TransactionContext5 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TxCntxt,omitempty"`
}

type Credentials1 struct {
	IdCd     Identification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 IdCd"`
	OthrIdCd Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrIdCd,omitempty"`
	IdVal    Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 IdVal"`
}

type Document struct {
	FrdDspstnRspn FraudDispositionResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 FrdDspstnRspn"`
}

type Environment10 struct {
	Acqrr   PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Acqrr,omitempty"`
	Sndr    PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Sndr,omitempty"`
	Rcvr    PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Rcvr,omitempty"`
	Card    CardData3              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Card,omitempty"`
	Crdhldr Cardholder15           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Crdhldr,omitempty"`
	Tkn     Token1                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tkn,omitempty"`
}

// Must match the pattern ([0-9A-F][0-9A-F]){1}
type Exact1HexBinaryText string

// Must match the pattern [a-zA-Z0-9]{2}
type Exact2AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type FeeAmount2 struct {
	Amt      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Amt"`
	Ccy      ISO3NumericCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Ccy,omitempty"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 XchgRate,omitempty"`
	QtnDt    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 QtnDt,omitempty"`
	Sgn      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Sgn,omitempty"`
}

type FraudDispositionResponse1 struct {
	Envt        Environment10        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Envt"`
	Cntxt       Context8             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Cntxt,omitempty"`
	Tx          Transaction93        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tx"`
	PrcgRslt    ProcessingResult4    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PrcgRslt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 SplmtryData,omitempty"`
}

type FraudDispositionResponseV01 struct {
	Hdr      Header48                  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Hdr"`
	Body     FraudDispositionResponse1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Body"`
	SctyTrlr ContentInformationType20  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 SctyTrlr,omitempty"`
}

type GenericIdentification172 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Id"`
	Tp     PartyType17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tp,omitempty"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrTp,omitempty"`
	Assgnr PartyType18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 ShrtNm,omitempty"`
}

type Header48 struct {
	MsgFctn        MessageFunction29Code       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 MsgFctn"`
	PrtcolVrsn     Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PrtcolVrsn"`
	XchgId         Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 XchgId,omitempty"`
	ReTrnsmssnCntr Max3NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CreDtTm"`
	BtchMgmtInf    BatchManagementInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 BtchMgmtInf,omitempty"`
	InitgPty       GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 InitgPty"`
	RcptPty        GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 RcptPty,omitempty"`
	TracData       []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TracData,omitempty"`
	Tracblt        []Traceability7             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Tracblt,omitempty"`
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCurrencyCode string

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

// May be one of DRID, NTID, PASS, SSYN, ARNB, OTHP, OTHN, EMAL, PHNB
type Identification2Code string

type MACData1 struct {
	Ctrl         Exact1HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Ctrl"`
	KeySetIdr    Max8NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 KeySetIdr"`
	DrvdInf      Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 DrvdInf,omitempty"`
	Algo         Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Algo"`
	KeyLngth     Max4NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 KeyLngth,omitempty"`
	KeyPrtcn     Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 KeyPrtcn,omitempty"`
	KeyIndx      Max5NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 KeyIndx,omitempty"`
	PddgMtd      Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PddgMtd,omitempty"`
	InitlstnVctr Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 InitlstnVctr,omitempty"`
}

// May be no more than 10 items long
type Max10Text string

// Must match the pattern [0-9]{1,11}
type Max11NumericText string

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 140 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// May be no more than 16 items long
type Max16Text string

// Must match the pattern [0-9]{1,19}
type Max19NumericText string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 20000 items long
type Max20KText string

// May be no more than 256 items long
type Max256Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,32}
type Max32HexBinaryText string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [0-9]{1,4}
type Max4NumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1,8}
type Max8HexBinaryText string

// Must match the pattern [0-9]{1,8}
type Max8NumericText string

// May be one of ADVC, NOTI
type MessageFunction29Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

// May be one of FLNM, MREF, OTHN, OTHP, SMSI, TEXT, URLI, HTML
type OutputFormat4Code string

type PartyIdentification197 struct {
	Id      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Id"`
	Assgnr  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Assgnr,omitempty"`
	Ctry    ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Ctry,omitempty"`
	ShrtNm  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 ShrtNm,omitempty"`
	AddtlId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlId,omitempty"`
}

// May be one of OTHN, OTHP, ACQR, ACQP, CISS, CISP, AGNT
type PartyType17Code string

// May be one of ACQR, CISS, CSCH, AGNT
type PartyType18Code string

// May be one of ACCP, ACQR, ACQP, CISS, CISP, AGNT, OTHN, OTHP
type PartyType19Code string

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type ProcessingResult4 struct {
	RsltData      ResultData5               `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 RsltData,omitempty"`
	OrgnlRsltData ResultData1               `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OrgnlRsltData,omitempty"`
	AddtlInf      []AdditionalInformation22 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlInf,omitempty"`
}

// May be one of PRCS, UNPR, UNRV, REJT, TECH, OTHN, OTHP
type Response8Code string

type ResultData1 struct {
	Rslt         Response8Code          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Rslt,omitempty"`
	OthrRslt     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrRslt,omitempty"`
	RsltDtls     Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 RsltDtls"`
	OthrRsltDtls Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrRsltDtls,omitempty"`
	AddtlRsltInf []AdditionalData1      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlRsltInf,omitempty"`
}

type ResultData5 struct {
	Rslt         Response8Code          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Rslt,omitempty"`
	OthrRslt     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 OthrRslt,omitempty"`
	RsltDtls     Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 RsltDtls"`
	AddtlRsltInf []AdditionalData1      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlRsltInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Token1 struct {
	PmtTkn        Max19NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 PmtTkn,omitempty"`
	TknXpryDt     Exact4NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TknXpryDt,omitempty"`
	TknRqstrId    Max11NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TknRqstrId,omitempty"`
	TknAssrncData Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TknAssrncData,omitempty"`
	TknAssrncMtd  Max2NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TknAssrncMtd,omitempty"`
	TknInittdInd  bool              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TknInittdInd,omitempty"`
}

type Traceability7 struct {
	RlayId      GenericIdentification172 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 RlayId"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TracDtTmIn,omitempty"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 TracDtTmOut,omitempty"`
}

type Transaction93 struct {
	FrdTxId   Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 FrdTxId"`
	AddtlFees []AdditionalFee1  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlFees,omitempty"`
	AddtlData []AdditionalData1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 AddtlData,omitempty"`
}

type TransactionContext5 struct {
	CardPrgrmmApld CardProgrammeMode1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.004.001.01 CardPrgrmmApld,omitempty"`
}

// May be one of INTC, FEEP, OTHN, OTHP, FEEA
type TypeOfAmount10Code string

// May be one of DSPU, FILE, LOGF, OTHP, OTHN
type UserInterface8Code string

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
