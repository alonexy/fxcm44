package marketdatarequestreject

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// MarketDataRequestReject is the fix44 MarketDataRequestReject type, MsgType = Y.
type MarketDataRequestReject struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataRequestReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataRequestReject {
	return MarketDataRequestReject{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataRequestReject initialized with the required fields for MarketDataRequestReject.
func New(mdreqid field.MDReqIDField) (m MarketDataRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Y"))
	m.Set(mdreqid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "Y", r
}

// SetText sets Text, Tag 58.
func (m MarketDataRequestReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetMDReqID sets MDReqID, Tag 262.
func (m MarketDataRequestReject) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

// SetMDReqRejReason sets MDReqRejReason, Tag 281.
func (m MarketDataRequestReject) SetMDReqRejReason(v enum.MDReqRejReason) {
	m.Set(field.NewMDReqRejReason(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m MarketDataRequestReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoAltMDSource sets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) SetNoAltMDSource(f NoAltMDSourceRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m MarketDataRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataRequestReject) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqRejReason gets MDReqRejReason, Tag 281.
func (m MarketDataRequestReject) GetMDReqRejReason() (v enum.MDReqRejReason, err quickfix.MessageRejectError) {
	var f field.MDReqRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m MarketDataRequestReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoAltMDSource gets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) GetNoAltMDSource() (NoAltMDSourceRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAltMDSourceRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m MarketDataRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataRequestReject) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasMDReqRejReason returns true if MDReqRejReason is present, Tag 281.
func (m MarketDataRequestReject) HasMDReqRejReason() bool {
	return m.Has(tag.MDReqRejReason)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m MarketDataRequestReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m MarketDataRequestReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoAltMDSource returns true if NoAltMDSource is present, Tag 816.
func (m MarketDataRequestReject) HasNoAltMDSource() bool {
	return m.Has(tag.NoAltMDSource)
}

// NoAltMDSource is a repeating group element, Tag 816.
type NoAltMDSource struct {
	*quickfix.Group
}

// SetAltMDSourceID sets AltMDSourceID, Tag 817.
func (m NoAltMDSource) SetAltMDSourceID(v string) {
	m.Set(field.NewAltMDSourceID(v))
}

// GetAltMDSourceID gets AltMDSourceID, Tag 817.
func (m NoAltMDSource) GetAltMDSourceID() (v string, err quickfix.MessageRejectError) {
	var f field.AltMDSourceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAltMDSourceID returns true if AltMDSourceID is present, Tag 817.
func (m NoAltMDSource) HasAltMDSourceID() bool {
	return m.Has(tag.AltMDSourceID)
}

// NoAltMDSourceRepeatingGroup is a repeating group, Tag 816.
type NoAltMDSourceRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAltMDSourceRepeatingGroup returns an initialized, NoAltMDSourceRepeatingGroup.
func NewNoAltMDSourceRepeatingGroup() NoAltMDSourceRepeatingGroup {
	return NoAltMDSourceRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAltMDSource,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AltMDSourceID)})}
}

// Add create and append a new NoAltMDSource to this group.
func (m NoAltMDSourceRepeatingGroup) Add() NoAltMDSource {
	g := m.RepeatingGroup.Add()
	return NoAltMDSource{g}
}

// Get returns the ith NoAltMDSource in the NoAltMDSourceRepeatinGroup.
func (m NoAltMDSourceRepeatingGroup) Get(i int) NoAltMDSource {
	return NoAltMDSource{m.RepeatingGroup.Get(i)}
}
edSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m MarketDataRequestReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m MarketDataRequestReject) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m MarketDataRequestReject) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m MarketDataRequestReject) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m MarketDataRequestReject) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m MarketDataRequestReject) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m MarketDataRequestReject) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m MarketDataRequestReject) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m MarketDataRequestReject) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetRoundLot sets RoundLot, Tag 561.
func (m MarketDataRequestReject) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m MarketDataRequestReject) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m MarketDataRequestReject) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetPool sets Pool, Tag 691.
func (m MarketDataRequestReject) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m MarketDataRequestReject) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetNoAltMDSource sets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) SetNoAltMDSource(f NoAltMDSourceRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoEvents sets NoEvents, Tag 864.
func (m MarketDataRequestReject) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m MarketDataRequestReject) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m MarketDataRequestReject) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m MarketDataRequestReject) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m MarketDataRequestReject) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m MarketDataRequestReject) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetFXCMSymID sets FXCMSymID, Tag 9000.
func (m MarketDataRequestReject) SetFXCMSymID(v string) {
	m.Set(field.NewFXCMSymID(v))
}

