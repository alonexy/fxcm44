package collateralinquiryack

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// CollateralInquiryAck is the fix44 CollateralInquiryAck type, MsgType = BG.
type CollateralInquiryAck struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a CollateralInquiryAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) CollateralInquiryAck {
	return CollateralInquiryAck{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m CollateralInquiryAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a CollateralInquiryAck initialized with the required fields for CollateralInquiryAck.
func New(collinquiryid field.CollInquiryIDField, collinquirystatus field.CollInquiryStatusField) (m CollateralInquiryAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BG"))
	m.Set(collinquiryid)
	m.Set(collinquirystatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg CollateralInquiryAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BG", r
}

// SetAccount sets Account, Tag 1.
func (m CollateralInquiryAck) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m CollateralInquiryAck) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetCurrency sets Currency, Tag 15.
func (m CollateralInquiryAck) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m CollateralInquiryAck) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m CollateralInquiryAck) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m CollateralInquiryAck) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetQuantity sets Quantity, Tag 53.
func (m CollateralInquiryAck) SetQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewQuantity(value, scale))
}

// SetSymbol sets Symbol, Tag 55.
func (m CollateralInquiryAck) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58.
func (m CollateralInquiryAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetSettlDate sets SettlDate, Tag 64.
func (m CollateralInquiryAck) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m CollateralInquiryAck) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m CollateralInquiryAck) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m CollateralInquiryAck) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetNoExecs sets NoExecs, Tag 124.
func (m CollateralInquiryAck) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

// SetSecurityType sets SecurityType, Tag 167.
func (m CollateralInquiryAck) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetSecondaryOrderID sets SecondaryOrderID, Tag 198.
func (m CollateralInquiryAck) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m CollateralInquiryAck) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m CollateralInquiryAck) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m CollateralInquiryAck) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m CollateralInquiryAck) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m CollateralInquiryAck) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m CollateralInquiryAck) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m CollateralInquiryAck) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m CollateralInquiryAck) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m CollateralInquiryAck) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m CollateralInquiryAck) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m CollateralInquiryAck) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m CollateralInquiryAck) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m CollateralInquiryAck) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248.
func (m CollateralInquiryAck) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

// SetLegIssueDate sets LegIssueDate, Tag 249.
func (m CollateralInquiryAck) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

// SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250.
func (m CollateralInquiryAck) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

// SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251.
func (m CollateralInquiryAck) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

// SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252.
func (m CollateralInquiryAck) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

// SetLegFactor sets LegFactor, Tag 253.
func (m CollateralInquiryAck) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

// SetLegRedemptionDate sets LegRedemptionDate, Tag 254.
func (m CollateralInquiryAck) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m CollateralInquiryAck) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetLegCreditRating sets LegCreditRating, Tag 257.
func (m CollateralInquiryAck) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m CollateralInquiryAck) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m CollateralInquiryAck) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m CollateralInquiryAck) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m CollateralInquiryAck) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m CollateralInquiryAck) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m CollateralInquiryAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m CollateralInquiryAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m CollateralInquiryAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m CollateralInquiryAck) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m CollateralInquiryAck) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m CollateralInquiryAck) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m CollateralInquiryAck) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m CollateralInquiryAck) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m CollateralInquiryAck) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526.
func (m CollateralInquiryAck) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m CollateralInquiryAck) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m CollateralInquiryAck) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetNoLegs sets NoLegs, Tag 555.
func (m CollateralInquiryAck) SetNoLegs(v int) {
	m.Set(field.NewNoLegs(v))
}

// SetLegCurrency sets LegCurrency, Tag 556.
func (m CollateralInquiryAck) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

// SetRoundLot sets RoundLot, Tag 561.
func (m CollateralInquiryAck) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

// SetAccountType sets AccountType, Tag 581.
func (m CollateralInquiryAck) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596.
func (m CollateralInquiryAck) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

// SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597.
func (m CollateralInquiryAck) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

// SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598.
func (m CollateralInquiryAck) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

// SetLegInstrRegistry sets LegInstrRegistry, Tag 599.
func (m CollateralInquiryAck) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

// SetLegSymbol sets LegSymbol, Tag 600.
func (m CollateralInquiryAck) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

// SetLegSymbolSfx sets LegSymbolSfx, Tag 601.
func (m CollateralInquiryAck) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

// SetLegSecurityID sets LegSecurityID, Tag 602.
func (m CollateralInquiryAck) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

// SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603.
func (m CollateralInquiryAck) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

// SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604.
func (m CollateralInquiryAck) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetLegProduct sets LegProduct, Tag 607.
func (m CollateralInquiryAck) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

// SetLegCFICode sets LegCFICode, Tag 608.
func (m CollateralInquiryAck) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

// SetLegSecurityType sets LegSecurityType, Tag 609.
func (m CollateralInquiryAck) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

// SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610.
func (m CollateralInquiryAck) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

// SetLegMaturityDate sets LegMaturityDate, Tag 611.
func (m CollateralInquiryAck) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

// SetLegStrikePrice sets LegStrikePrice, Tag 612.
func (m CollateralInquiryAck) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

// SetLegOptAttribute sets LegOptAttribute, Tag 613.
func (m CollateralInquiryAck) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

// SetLegContractMultiplier sets LegContractMultiplier, Tag 614.
func (m CollateralInquiryAck) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

// SetLegCouponRate sets LegCouponRate, Tag 615.
func (m CollateralInquiryAck) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

// SetLegSecurityExchange sets LegSecurityExchange, Tag 616.
func (m CollateralInquiryAck) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

// SetLegIssuer sets LegIssuer, Tag 617.
func (m CollateralInquiryAck) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

// SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618.
func (m CollateralInquiryAck) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

// SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619.
func (m CollateralInquiryAck) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

// SetLegSecurityDesc sets LegSecurityDesc, Tag 620.
func (m CollateralInquiryAck) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

// SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621.
func (m CollateralInquiryAck) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

// SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622.
func (m CollateralInquiryAck) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

// SetLegRatioQty sets LegRatioQty, Tag 623.
func (m CollateralInquiryAck) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

// SetLegSide sets LegSide, Tag 624.
func (m CollateralInquiryAck) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m CollateralInquiryAck) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m CollateralInquiryAck) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetPool sets Pool, Tag 691.
func (m CollateralInquiryAck) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetNoUnderlyings sets NoUnderlyings, Tag 711.
func (m CollateralInquiryAck) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

// SetClearingBusinessDate sets ClearingBusinessDate, Tag 715.
func (m CollateralInquiryAck) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

// SetSettlSessID sets SettlSessID, Tag 716.
func (m CollateralInquiryAck) SetSettlSessID(v enum.SettlSessID) {
	m.Set(field.NewSettlSessID(v))
}

// SetSettlSessSubID sets SettlSessSubID, Tag 717.
func (m CollateralInquiryAck) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

// SetResponseTransportType sets ResponseTransportType, Tag 725.
func (m CollateralInquiryAck) SetResponseTransportType(v enum.ResponseTransportType) {
	m.Set(field.NewResponseTransportType(v))
}

// SetResponseDestination sets ResponseDestination, Tag 726.
func (m CollateralInquiryAck) SetResponseDestination(v string) {
	m.Set(field.NewResponseDestination(v))
}

// SetLegDatedDate sets LegDatedDate, Tag 739.
func (m CollateralInquiryAck) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

// SetLegPool sets LegPool, Tag 740.
func (m CollateralInquiryAck) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m CollateralInquiryAck) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetLegSecuritySubType sets LegSecuritySubType, Tag 764.
func (m CollateralInquiryAck) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

// SetTerminationType sets TerminationType, Tag 788.
func (m CollateralInquiryAck) SetTerminationType(v enum.TerminationType) {
	m.Set(field.NewTerminationType(v))
}

// SetQtyType sets QtyType, Tag 854.
func (m CollateralInquiryAck) SetQtyType(v enum.QtyType) {
	m.Set(field.NewQtyType(v))
}

// SetNoEvents sets NoEvents, Tag 864.
func (m CollateralInquiryAck) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m CollateralInquiryAck) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m CollateralInquiryAck) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m CollateralInquiryAck) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m CollateralInquiryAck) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetNoTrades sets NoTrades, Tag 897.
func (m CollateralInquiryAck) SetNoTrades(f NoTradesRepeatingGroup) {
	m.SetGroup(f)
}

// SetMarginRatio sets MarginRatio, Tag 898.
func (m CollateralInquiryAck) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

// SetCollInquiryID sets CollInquiryID, Tag 909.
func (m CollateralInquiryAck) SetCollInquiryID(v string) {
	m.Set(field.NewCollInquiryID(v))
}

// SetTotNumReports sets TotNumReports, Tag 911.
func (m CollateralInquiryAck) SetTotNumReports(v int) {
	m.Set(field.NewTotNumReports(v))
}

// SetAgreementDesc sets AgreementDesc, Tag 913.
func (m CollateralInquiryAck) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

// SetAgreementID sets AgreementID, Tag 914.
func (m CollateralInquiryAck) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

// SetAgreementDate sets AgreementDate, Tag 915.
func (m CollateralInquiryAck) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

// SetStartDate sets StartDate, Tag 916.
func (m CollateralInquiryAck) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

// SetEndDate sets EndDate, Tag 917.
func (m CollateralInquiryAck) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

// SetAgreementCurrency sets AgreementCurrency, Tag 918.
func (m CollateralInquiryAck) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

// SetDeliveryType sets DeliveryType, Tag 919.
func (m CollateralInquiryAck) SetDeliveryType(v enum.DeliveryType) {
	m.Set(field.NewDeliveryType(v))
}

// SetNoCollInquiryQualifier sets NoCollInquiryQualifier, Tag 938.
func (m CollateralInquiryAck) SetNoCollInquiryQualifier(f NoCollInquiryQualifierRepeatingGroup) {
	m.SetGroup(f)
}

// SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942.
func (m CollateralInquiryAck) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

// SetCollInquiryStatus sets CollInquiryStatus, Tag 945.
func (m CollateralInquiryAck) SetCollInquiryStatus(v enum.CollInquiryStatus) {
	m.Set(field.NewCollInquiryStatus(v))
}

// SetCollInquiryResult sets CollInquiryResult, Tag 946.
func (m CollateralInquiryAck) SetCollInquiryResult(v enum.CollInquiryResult) {
	m.Set(field.NewCollInquiryResult(v))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m CollateralInquiryAck) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955.
func (m CollateralInquiryAck) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

// SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956.
func (m CollateralInquiryAck) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

// SetFXCMSymID sets FXCMSymID, Tag 9000.
func (m CollateralInquiryAck) SetFXCMSymID(v string) {
	m.Set(field.NewFXCMSymID(v))
}

// SetFXCMSymPrecision sets FXCMSymPrecision, Tag 9001.
func (m CollateralInquiryAck) SetFXCMSymPrecision(v int) {
	m.Set(field.NewFXCMSymPrecision(v))
}

// SetFXCMSymPointSize sets FXCMSymPointSize, Tag 9002.
func (m CollateralInquiryAck) SetFXCMSymPointSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymPointSize(value, scale))
}

// SetFXCMSymInterestBuy sets FXCMSymInterestBuy, Tag 9003.
func (m CollateralInquiryAck) SetFXCMSymInterestBuy(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestBuy(value, scale))
}

// SetFXCMSymInterestSell sets FXCMSymInterestSell, Tag 9004.
func (m CollateralInquiryAck) SetFXCMSymInterestSell(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestSell(value, scale))
}

