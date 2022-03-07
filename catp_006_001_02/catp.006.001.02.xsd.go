// Code generated by download. DO NOT EDIT.

package iso20022_catp_006_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Vrsn"`
}

type ATMContext14 struct {
	SsnRef Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SsnRef,omitempty"`
	Svc    ATMService15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Svc"`
}

type ATMCustomer4 struct {
	Prfl         ATMCustomerProfile4              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Prfl,omitempty"`
	SelctdLang   string                           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SelctdLang,omitempty"`
	Authntcn     []CardholderAuthentication8      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Authntcn"`
	AuthntcnRslt []TransactionVerificationResult5 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AuthntcnRslt,omitempty"`
}

// May be one of CRDF, OREQ, PREQ
type ATMCustomerProfile1Code string

type ATMCustomerProfile4 struct {
	RtrvlMd ATMCustomerProfile1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 RtrvlMd"`
	PrflRef Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrflRef,omitempty"`
	CstmrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CstmrId,omitempty"`
}

// May be one of ALRM, BRCD, CAMR, CRDD, CRDR, CSHD, CSHI, CSHR, CHCK, CDIS, DPST, DPRN, DOOR, INPM, JRNL, JPRN, SNSR, PSBK, PINR, RPRN, SCAN, RWDR
type ATMDevice2Code string

type ATMEnvironment14 struct {
	Acqrr    Acquirer7                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Acqrr,omitempty"`
	ATMMgrId Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATMMgrId,omitempty"`
	HstgNtty TerminalHosting1         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATM"`
	Cstmr    ATMCustomer4             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Cstmr,omitempty"`
	Card     PaymentCard22            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Card,omitempty"`
}

type ATMEquipment1 struct {
	Manfctr    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Manfctr,omitempty"`
	Mdl        Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Mdl,omitempty"`
	SrlNb      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SrlNb,omitempty"`
	ApplPrvdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ApplPrvdr,omitempty"`
	ApplNm     Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ApplNm,omitempty"`
	ApplVrsn   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ApplVrsn,omitempty"`
	ApprvlNb   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ApprvlNb,omitempty"`
	CfgtnParam []ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CfgtnParam,omitempty"`
}

type ATMInquiryRequest2 struct {
	Envt  ATMEnvironment14 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Envt"`
	Cntxt ATMContext14     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Cntxt"`
	Tx    ATMTransaction29 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Tx"`
}

type ATMInquiryRequestV02 struct {
	Hdr              Header31                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Hdr"`
	PrtctdATMNqryReq ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtctdATMNqryReq,omitempty"`
	ATMNqryReq       ATMInquiryRequest2       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATMNqryReq,omitempty"`
	SctyTrlr         ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SctyTrlr,omitempty"`
}

// May be one of CARD, COIN, CMDT, CPNS, NOTE, STMP, UDTM
type ATMMediaType1Code string

// May be one of CARD, COIN, CMDT, CPNS, NOTE, STMP, UDTM, CHCK
type ATMMediaType2Code string

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 HstSvcCd,omitempty"`
}

type ATMService15 struct {
	SvcRef     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SvcRef,omitempty"`
	ATMSvcCd   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATMSvcCd,omitempty"`
	SvcTp      ATMServiceType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SvcTp"`
	SvcVarntId []Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SvcVarntId,omitempty"`
}

// May be one of ASTS, CDVF, DCCS, XRTD, XRTW, EMVS, CMPF, BLCQ
type ATMServiceType3Code string

type ATMTransaction29 struct {
	TxId           TransactionIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TxId"`
	AcctData       CardAccount7             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AcctData,omitempty"`
	PrtctdAcctData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtctdAcctData,omitempty"`
	TtlReqdAmt     AmountAndCurrency1       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TtlReqdAmt,omitempty"`
	DtldReqdAmt    DetailedAmount12         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DtldReqdAmt,omitempty"`
	ICCRltdData    Max10000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ICCRltdData,omitempty"`
}

// May be one of ACSL, ENTR, IMAC, IMPL, NOSL, TPSL
type AccountChoiceMethod1Code string

type AccountIdentification31Choice struct {
	IBAN     IBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 IBAN,omitempty"`
	BBAN     BBANIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BBAN,omitempty"`
	UPIC     UPICIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 UPIC,omitempty"`
	DmstAcct SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DmstAcct,omitempty"`
}

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Brnch,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

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
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Param,omitempty"`
}