// SetFXCMSymPrecision sets FXCMSymPrecision, Tag 9001.
func (m MarketDataRequestReject) SetFXCMSymPrecision(v int) {
	m.Set(field.NewFXCMSymPrecision(v))
}

// SetFXCMSymPointSize sets FXCMSymPointSize, Tag 9002.
func (m MarketDataRequestReject) SetFXCMSymPointSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymPointSize(value, scale))
}

// SetFXCMSymInterestBuy sets FXCMSymInterestBuy, Tag 9003.
func (m MarketDataRequestReject) SetFXCMSymInterestBuy(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestBuy(value, scale))
}

// SetFXCMSymInterestSell sets FXCMSymInterestSell, Tag 9004.
func (m MarketDataRequestReject) SetFXCMSymInterestSell(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestSell(value, scale))
}

// SetFXCMSymSortOrder sets FXCMSymSortOrder, Tag 9005.
func (m MarketDataRequestReject) SetFXCMSymSortOrder(v int) {
	m.Set(field.NewFXCMSymSortOrder(v))
}

// SetFXCMSymMarginRatio sets FXCMSymMarginRatio, Tag 9006.
func (m MarketDataRequestReject) SetFXCMSymMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymMarginRatio(value, scale))
}

// SetFXCMTimingInterval sets FXCMTimingInterval, Tag 9011.
func (m MarketDataRequestReject) SetFXCMTimingInterval(v enum.FXCMTimingInterval) {
	m.Set(field.NewFXCMTimingInterval(v))
}

// SetFXCMRequestRejectReason sets FXCMRequestRejectReason, Tag 9025.
func (m MarketDataRequestReject) SetFXCMRequestRejectReason(v enum.FXCMRequestRejectReason) {
	m.Set(field.NewFXCMRequestRejectReason(v))
}

// SetFXCMErrorDetails sets FXCMErrorDetails, Tag 9029.
func (m MarketDataRequestReject) SetFXCMErrorDetails(v string) {
	m.Set(field.NewFXCMErrorDetails(v))
}

// SetFXCMSubscriptionStatus sets FXCMSubscriptionStatus, Tag 9076.
func (m MarketDataRequestReject) SetFXCMSubscriptionStatus(v string) {
	m.Set(field.NewFXCMSubscriptionStatus(v))
}

// SetFXCMProductID sets FXCMProductID, Tag 9080.
func (m MarketDataRequestReject) SetFXCMProductID(v int) {
	m.Set(field.NewFXCMProductID(v))
}

// SetFXCMCondDistStop sets FXCMCondDistStop, Tag 9090.
func (m MarketDataRequestReject) SetFXCMCondDistStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistStop(value, scale))
}

// SetFXCMCondDistLimit sets FXCMCondDistLimit, Tag 9091.
func (m MarketDataRequestReject) SetFXCMCondDistLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistLimit(value, scale))
}

// SetFXCMCondDistEntryStop sets FXCMCondDistEntryStop, Tag 9092.
func (m MarketDataRequestReject) SetFXCMCondDistEntryStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryStop(value, scale))
}

// SetFXCMCondDistEntryLimit sets FXCMCondDistEntryLimit, Tag 9093.
func (m MarketDataRequestReject) SetFXCMCondDistEntryLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryLimit(value, scale))
}

// SetFXCMMaxQuantity sets FXCMMaxQuantity, Tag 9094.
func (m MarketDataRequestReject) SetFXCMMaxQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMaxQuantity(value, scale))
}

// SetFXCMMinQuantity sets FXCMMinQuantity, Tag 9095.
func (m MarketDataRequestReject) SetFXCMMinQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMinQuantity(value, scale))
}

// SetFXCMTradingStatus sets FXCMTradingStatus, Tag 9096.
func (m MarketDataRequestReject) SetFXCMTradingStatus(v string) {
	m.Set(field.NewFXCMTradingStatus(v))
}