// SetFXCMSymSortOrder sets FXCMSymSortOrder, Tag 9005.
func (m CollateralInquiryAck) SetFXCMSymSortOrder(v int) {
	m.Set(field.NewFXCMSymSortOrder(v))
}

// SetFXCMSymMarginRatio sets FXCMSymMarginRatio, Tag 9006.
func (m CollateralInquiryAck) SetFXCMSymMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymMarginRatio(value, scale))
}

// SetFXCMRequestRejectReason sets FXCMRequestRejectReason, Tag 9025.
func (m CollateralInquiryAck) SetFXCMRequestRejectReason(v enum.FXCMRequestRejectReason) {
	m.Set(field.NewFXCMRequestRejectReason(v))
}

// SetFXCMErrorDetails sets FXCMErrorDetails, Tag 9029.
func (m CollateralInquiryAck) SetFXCMErrorDetails(v string) {
	m.Set(field.NewFXCMErrorDetails(v))
}

// SetFXCMSubscriptionStatus sets FXCMSubscriptionStatus, Tag 9076.
func (m CollateralInquiryAck) SetFXCMSubscriptionStatus(v string) {
	m.Set(field.NewFXCMSubscriptionStatus(v))
}

// SetFXCMProductID sets FXCMProductID, Tag 9080.
func (m CollateralInquiryAck) SetFXCMProductID(v int) {
	m.Set(field.NewFXCMProductID(v))
}

// SetFXCMCondDistStop sets FXCMCondDistStop, Tag 9090.
func (m CollateralInquiryAck) SetFXCMCondDistStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistStop(value, scale))
}

// SetFXCMCondDistLimit sets FXCMCondDistLimit, Tag 9091.
func (m CollateralInquiryAck) SetFXCMCondDistLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistLimit(value, scale))
}

// SetFXCMCondDistEntryStop sets FXCMCondDistEntryStop, Tag 9092.
func (m CollateralInquiryAck) SetFXCMCondDistEntryStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryStop(value, scale))
}

// SetFXCMCondDistEntryLimit sets FXCMCondDistEntryLimit, Tag 9093.
func (m CollateralInquiryAck) SetFXCMCondDistEntryLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryLimit(value, scale))
}

// SetFXCMMaxQuantity sets FXCMMaxQuantity, Tag 9094.
func (m CollateralInquiryAck) SetFXCMMaxQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMaxQuantity(value, scale))
}

// SetFXCMMinQuantity sets FXCMMinQuantity, Tag 9095.
func (m CollateralInquiryAck) SetFXCMMinQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMinQuantity(value, scale))
}

// SetFXCMTradingStatus sets FXCMTradingStatus, Tag 9096.
func (m CollateralInquiryAck) SetFXCMTradingStatus(v string) {
	m.Set(field.NewFXCMTradingStatus(v))
}