type AmountAndCurrency1 struct {
	Amt float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Amt"`
	Ccy ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ccy,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of TOKA, BIOM, MOBL, OTHR, FPIN, NPIN, PSWD, SCRT, SCNL
type AuthenticationMethod7Code string

type AutomatedTellerMachine10 struct {
	Id       Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Id"`
	AddtlId  Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AddtlId,omitempty"`
	SeqNb    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SeqNb,omitempty"`
	BaseCcy  ActiveCurrencyCode              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BaseCcy"`
	Lctn     PostalAddress17                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Lctn,omitempty"`
	LctnCtgy TransactionEnvironment2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 LctnCtgy,omitempty"`
	Cpblties PointOfInteractionCapabilities7 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Cpblties,omitempty"`
	Eqpmnt   ATMEquipment1                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Eqpmnt,omitempty"`
	AvlblDvc []ATMDevice2Code                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AvlblDvc,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardAccount7 struct {
	SelctnMtd    AccountChoiceMethod1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SelctnMtd,omitempty"`
	SelctdAcctTp CardAccountType3Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SelctdAcctTp,omitempty"`
	AcctNm       Max70Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AcctNm,omitempty"`
	AcctOwnr     NameAndAddress3               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AcctOwnr,omitempty"`
	Ccy          ActiveCurrencyCode            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ccy,omitempty"`
	AcctIdr      AccountIdentification31Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AcctIdr,omitempty"`
	Svcr         PartyIdentification72Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Svcr,omitempty"`
}

// May be one of CTDP, CHCK, CRDT, CURR, CDBT, DFLT, EPRS, HEQL, ISTL, INVS, LCDT, MBNW, MNMK, MNMC, MTGL, RTRM, RVLV, SVNG, STBD, UVRL, PRPD, FLTC
type CardAccountType3Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of ECTL, CICC, MGST, CTLS
type CardDataReading4Code string

// May be one of FFLB, SFLB, NFLB
type CardFallback1Code string

type CardholderAuthentication8 struct {
	AuthntcnMtd       AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AuthntcnMtd"`
	TknReqd           bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TknReqd,omitempty"`
	AuthntcnVal       Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AuthntcnVal,omitempty"`
	PrtctdAuthntcnVal ContentInformationType10  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtctdAuthntcnVal,omitempty"`
	CrdhldrOnLinePIN  OnLinePIN5                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CrdhldrOnLinePIN,omitempty"`
}

// May be one of NPIN, FCPN, FEPN, FDSG, FBIO, FBIG, PKIS, PCOD
type CardholderVerificationCapability3Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 EnvlpdData"`
}

type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type DetailedAmount12 struct {
	AmtToDspns float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AmtToDspns"`
	Ccy        ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ccy,omitempty"`
	Fees       []DetailedAmount13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Fees,omitempty"`
	Dontn      []DetailedAmount13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Dontn,omitempty"`
}

type DetailedAmount13 struct {
	Amt  float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Amt"`
	Ccy  ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ccy,omitempty"`
	Labl Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Labl,omitempty"`
}

type DisplayCapabilities5 struct {
	Dstn      []UserInterface5Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Dstn"`
	AvlblFrmt []OutputFormat1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AvlblFrmt,omitempty"`
	NbOfLines float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NbOfLines,omitempty"`
	LineWidth float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 LineWidth,omitempty"`
	AvlblLang []string             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AvlblLang,omitempty"`
}

type Document struct {
	ATMNqryReq ATMInquiryRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ATMNqryReq"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptdCntt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Issr,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 GeogcCordints,omitempty"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 UTMCordints,omitempty"`
}

type Header31 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Tracblt,omitempty"`
}

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptdKey"`
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

// May be no more than 3 items long
type Max3Text string

// May be no more than 45 items long
type Max45Text string

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

type NameAndAddress3 struct {
	Nm  Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Adr"`
}

type OnLinePIN5 struct {
	NcrptdPINBlck ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptdPINBlck"`
	PINFrmt       PINFormat4Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PINFrmt"`
	AddtlInpt     Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AddtlInpt,omitempty"`
}

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

// May be one of ANSI, BNCM, BKSY, DBLD, DBLC, ECI2, ECI3, EMVS, IBM3, ISO0, ISO1, ISO2, ISO3, ISO4, ISO5, VIS2, VIS3
type PINFormat4Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BPddg,omitempty"`
}

type PartyIdentification72Choice struct {
	AnyBIC  AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AnyBIC,omitempty"`
	PrtryId GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtryId,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PaymentCard22 struct {
	CardDataNtryMd CardDataReading1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardDataNtryMd"`
	FllbckInd      CardFallback1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 FllbckInd,omitempty"`
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData18          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PlainCardData,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardCcyCd,omitempty"`
}

type PlainCardData18 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PAN,omitempty"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 XpryDt,omitempty"`
	SvcCd     Exact3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SvcCd,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Trck3,omitempty"`
	CrdhldrNm Max45Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CrdhldrNm,omitempty"`
}

type PointOfInteractionCapabilities7 struct {
	CardRdData       []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardRdData,omitempty"`
	CardWrtData      []CardDataReading4Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardWrtData,omitempty"`
	Authntcn         []CardholderVerificationCapability3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Authntcn,omitempty"`
	PINLngthCpblties float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PINLngthCpblties,omitempty"`
	ApprvlCdLngth    float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 ApprvlCdLngth,omitempty"`
	MxScrptLngth     float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MxScrptLngth,omitempty"`
	CardCaptrCpbl    bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CardCaptrCpbl,omitempty"`
	WdrwlMdia        []ATMMediaType1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 WdrwlMdia,omitempty"`
	DpstdMdia        []ATMMediaType2Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 DpstdMdia,omitempty"`
	MsgCpblties      []DisplayCapabilities5                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 MsgCpblties,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ctry"`
}

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyTrnsprt,omitempty"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KEK,omitempty"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyIdr,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AttrVal"`
}

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Id"`
}

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TracDtTmOut"`
}

// May be one of PRIV, PUBL
type TransactionEnvironment2Code string

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 TxRef"`
}

type TransactionVerificationResult5 struct {
	Mtd         AuthenticationMethod7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Mtd"`
	VrfctnNtty  AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 VrfctnNtty,omitempty"`
	Rslt        Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Rslt,omitempty"`
	AddtlRslt   Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AddtlRslt,omitempty"`
	AuthntcnTkn Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 AuthntcnTkn,omitempty"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 UTMNrthwrd"`
}

// May be one of CDSP, CRCP, CRDO
type UserInterface5Code string

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