// GetCurrency gets Currency, Tag 15.
func (m MarketDataRequestReject) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m MarketDataRequestReject) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m MarketDataRequestReject) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m MarketDataRequestReject) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m MarketDataRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m MarketDataRequestReject) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m MarketDataRequestReject) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m MarketDataRequestReject) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m MarketDataRequestReject) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m MarketDataRequestReject) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m MarketDataRequestReject) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m MarketDataRequestReject) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m MarketDataRequestReject) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m MarketDataRequestReject) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m MarketDataRequestReject) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m MarketDataRequestReject) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m MarketDataRequestReject) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m MarketDataRequestReject) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m MarketDataRequestReject) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m MarketDataRequestReject) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m MarketDataRequestReject) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m MarketDataRequestReject) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m MarketDataRequestReject) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataRequestReject) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqRejReason gets MDReqRejReason, Tag 281.
func (m MarketDataRequestReject) GetMDReqRejReason() (v enum.MDReqRejReason, err quickfix.MessageRejectError) {
	var f field.MDReqRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m MarketDataRequestReject) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m MarketDataRequestReject) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m MarketDataRequestReject) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m MarketDataRequestReject) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m MarketDataRequestReject) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m MarketDataRequestReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m MarketDataRequestReject) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m MarketDataRequestReject) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m MarketDataRequestReject) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m MarketDataRequestReject) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m MarketDataRequestReject) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m MarketDataRequestReject) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m MarketDataRequestReject) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m MarketDataRequestReject) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoundLot gets RoundLot, Tag 561.
func (m MarketDataRequestReject) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m MarketDataRequestReject) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m MarketDataRequestReject) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m MarketDataRequestReject) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m MarketDataRequestReject) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoAltMDSource gets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) GetNoAltMDSource() (NoAltMDSourceRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAltMDSourceRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoEvents gets NoEvents, Tag 864.
func (m MarketDataRequestReject) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m MarketDataRequestReject) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m MarketDataRequestReject) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m MarketDataRequestReject) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m MarketDataRequestReject) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m MarketDataRequestReject) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymID gets FXCMSymID, Tag 9000.
func (m MarketDataRequestReject) GetFXCMSymID() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSymIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPrecision gets FXCMSymPrecision, Tag 9001.
func (m MarketDataRequestReject) GetFXCMSymPrecision() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPointSize gets FXCMSymPointSize, Tag 9002.
func (m MarketDataRequestReject) GetFXCMSymPointSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymPointSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestBuy gets FXCMSymInterestBuy, Tag 9003.
func (m MarketDataRequestReject) GetFXCMSymInterestBuy() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestBuyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestSell gets FXCMSymInterestSell, Tag 9004.
func (m MarketDataRequestReject) GetFXCMSymInterestSell() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestSellField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymSortOrder gets FXCMSymSortOrder, Tag 9005.
func (m MarketDataRequestReject) GetFXCMSymSortOrder() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymSortOrderField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymMarginRatio gets FXCMSymMarginRatio, Tag 9006.
func (m MarketDataRequestReject) GetFXCMSymMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymMarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMTimingInterval gets FXCMTimingInterval, Tag 9011.
func (m MarketDataRequestReject) GetFXCMTimingInterval() (v enum.FXCMTimingInterval, err quickfix.MessageRejectError) {
	var f field.FXCMTimingIntervalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMRequestRejectReason gets FXCMRequestRejectReason, Tag 9025.
func (m MarketDataRequestReject) GetFXCMRequestRejectReason() (v enum.FXCMRequestRejectReason, err quickfix.MessageRejectError) {
	var f field.FXCMRequestRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMErrorDetails gets FXCMErrorDetails, Tag 9029.
func (m MarketDataRequestReject) GetFXCMErrorDetails() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMErrorDetailsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSubscriptionStatus gets FXCMSubscriptionStatus, Tag 9076.
func (m MarketDataRequestReject) GetFXCMSubscriptionStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSubscriptionStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMProductID gets FXCMProductID, Tag 9080.
func (m MarketDataRequestReject) GetFXCMProductID() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMProductIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistStop gets FXCMCondDistStop, Tag 9090.
func (m MarketDataRequestReject) GetFXCMCondDistStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistLimit gets FXCMCondDistLimit, Tag 9091.
func (m MarketDataRequestReject) GetFXCMCondDistLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryStop gets FXCMCondDistEntryStop, Tag 9092.
func (m MarketDataRequestReject) GetFXCMCondDistEntryStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryLimit gets FXCMCondDistEntryLimit, Tag 9093.
func (m MarketDataRequestReject) GetFXCMCondDistEntryLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMaxQuantity gets FXCMMaxQuantity, Tag 9094.
func (m MarketDataRequestReject) GetFXCMMaxQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMaxQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMinQuantity gets FXCMMinQuantity, Tag 9095.
func (m MarketDataRequestReject) GetFXCMMinQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMinQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMTradingStatus gets FXCMTradingStatus, Tag 9096.
func (m MarketDataRequestReject) GetFXCMTradingStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m MarketDataRequestReject) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m MarketDataRequestReject) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m MarketDataRequestReject) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m MarketDataRequestReject) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m MarketDataRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m MarketDataRequestReject) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m MarketDataRequestReject) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m MarketDataRequestReject) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m MarketDataRequestReject) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m MarketDataRequestReject) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m MarketDataRequestReject) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m MarketDataRequestReject) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m MarketDataRequestReject) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m MarketDataRequestReject) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m MarketDataRequestReject) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m MarketDataRequestReject) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m MarketDataRequestReject) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m MarketDataRequestReject) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m MarketDataRequestReject) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m MarketDataRequestReject) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m MarketDataRequestReject) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m MarketDataRequestReject) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m MarketDataRequestReject) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataRequestReject) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasMDReqRejReason returns true if MDReqRejReason is present, Tag 281.
func (m MarketDataRequestReject) HasMDReqRejReason() bool {
	return m.Has(tag.MDReqRejReason)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m MarketDataRequestReject) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m MarketDataRequestReject) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m MarketDataRequestReject) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m MarketDataRequestReject) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m MarketDataRequestReject) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m MarketDataRequestReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m MarketDataRequestReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m MarketDataRequestReject) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m MarketDataRequestReject) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m MarketDataRequestReject) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m MarketDataRequestReject) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m MarketDataRequestReject) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m MarketDataRequestReject) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m MarketDataRequestReject) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m MarketDataRequestReject) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasRoundLot returns true if RoundLot is present, Tag 561.