// GetAccount gets Account, Tag 1.
func (m CollateralInquiryAck) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m CollateralInquiryAck) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m CollateralInquiryAck) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m CollateralInquiryAck) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m CollateralInquiryAck) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m CollateralInquiryAck) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuantity gets Quantity, Tag 53.
func (m CollateralInquiryAck) GetQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.QuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m CollateralInquiryAck) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m CollateralInquiryAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlDate gets SettlDate, Tag 64.
func (m CollateralInquiryAck) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m CollateralInquiryAck) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m CollateralInquiryAck) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m CollateralInquiryAck) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoExecs gets NoExecs, Tag 124.
func (m CollateralInquiryAck) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSecurityType gets SecurityType, Tag 167.
func (m CollateralInquiryAck) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryOrderID gets SecondaryOrderID, Tag 198.
func (m CollateralInquiryAck) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m CollateralInquiryAck) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m CollateralInquiryAck) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m CollateralInquiryAck) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m CollateralInquiryAck) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m CollateralInquiryAck) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m CollateralInquiryAck) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m CollateralInquiryAck) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m CollateralInquiryAck) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m CollateralInquiryAck) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m CollateralInquiryAck) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m CollateralInquiryAck) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m CollateralInquiryAck) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m CollateralInquiryAck) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248.
func (m CollateralInquiryAck) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssueDate gets LegIssueDate, Tag 249.
func (m CollateralInquiryAck) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250.
func (m CollateralInquiryAck) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251.
func (m CollateralInquiryAck) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252.
func (m CollateralInquiryAck) GetLegRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFactor gets LegFactor, Tag 253.
func (m CollateralInquiryAck) GetLegFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRedemptionDate gets LegRedemptionDate, Tag 254.
func (m CollateralInquiryAck) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m CollateralInquiryAck) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCreditRating gets LegCreditRating, Tag 257.
func (m CollateralInquiryAck) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m CollateralInquiryAck) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m CollateralInquiryAck) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m CollateralInquiryAck) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m CollateralInquiryAck) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m CollateralInquiryAck) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m CollateralInquiryAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m CollateralInquiryAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m CollateralInquiryAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m CollateralInquiryAck) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m CollateralInquiryAck) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m CollateralInquiryAck) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m CollateralInquiryAck) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m CollateralInquiryAck) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m CollateralInquiryAck) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526.
func (m CollateralInquiryAck) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m CollateralInquiryAck) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m CollateralInquiryAck) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegs gets NoLegs, Tag 555.
func (m CollateralInquiryAck) GetNoLegs() (v int, err quickfix.MessageRejectError) {
	var f field.NoLegsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCurrency gets LegCurrency, Tag 556.
func (m CollateralInquiryAck) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoundLot gets RoundLot, Tag 561.
func (m CollateralInquiryAck) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581.
func (m CollateralInquiryAck) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596.
func (m CollateralInquiryAck) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597.
func (m CollateralInquiryAck) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598.
func (m CollateralInquiryAck) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInstrRegistry gets LegInstrRegistry, Tag 599.
func (m CollateralInquiryAck) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSymbol gets LegSymbol, Tag 600.
func (m CollateralInquiryAck) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSymbolSfx gets LegSymbolSfx, Tag 601.
func (m CollateralInquiryAck) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityID gets LegSecurityID, Tag 602.
func (m CollateralInquiryAck) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603.
func (m CollateralInquiryAck) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604.
func (m CollateralInquiryAck) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLegProduct gets LegProduct, Tag 607.
func (m CollateralInquiryAck) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCFICode gets LegCFICode, Tag 608.
func (m CollateralInquiryAck) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityType gets LegSecurityType, Tag 609.
func (m CollateralInquiryAck) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610.
func (m CollateralInquiryAck) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityDate gets LegMaturityDate, Tag 611.
func (m CollateralInquiryAck) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikePrice gets LegStrikePrice, Tag 612.
func (m CollateralInquiryAck) GetLegStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptAttribute gets LegOptAttribute, Tag 613.
func (m CollateralInquiryAck) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplier gets LegContractMultiplier, Tag 614.
func (m CollateralInquiryAck) GetLegContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponRate gets LegCouponRate, Tag 615.
func (m CollateralInquiryAck) GetLegCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityExchange gets LegSecurityExchange, Tag 616.
func (m CollateralInquiryAck) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssuer gets LegIssuer, Tag 617.
func (m CollateralInquiryAck) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618.
func (m CollateralInquiryAck) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619.
func (m CollateralInquiryAck) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityDesc gets LegSecurityDesc, Tag 620.
func (m CollateralInquiryAck) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621.
func (m CollateralInquiryAck) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622.
func (m CollateralInquiryAck) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRatioQty gets LegRatioQty, Tag 623.
func (m CollateralInquiryAck) GetLegRatioQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSide gets LegSide, Tag 624.
func (m CollateralInquiryAck) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m CollateralInquiryAck) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m CollateralInquiryAck) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m CollateralInquiryAck) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyings gets NoUnderlyings, Tag 711.
func (m CollateralInquiryAck) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetClearingBusinessDate gets ClearingBusinessDate, Tag 715.
func (m CollateralInquiryAck) GetClearingBusinessDate() (v string, err quickfix.MessageRejectError) {
	var f field.ClearingBusinessDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessID gets SettlSessID, Tag 716.
func (m CollateralInquiryAck) GetSettlSessID() (v enum.SettlSessID, err quickfix.MessageRejectError) {
	var f field.SettlSessIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessSubID gets SettlSessSubID, Tag 717.
func (m CollateralInquiryAck) GetSettlSessSubID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlSessSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetResponseTransportType gets ResponseTransportType, Tag 725.
func (m CollateralInquiryAck) GetResponseTransportType() (v enum.ResponseTransportType, err quickfix.MessageRejectError) {
	var f field.ResponseTransportTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetResponseDestination gets ResponseDestination, Tag 726.
func (m CollateralInquiryAck) GetResponseDestination() (v string, err quickfix.MessageRejectError) {
	var f field.ResponseDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegDatedDate gets LegDatedDate, Tag 739.
func (m CollateralInquiryAck) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPool gets LegPool, Tag 740.
func (m CollateralInquiryAck) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m CollateralInquiryAck) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecuritySubType gets LegSecuritySubType, Tag 764.
func (m CollateralInquiryAck) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTerminationType gets TerminationType, Tag 788.
func (m CollateralInquiryAck) GetTerminationType() (v enum.TerminationType, err quickfix.MessageRejectError) {
	var f field.TerminationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQtyType gets QtyType, Tag 854.
func (m CollateralInquiryAck) GetQtyType() (v enum.QtyType, err quickfix.MessageRejectError) {
	var f field.QtyTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864.
func (m CollateralInquiryAck) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m CollateralInquiryAck) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m CollateralInquiryAck) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m CollateralInquiryAck) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m CollateralInquiryAck) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTrades gets NoTrades, Tag 897.
func (m CollateralInquiryAck) GetNoTrades() (NoTradesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMarginRatio gets MarginRatio, Tag 898.
func (m CollateralInquiryAck) GetMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollInquiryID gets CollInquiryID, Tag 909.
func (m CollateralInquiryAck) GetCollInquiryID() (v string, err quickfix.MessageRejectError) {
	var f field.CollInquiryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNumReports gets TotNumReports, Tag 911.
func (m CollateralInquiryAck) GetTotNumReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNumReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDesc gets AgreementDesc, Tag 913.
func (m CollateralInquiryAck) GetAgreementDesc() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementID gets AgreementID, Tag 914.
func (m CollateralInquiryAck) GetAgreementID() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDate gets AgreementDate, Tag 915.
func (m CollateralInquiryAck) GetAgreementDate() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStartDate gets StartDate, Tag 916.
func (m CollateralInquiryAck) GetStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.StartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndDate gets EndDate, Tag 917.
func (m CollateralInquiryAck) GetEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementCurrency gets AgreementCurrency, Tag 918.
func (m CollateralInquiryAck) GetAgreementCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeliveryType gets DeliveryType, Tag 919.
func (m CollateralInquiryAck) GetDeliveryType() (v enum.DeliveryType, err quickfix.MessageRejectError) {
	var f field.DeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoCollInquiryQualifier gets NoCollInquiryQualifier, Tag 938.
func (m CollateralInquiryAck) GetNoCollInquiryQualifier() (NoCollInquiryQualifierRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCollInquiryQualifierRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942.
func (m CollateralInquiryAck) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollInquiryStatus gets CollInquiryStatus, Tag 945.
func (m CollateralInquiryAck) GetCollInquiryStatus() (v enum.CollInquiryStatus, err quickfix.MessageRejectError) {
	var f field.CollInquiryStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollInquiryResult gets CollInquiryResult, Tag 946.
func (m CollateralInquiryAck) GetCollInquiryResult() (v enum.CollInquiryResult, err quickfix.MessageRejectError) {
	var f field.CollInquiryResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m CollateralInquiryAck) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955.
func (m CollateralInquiryAck) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956.
func (m CollateralInquiryAck) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymID gets FXCMSymID, Tag 9000.
func (m CollateralInquiryAck) GetFXCMSymID() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSymIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPrecision gets FXCMSymPrecision, Tag 9001.
func (m CollateralInquiryAck) GetFXCMSymPrecision() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPointSize gets FXCMSymPointSize, Tag 9002.
func (m CollateralInquiryAck) GetFXCMSymPointSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymPointSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestBuy gets FXCMSymInterestBuy, Tag 9003.
func (m CollateralInquiryAck) GetFXCMSymInterestBuy() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestBuyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestSell gets FXCMSymInterestSell, Tag 9004.
func (m CollateralInquiryAck) GetFXCMSymInterestSell() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestSellField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymSortOrder gets FXCMSymSortOrder, Tag 9005.
func (m CollateralInquiryAck) GetFXCMSymSortOrder() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymSortOrderField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymMarginRatio gets FXCMSymMarginRatio, Tag 9006.
func (m CollateralInquiryAck) GetFXCMSymMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymMarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMRequestRejectReason gets FXCMRequestRejectReason, Tag 9025.
func (m CollateralInquiryAck) GetFXCMRequestRejectReason() (v enum.FXCMRequestRejectReason, err quickfix.MessageRejectError) {
	var f field.FXCMRequestRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMErrorDetails gets FXCMErrorDetails, Tag 9029.
func (m CollateralInquiryAck) GetFXCMErrorDetails() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMErrorDetailsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSubscriptionStatus gets FXCMSubscriptionStatus, Tag 9076.
func (m CollateralInquiryAck) GetFXCMSubscriptionStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSubscriptionStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMProductID gets FXCMProductID, Tag 9080.
func (m CollateralInquiryAck) GetFXCMProductID() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMProductIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistStop gets FXCMCondDistStop, Tag 9090.
func (m CollateralInquiryAck) GetFXCMCondDistStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistLimit gets FXCMCondDistLimit, Tag 9091.
func (m CollateralInquiryAck) GetFXCMCondDistLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryStop gets FXCMCondDistEntryStop, Tag 9092.
func (m CollateralInquiryAck) GetFXCMCondDistEntryStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryLimit gets FXCMCondDistEntryLimit, Tag 9093.
func (m CollateralInquiryAck) GetFXCMCondDistEntryLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMaxQuantity gets FXCMMaxQuantity, Tag 9094.
func (m CollateralInquiryAck) GetFXCMMaxQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMaxQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMinQuantity gets FXCMMinQuantity, Tag 9095.
func (m CollateralInquiryAck) GetFXCMMinQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMinQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMTradingStatus gets FXCMTradingStatus, Tag 9096.
func (m CollateralInquiryAck) GetFXCMTradingStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m CollateralInquiryAck) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m CollateralInquiryAck) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m CollateralInquiryAck) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m CollateralInquiryAck) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m CollateralInquiryAck) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m CollateralInquiryAck) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasQuantity returns true if Quantity is present, Tag 53.
func (m CollateralInquiryAck) HasQuantity() bool {
	return m.Has(tag.Quantity)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m CollateralInquiryAck) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m CollateralInquiryAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasSettlDate returns true if SettlDate is present, Tag 64.
func (m CollateralInquiryAck) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m CollateralInquiryAck) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m CollateralInquiryAck) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m CollateralInquiryAck) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasNoExecs returns true if NoExecs is present, Tag 124.
func (m CollateralInquiryAck) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m CollateralInquiryAck) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198.
func (m CollateralInquiryAck) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m CollateralInquiryAck) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m CollateralInquiryAck) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m CollateralInquiryAck) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m CollateralInquiryAck) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m CollateralInquiryAck) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m CollateralInquiryAck) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m CollateralInquiryAck) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m CollateralInquiryAck) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m CollateralInquiryAck) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m CollateralInquiryAck) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m CollateralInquiryAck) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m CollateralInquiryAck) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m CollateralInquiryAck) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248.
func (m CollateralInquiryAck) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

