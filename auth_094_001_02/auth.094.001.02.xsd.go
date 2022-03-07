// Code generated by download. DO NOT EDIT.

package iso20022_auth_094_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// May be one of ANYM
type AnyMIC1Code string

// May be one of GBBK, BOND, CASH, COMM, INSU, LCRE, OTHR, PHYS, SECU, STCF
type CollateralType6Code string

type CorporateSectorCriteria5 struct {
	FISctr  []FinancialPartySectorType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 FISctr,omitempty"`
	NFISctr []NACEDomainIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NFISctr,omitempty"`
	NotRptd NotReported1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NotRptd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateOrBlankQuery2Choice struct {
	Rg      DatePeriod1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Rg,omitempty"`
	NotRptd NotReported1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NotRptd,omitempty"`
}

type DatePeriod1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 FrDt,omitempty"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ToDt"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ToDtTm"`
}

type Document struct {
	SctiesFincgRptgTxQry SecuritiesFinancingReportingTransactionQueryV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 SctiesFincgRptgTxQry"`
}

// May be one of SBSC, MGLD, SLEB, REPO
type ExposureType10Code string

// May be one of AIFD, CSDS, CCPS, CDTI, INUN, ORPI, INVF, REIN, UCIT
type FinancialPartySectorType2Code string

// May be one of DAIL, WEEK, MNTH, ADHO
type Frequency14Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Issr,omitempty"`
}

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

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// May be no more than 1000 items long
type Max1000Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 50 items long
type Max50Text string

// May be no more than 70 items long
type Max70Text string

// Must match the pattern [A-U]{1,1}
type NACEDomainIdentifier string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Adr,omitempty"`
}

// May be one of NORP
type NotReported1Code string

// May be one of ANDD, ORRR
type Operation3Code string

type PartyIdentification121Choice struct {
	AnyBIC     AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AnyBIC,omitempty"`
	LglNttyIdr LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 LglNttyIdr,omitempty"`
	NmAndAdr   NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NmAndAdr,omitempty"`
	PrtryId    GenericIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 PrtryId,omitempty"`
}

// May be one of OTHR, NFIN, FIIN, CCPS
type PartyNatureType1Code string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Ctry"`
}

type SecuritiesFinancingReportingTransactionQueryV02 struct {
	RqstngAuthrty PartyIdentification121Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 RqstngAuthrty"`
	TradQryData   TradeReportQuery13Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TradQryData"`
	SplmtryData   []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 SplmtryData,omitempty"`
}

type SecuritiesTradeVenueCriteria1Choice struct {
	MIC    []MICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 MIC,omitempty"`
	AnyMIC AnyMIC1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AnyMIC,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeAdditionalQueryCriteria7 struct {
	ActnTp      []TransactionOperationType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ActnTp,omitempty"`
	ExctnVn     SecuritiesTradeVenueCriteria1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ExctnVn,omitempty"`
	NtrOfCtrPty []PartyNatureType1Code              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NtrOfCtrPty,omitempty"`
	CorpSctr    []CorporateSectorCriteria5          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 CorpSctr,omitempty"`
}

type TradeDateTimeQueryCriteria2 struct {
	RptgDtTm  DateTimePeriod1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 RptgDtTm,omitempty"`
	ExctnDtTm DateTimePeriod1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ExctnDtTm,omitempty"`
	MtrtyDt   DateOrBlankQuery2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 MtrtyDt,omitempty"`
	TermntnDt DateOrBlankQuery2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TermntnDt,omitempty"`
}

type TradePartyIdentificationQuery8 struct {
	LEI     []LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 LEI,omitempty"`
	AnyBIC  []AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AnyBIC,omitempty"`
	ClntId  []Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ClntId,omitempty"`
	NotRptd NotReported1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NotRptd,omitempty"`
}

type TradePartyIdentificationQuery9 struct {
	LEI     []LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 LEI,omitempty"`
	CtryCd  []CountryCode             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 CtryCd,omitempty"`
	AnyBIC  []AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AnyBIC,omitempty"`
	ClntId  []Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 ClntId,omitempty"`
	NotRptd NotReported1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 NotRptd,omitempty"`
}

type TradePartyQueryCriteria5 struct {
	Oprtr           Operation3Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Oprtr"`
	RptgCtrPty      TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 RptgCtrPty,omitempty"`
	RptgCtrPtyBrnch TradePartyIdentificationQuery9 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 RptgCtrPtyBrnch,omitempty"`
	OthrCtrPty      TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 OthrCtrPty,omitempty"`
	OthrCtrPtyBrnch TradePartyIdentificationQuery9 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 OthrCtrPtyBrnch,omitempty"`
	Bnfcry          TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Bnfcry,omitempty"`
	SubmitgAgt      TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 SubmitgAgt,omitempty"`
	Brkr            TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Brkr,omitempty"`
	CCP             TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 CCP,omitempty"`
	AgtLndr         TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AgtLndr,omitempty"`
	TrptyAgt        TradePartyIdentificationQuery8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TrptyAgt,omitempty"`
}

type TradeQueryCriteria10 struct {
	TradLifeCyclHstry bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TradLifeCyclHstry"`
	OutsdngTradInd    bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 OutsdngTradInd"`
	TradPtyCrit       TradePartyQueryCriteria5      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TradPtyCrit,omitempty"`
	TradTpCrit        TradeTypeQueryCriteria2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TradTpCrit,omitempty"`
	TmCrit            TradeDateTimeQueryCriteria2   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 TmCrit,omitempty"`
	OthrCrit          TradeAdditionalQueryCriteria7 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 OthrCrit,omitempty"`
}

type TradeQueryExecutionFrequency3 struct {
	FrqcyTp   Frequency14Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 FrqcyTp"`
	DlvryDay  []WeekDay3Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 DlvryDay,omitempty"`
	DayOfMnth []float64       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 DayOfMnth,omitempty"`
}

type TradeRecurrentQuery5 struct {
	QryTp    Max1000Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 QryTp"`
	Frqcy    TradeQueryExecutionFrequency3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Frqcy"`
	VldUntil ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 VldUntil"`
}

type TradeReportQuery13Choice struct {
	AdHocQry TradeQueryCriteria10 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 AdHocQry,omitempty"`
	RcrntQry TradeRecurrentQuery5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 RcrntQry,omitempty"`
}

type TradeTypeQueryCriteria2 struct {
	Oprtr           Operation3Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 Oprtr"`
	SctiesFincgTxTp []ExposureType10Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 SctiesFincgTxTp,omitempty"`
	CollCmpntTp     []CollateralType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.094.001.02 CollCmpntTp,omitempty"`
}

// May be one of REUU, COLU, CORR, ETRM, VALU, POSC, NEWT, MODI, MARU, EROR
type TransactionOperationType6Code string

// May be one of ALLD, XBHL, IBHL, FRID, MOND, SATD, SUND, THUD, TUED, WEDD, WDAY, WEND
type WeekDay3Code string

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