func (m MarketDataRequestReject) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m MarketDataRequestReject) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m MarketDataRequestReject) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasPool returns true if Pool is present, Tag 691.
func (m MarketDataRequestReject) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m MarketDataRequestReject) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasNoAltMDSource returns true if NoAltMDSource is present, Tag 816.
func (m MarketDataRequestReject) HasNoAltMDSource() bool {
	return m.Has(tag.NoAltMDSource)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m MarketDataRequestReject) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m MarketDataRequestReject) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m MarketDataRequestReject) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m MarketDataRequestReject) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m MarketDataRequestReject) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m MarketDataRequestReject) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasFXCMSymID returns true if FXCMSymID is present, Tag 9000.
func (m MarketDataRequestReject) HasFXCMSymID() bool {
	return m.Has(tag.FXCMSymID)
}

// HasFXCMSymPrecision returns true if FXCMSymPrecision is present, Tag 9001.
func (m MarketDataRequestReject) HasFXCMSymPrecision() bool {
	return m.Has(tag.FXCMSymPrecision)
}

// HasFXCMSymPointSize returns true if FXCMSymPointSize is present, Tag 9002.
func (m MarketDataRequestReject) HasFXCMSymPointSize() bool {
	return m.Has(tag.FXCMSymPointSize)
}

// HasFXCMSymInterestBuy returns true if FXCMSymInterestBuy is present, Tag 9003.
func (m MarketDataRequestReject) HasFXCMSymInterestBuy() bool {
	return m.Has(tag.FXCMSymInterestBuy)
}

// HasFXCMSymInterestSell returns true if FXCMSymInterestSell is present, Tag 9004.
func (m MarketDataRequestReject) HasFXCMSymInterestSell() bool {
	return m.Has(tag.FXCMSymInterestSell)
}

// HasFXCMSymSortOrder returns true if FXCMSymSortOrder is present, Tag 9005.
func (m MarketDataRequestReject) HasFXCMSymSortOrder() bool {
	return m.Has(tag.FXCMSymSortOrder)
}

// HasFXCMSymMarginRatio returns true if FXCMSymMarginRatio is present, Tag 9006.
func (m MarketDataRequestReject) HasFXCMSymMarginRatio() bool {
	return m.Has(tag.FXCMSymMarginRatio)
}

// HasFXCMTimingInterval returns true if FXCMTimingInterval is present, Tag 9011.
func (m MarketDataRequestReject) HasFXCMTimingInterval() bool {
	return m.Has(tag.FXCMTimingInterval)
}