// HasLegIssueDate returns true if LegIssueDate is present, Tag 249.
func (m CollateralInquiryAck) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

// HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250.
func (m CollateralInquiryAck) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

// HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251.
func (m CollateralInquiryAck) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

// HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252.
func (m CollateralInquiryAck) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

// HasLegFactor returns true if LegFactor is present, Tag 253.
func (m CollateralInquiryAck) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

// HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254.
func (m CollateralInquiryAck) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m CollateralInquiryAck) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasLegCreditRating returns true if LegCreditRating is present, Tag 257.
func (m CollateralInquiryAck) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m CollateralInquiryAck) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m CollateralInquiryAck) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m CollateralInquiryAck) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m CollateralInquiryAck) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m CollateralInquiryAck) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m CollateralInquiryAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m CollateralInquiryAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m CollateralInquiryAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m CollateralInquiryAck) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m CollateralInquiryAck) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m CollateralInquiryAck) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m CollateralInquiryAck) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m CollateralInquiryAck) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m CollateralInquiryAck) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526.
func (m CollateralInquiryAck) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m CollateralInquiryAck) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m CollateralInquiryAck) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasNoLegs returns true if NoLegs is present, Tag 555.
func (m CollateralInquiryAck) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

// HasLegCurrency returns true if LegCurrency is present, Tag 556.
func (m CollateralInquiryAck) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

// HasRoundLot returns true if RoundLot is present, Tag 561.
func (m CollateralInquiryAck) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m CollateralInquiryAck) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596.
func (m CollateralInquiryAck) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

// HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597.
func (m CollateralInquiryAck) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

// HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598.
func (m CollateralInquiryAck) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

// HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599.
func (m CollateralInquiryAck) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

// HasLegSymbol returns true if LegSymbol is present, Tag 600.
func (m CollateralInquiryAck) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

// HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601.
func (m CollateralInquiryAck) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

// HasLegSecurityID returns true if LegSecurityID is present, Tag 602.
func (m CollateralInquiryAck) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

// HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603.
func (m CollateralInquiryAck) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

// HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604.
func (m CollateralInquiryAck) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

// HasLegProduct returns true if LegProduct is present, Tag 607.
func (m CollateralInquiryAck) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

// HasLegCFICode returns true if LegCFICode is present, Tag 608.
func (m CollateralInquiryAck) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

// HasLegSecurityType returns true if LegSecurityType is present, Tag 609.
func (m CollateralInquiryAck) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

// HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610.
func (m CollateralInquiryAck) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

// HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611.
func (m CollateralInquiryAck) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

// HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612.
func (m CollateralInquiryAck) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

// HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613.
func (m CollateralInquiryAck) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

// HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614.
func (m CollateralInquiryAck) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

// HasLegCouponRate returns true if LegCouponRate is present, Tag 615.
func (m CollateralInquiryAck) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

// HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616.
func (m CollateralInquiryAck) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

// HasLegIssuer returns true if LegIssuer is present, Tag 617.
func (m CollateralInquiryAck) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

// HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618.
func (m CollateralInquiryAck) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

// HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619.
func (m CollateralInquiryAck) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

// HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620.
func (m CollateralInquiryAck) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

// HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621.
func (m CollateralInquiryAck) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

// HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622.
func (m CollateralInquiryAck) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

// HasLegRatioQty returns true if LegRatioQty is present, Tag 623.
func (m CollateralInquiryAck) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

// HasLegSide returns true if LegSide is present, Tag 624.
func (m CollateralInquiryAck) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m CollateralInquiryAck) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m CollateralInquiryAck) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasPool returns true if Pool is present, Tag 691.
func (m CollateralInquiryAck) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711.
func (m CollateralInquiryAck) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

// HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715.
func (m CollateralInquiryAck) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

// HasSettlSessID returns true if SettlSessID is present, Tag 716.
func (m CollateralInquiryAck) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

// HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717.
func (m CollateralInquiryAck) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

// HasResponseTransportType returns true if ResponseTransportType is present, Tag 725.
func (m CollateralInquiryAck) HasResponseTransportType() bool {
	return m.Has(tag.ResponseTransportType)
}

// HasResponseDestination returns true if ResponseDestination is present, Tag 726.
func (m CollateralInquiryAck) HasResponseDestination() bool {
	return m.Has(tag.ResponseDestination)
}

// HasLegDatedDate returns true if LegDatedDate is present, Tag 739.
func (m CollateralInquiryAck) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

// HasLegPool returns true if LegPool is present, Tag 740.
func (m CollateralInquiryAck) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m CollateralInquiryAck) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764.
func (m CollateralInquiryAck) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

// HasTerminationType returns true if TerminationType is present, Tag 788.
func (m CollateralInquiryAck) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

// HasQtyType returns true if QtyType is present, Tag 854.
func (m CollateralInquiryAck) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m CollateralInquiryAck) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m CollateralInquiryAck) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m CollateralInquiryAck) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m CollateralInquiryAck) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m CollateralInquiryAck) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasNoTrades returns true if NoTrades is present, Tag 897.
func (m CollateralInquiryAck) HasNoTrades() bool {
	return m.Has(tag.NoTrades)
}

// HasMarginRatio returns true if MarginRatio is present, Tag 898.
func (m CollateralInquiryAck) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

// HasCollInquiryID returns true if CollInquiryID is present, Tag 909.
func (m CollateralInquiryAck) HasCollInquiryID() bool {
	return m.Has(tag.CollInquiryID)
}

// HasTotNumReports returns true if TotNumReports is present, Tag 911.
func (m CollateralInquiryAck) HasTotNumReports() bool {
	return m.Has(tag.TotNumReports)
}

// HasAgreementDesc returns true if AgreementDesc is present, Tag 913.
func (m CollateralInquiryAck) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

// HasAgreementID returns true if AgreementID is present, Tag 914.
func (m CollateralInquiryAck) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

// HasAgreementDate returns true if AgreementDate is present, Tag 915.
func (m CollateralInquiryAck) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

// HasStartDate returns true if StartDate is present, Tag 916.
func (m CollateralInquiryAck) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

// HasEndDate returns true if EndDate is present, Tag 917.
func (m CollateralInquiryAck) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

// HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918.
func (m CollateralInquiryAck) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

// HasDeliveryType returns true if DeliveryType is present, Tag 919.
func (m CollateralInquiryAck) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

// HasNoCollInquiryQualifier returns true if NoCollInquiryQualifier is present, Tag 938.
func (m CollateralInquiryAck) HasNoCollInquiryQualifier() bool {
	return m.Has(tag.NoCollInquiryQualifier)
}

// HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942.
func (m CollateralInquiryAck) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

// HasCollInquiryStatus returns true if CollInquiryStatus is present, Tag 945.
func (m CollateralInquiryAck) HasCollInquiryStatus() bool {
	return m.Has(tag.CollInquiryStatus)
}

// HasCollInquiryResult returns true if CollInquiryResult is present, Tag 946.
func (m CollateralInquiryAck) HasCollInquiryResult() bool {
	return m.Has(tag.CollInquiryResult)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m CollateralInquiryAck) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955.
func (m CollateralInquiryAck) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

// HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956.
func (m CollateralInquiryAck) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

// HasFXCMSymID returns true if FXCMSymID is present, Tag 9000.
func (m CollateralInquiryAck) HasFXCMSymID() bool {
	return m.Has(tag.FXCMSymID)
}

// HasFXCMSymPrecision returns true if FXCMSymPrecision is present, Tag 9001.
func (m CollateralInquiryAck) HasFXCMSymPrecision() bool {
	return m.Has(tag.FXCMSymPrecision)
}

// HasFXCMSymPointSize returns true if FXCMSymPointSize is present, Tag 9002.
func (m CollateralInquiryAck) HasFXCMSymPointSize() bool {
	return m.Has(tag.FXCMSymPointSize)
}

// HasFXCMSymInterestBuy returns true if FXCMSymInterestBuy is present, Tag 9003.
func (m CollateralInquiryAck) HasFXCMSymInterestBuy() bool {
	return m.Has(tag.FXCMSymInterestBuy)
}

// HasFXCMSymInterestSell returns true if FXCMSymInterestSell is present, Tag 9004.
func (m CollateralInquiryAck) HasFXCMSymInterestSell() bool {
	return m.Has(tag.FXCMSymInterestSell)
}

// HasFXCMSymSortOrder returns true if FXCMSymSortOrder is present, Tag 9005.
func (m CollateralInquiryAck) HasFXCMSymSortOrder() bool {
	return m.Has(tag.FXCMSymSortOrder)
}

// HasFXCMSymMarginRatio returns true if FXCMSymMarginRatio is present, Tag 9006.
func (m CollateralInquiryAck) HasFXCMSymMarginRatio() bool {
	return m.Has(tag.FXCMSymMarginRatio)
}

// HasFXCMRequestRejectReason returns true if FXCMRequestRejectReason is present, Tag 9025.
func (m CollateralInquiryAck) HasFXCMRequestRejectReason() bool {
	return m.Has(tag.FXCMRequestRejectReason)
}

// HasFXCMErrorDetails returns true if FXCMErrorDetails is present, Tag 9029.
func (m CollateralInquiryAck) HasFXCMErrorDetails() bool {
	return m.Has(tag.FXCMErrorDetails)
}

// HasFXCMSubscriptionStatus returns true if FXCMSubscriptionStatus is present, Tag 9076.
func (m CollateralInquiryAck) HasFXCMSubscriptionStatus() bool {
	return m.Has(tag.FXCMSubscriptionStatus)
}

// HasFXCMProductID returns true if FXCMProductID is present, Tag 9080.
func (m CollateralInquiryAck) HasFXCMProductID() bool {
	return m.Has(tag.FXCMProductID)
}

// HasFXCMCondDistStop returns true if FXCMCondDistStop is present, Tag 9090.
func (m CollateralInquiryAck) HasFXCMCondDistStop() bool {
	return m.Has(tag.FXCMCondDistStop)
}

// HasFXCMCondDistLimit returns true if FXCMCondDistLimit is present, Tag 9091.
func (m CollateralInquiryAck) HasFXCMCondDistLimit() bool {
	return m.Has(tag.FXCMCondDistLimit)
}

// HasFXCMCondDistEntryStop returns true if FXCMCondDistEntryStop is present, Tag 9092.
func (m CollateralInquiryAck) HasFXCMCondDistEntryStop() bool {
	return m.Has(tag.FXCMCondDistEntryStop)
}

// HasFXCMCondDistEntryLimit returns true if FXCMCondDistEntryLimit is present, Tag 9093.
func (m CollateralInquiryAck) HasFXCMCondDistEntryLimit() bool {
	return m.Has(tag.FXCMCondDistEntryLimit)
}

// HasFXCMMaxQuantity returns true if FXCMMaxQuantity is present, Tag 9094.
func (m CollateralInquiryAck) HasFXCMMaxQuantity() bool {
	return m.Has(tag.FXCMMaxQuantity)
}

// HasFXCMMinQuantity returns true if FXCMMinQuantity is present, Tag 9095.
func (m CollateralInquiryAck) HasFXCMMinQuantity() bool {
	return m.Has(tag.FXCMMinQuantity)
}

// HasFXCMTradingStatus returns true if FXCMTradingStatus is present, Tag 9096.
func (m CollateralInquiryAck) HasFXCMTradingStatus() bool {
	return m.Has(tag.FXCMTradingStatus)
}

// NoExecs is a repeating group element, Tag 124.
type NoExecs struct {
	*quickfix.Group
}

// SetExecID sets ExecID, Tag 17.
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

// GetExecID gets ExecID, Tag 17.
func (m NoExecs) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasExecID returns true if ExecID is present, Tag 17.
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

// NoExecsRepeatingGroup is a repeating group, Tag 124.
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup.
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecID)})}
}

// Add create and append a new NoExecs to this group.
func (m NoExecsRepeatingGroup) Add() NoExecs {
	g := m.RepeatingGroup.Add()
	return NoExecs{g}
}