// HasFXCMRequestRejectReason returns true if FXCMRequestRejectReason is present, Tag 9025.
func (m MarketDataRequestReject) HasFXCMRequestRejectReason() bool {
	return m.Has(tag.FXCMRequestRejectReason)
}

// HasFXCMErrorDetails returns true if FXCMErrorDetails is present, Tag 9029.
func (m MarketDataRequestReject) HasFXCMErrorDetails() bool {
	return m.Has(tag.FXCMErrorDetails)
}

// HasFXCMSubscriptionStatus returns true if FXCMSubscriptionStatus is present, Tag 9076.
func (m MarketDataRequestReject) HasFXCMSubscriptionStatus() bool {
	return m.Has(tag.FXCMSubscriptionStatus)
}

// HasFXCMProductID returns true if FXCMProductID is present, Tag 9080.
func (m MarketDataRequestReject) HasFXCMProductID() bool {
	return m.Has(tag.FXCMProductID)
}

// HasFXCMCondDistStop returns true if FXCMCondDistStop is present, Tag 9090.
func (m MarketDataRequestReject) HasFXCMCondDistStop() bool {
	return m.Has(tag.FXCMCondDistStop)
}

// HasFXCMCondDistLimit returns true if FXCMCondDistLimit is present, Tag 9091.
func (m MarketDataRequestReject) HasFXCMCondDistLimit() bool {
	return m.Has(tag.FXCMCondDistLimit)
}

// HasFXCMCondDistEntryStop returns true if FXCMCondDistEntryStop is present, Tag 9092.
func (m MarketDataRequestReject) HasFXCMCondDistEntryStop() bool {
	return m.Has(tag.FXCMCondDistEntryStop)
}

// HasFXCMCondDistEntryLimit returns true if FXCMCondDistEntryLimit is present, Tag 9093.
func (m MarketDataRequestReject) HasFXCMCondDistEntryLimit() bool {
	return m.Has(tag.FXCMCondDistEntryLimit)
}

// HasFXCMMaxQuantity returns true if FXCMMaxQuantity is present, Tag 9094.
func (m MarketDataRequestReject) HasFXCMMaxQuantity() bool {
	return m.Has(tag.FXCMMaxQuantity)
}

// HasFXCMMinQuantity returns true if FXCMMinQuantity is present, Tag 9095.
func (m MarketDataRequestReject) HasFXCMMinQuantity() bool {
	return m.Has(tag.FXCMMinQuantity)
}

// HasFXCMTradingStatus returns true if FXCMTradingStatus is present, Tag 9096.
func (m MarketDataRequestReject) HasFXCMTradingStatus() bool {
	return m.Has(tag.FXCMTradingStatus)
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

// NoAltMDSource is a repeating group element, Tag 816.
type NoAltMDSource struct {
	*quickfix.Group
}

// SetAltMDSourceID sets AltMDSourceID, Tag 817.
func (m NoAltMDSource) SetAltMDSourceID(v string) {
	m.Set(field.NewAltMDSourceID(v))
}

// GetAltMDSourceID gets AltMDSourceID, Tag 817.
func (m NoAltMDSource) GetAltMDSourceID() (v string, err quickfix.MessageRejectError) {
	var f field.AltMDSourceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAltMDSourceID returns true if AltMDSourceID is present, Tag 817.
func (m NoAltMDSource) HasAltMDSourceID() bool {
	return m.Has(tag.AltMDSourceID)
}

// NoAltMDSourceRepeatingGroup is a repeating group, Tag 816.
type NoAltMDSourceRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAltMDSourceRepeatingGroup returns an initialized, NoAltMDSourceRepeatingGroup.
func NewNoAltMDSourceRepeatingGroup() NoAltMDSourceRepeatingGroup {
	return NoAltMDSourceRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAltMDSource,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AltMDSourceID)})}
}

// Add create and append a new NoAltMDSource to this group.
func (m NoAltMDSourceRepeatingGroup) Add() NoAltMDSource {
	g := m.RepeatingGroup.Add()
	return NoAltMDSource{g}
}

// Get returns the ith NoAltMDSource in the NoAltMDSourceRepeatinGroup.
func (m NoAltMDSourceRepeatingGroup) Get(i int) NoAltMDSource {
	return NoAltMDSource{m.RepeatingGroup.Get(i)}
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