// Get returns the ith NoExecs in the NoExecsRepeatinGroup.
func (m NoExecsRepeatingGroup) Get(i int) NoExecs {
	return NoExecs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoSecurityAltID is a repeating group element, Tag 454.
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455.
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455.
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455.
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456.
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454.
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup.
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

// Add create and append a new NoSecurityAltID to this group.
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup.
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoLegSecurityAltID is a repeating group element, Tag 604.
type NoLegSecurityAltID struct {
	*quickfix.Group
}

// SetLegSecurityAltID sets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

// SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

// GetLegSecurityAltID gets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605.
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

// HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606.
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

// NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604.
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup.
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

// Add create and append a new NoLegSecurityAltID to this group.
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

// Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup.
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyings is a repeating group element, Tag 711.
type NoUnderlyings struct {
	*quickfix.Group
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

// SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

// SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

// SetUnderlyingQty sets UnderlyingQty, Tag 879.
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

// SetUnderlyingPx sets UnderlyingPx, Tag 810.
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

// SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

// SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

// SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

// SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

// SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

// SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingQty gets UnderlyingQty, Tag 879.
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPx gets UnderlyingPx, Tag 810.
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311.
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312.
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309.
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305.
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457.
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462.
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463.
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310.
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763.
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313.
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542.
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241.
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242.
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243.
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244.
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245.
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246.
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256.
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595.
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592.
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593.
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594.
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247.
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316.
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941.
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317.
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436.
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435.
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308.
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306.
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362.
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363.
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307.
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877.
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

// HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878.
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

// HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318.
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

// HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879.
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

// HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810.
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

// HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882.
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

// HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883.
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

// HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884.
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

// HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885.
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

// HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886.
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

// HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887.
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457.
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457.
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup.
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingStips is a repeating group element, Tag 887.
type NoUnderlyingStips struct {
	*quickfix.Group
}

// SetUnderlyingStipType sets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

// SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

// GetUnderlyingStipType gets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888.
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

// HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889.
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

// NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887.
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup.
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

// Add create and append a new NoUnderlyingStips to this group.
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

// Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup.
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingsRepeatingGroup is a repeating group, Tag 711.
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup.
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup()})}
}

// Add create and append a new NoUnderlyings to this group.
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

// Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup.
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

// NoEvents is a repeating group element, Tag 864.
type NoEvents struct {
	*quickfix.Group
}

// SetEventType sets EventType, Tag 865.
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

// SetEventDate sets EventDate, Tag 866.
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

// SetEventPx sets EventPx, Tag 867.
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

// SetEventText sets EventText, Tag 868.
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

// GetEventType gets EventType, Tag 865.
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventDate gets EventDate, Tag 866.
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventPx gets EventPx, Tag 867.
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventText gets EventText, Tag 868.
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEventType returns true if EventType is present, Tag 865.
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

// HasEventDate returns true if EventDate is present, Tag 866.
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

// HasEventPx returns true if EventPx is present, Tag 867.
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

// HasEventText returns true if EventText is present, Tag 868.
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864.
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup.
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText)})}
}

// Add create and append a new NoEvents to this group.
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

// Get returns the ith NoEvents in the NoEventsRepeatinGroup.
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

// NoTrades is a repeating group element, Tag 897.
type NoTrades struct {
	*quickfix.Group
}

// SetTradeReportID sets TradeReportID, Tag 571.
func (m NoTrades) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

// SetSecondaryTradeReportID sets SecondaryTradeReportID, Tag 818.
func (m NoTrades) SetSecondaryTradeReportID(v string) {
	m.Set(field.NewSecondaryTradeReportID(v))
}

// GetTradeReportID gets TradeReportID, Tag 571.
func (m NoTrades) GetTradeReportID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryTradeReportID gets SecondaryTradeReportID, Tag 818.
func (m NoTrades) GetSecondaryTradeReportID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryTradeReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTradeReportID returns true if TradeReportID is present, Tag 571.
func (m NoTrades) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

// HasSecondaryTradeReportID returns true if SecondaryTradeReportID is present, Tag 818.
func (m NoTrades) HasSecondaryTradeReportID() bool {
	return m.Has(tag.SecondaryTradeReportID)
}

// NoTradesRepeatingGroup is a repeating group, Tag 897.
type NoTradesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTradesRepeatingGroup returns an initialized, NoTradesRepeatingGroup.
func NewNoTradesRepeatingGroup() NoTradesRepeatingGroup {
	return NoTradesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrades,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradeReportID), quickfix.GroupElement(tag.SecondaryTradeReportID)})}
}

// Add create and append a new NoTrades to this group.
func (m NoTradesRepeatingGroup) Add() NoTrades {
	g := m.RepeatingGroup.Add()
	return NoTrades{g}
}

// Get returns the ith NoTrades in the NoTradesRepeatinGroup.
func (m NoTradesRepeatingGroup) Get(i int) NoTrades {
	return NoTrades{m.RepeatingGroup.Get(i)}
}

// NoCollInquiryQualifier is a repeating group element, Tag 938.
type NoCollInquiryQualifier struct {
	*quickfix.Group
}

// SetCollInquiryQualifier sets CollInquiryQualifier, Tag 896.
func (m NoCollInquiryQualifier) SetCollInquiryQualifier(v enum.CollInquiryQualifier) {
	m.Set(field.NewCollInquiryQualifier(v))
}

// GetCollInquiryQualifier gets CollInquiryQualifier, Tag 896.
func (m NoCollInquiryQualifier) GetCollInquiryQualifier() (v enum.CollInquiryQualifier, err quickfix.MessageRejectError) {
	var f field.CollInquiryQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCollInquiryQualifier returns true if CollInquiryQualifier is present, Tag 896.
func (m NoCollInquiryQualifier) HasCollInquiryQualifier() bool {
	return m.Has(tag.CollInquiryQualifier)
}

// NoCollInquiryQualifierRepeatingGroup is a repeating group, Tag 938.
type NoCollInquiryQualifierRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoCollInquiryQualifierRepeatingGroup returns an initialized, NoCollInquiryQualifierRepeatingGroup.
func NewNoCollInquiryQualifierRepeatingGroup() NoCollInquiryQualifierRepeatingGroup {
	return NoCollInquiryQualifierRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoCollInquiryQualifier,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.CollInquiryQualifier)})}
}

// Add create and append a new NoCollInquiryQualifier to this group.
func (m NoCollInquiryQualifierRepeatingGroup) Add() NoCollInquiryQualifier {
	g := m.RepeatingGroup.Add()
	return NoCollInquiryQualifier{g}
}

// Get returns the ith NoCollInquiryQualifier in the NoCollInquiryQualifierRepeatinGroup.
func (m NoCollInquiryQualifierRepeatingGroup) Get(i int) NoCollInquiryQualifier {
	return NoCollInquiryQualifier{m.RepeatingGroup.Get(i)}
}
