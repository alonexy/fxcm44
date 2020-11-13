package tradingsessionstatus

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// TradingSessionStatus is the fix44 TradingSessionStatus type, MsgType = h.
type TradingSessionStatus struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TradingSessionStatus from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) TradingSessionStatus {
	return TradingSessionStatus{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m TradingSessionStatus) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a TradingSessionStatus initialized with the required fields for TradingSessionStatus.
func New(tradingsessionid field.TradingSessionIDField, tradsesstatus field.TradSesStatusField) (m TradingSessionStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("h"))
	m.Set(tradingsessionid)
	m.Set(tradsesstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "h", r
}

// SetText sets Text, Tag 58.
func (m TradingSessionStatus) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m TradingSessionStatus) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m TradingSessionStatus) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325.
func (m TradingSessionStatus) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

// SetTradSesReqID sets TradSesReqID, Tag 335.
func (m TradingSessionStatus) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m TradingSessionStatus) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradSesMethod sets TradSesMethod, Tag 338.
func (m TradingSessionStatus) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

// SetTradSesMode sets TradSesMode, Tag 339.
func (m TradingSessionStatus) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

// SetTradSesStatus sets TradSesStatus, Tag 340.
func (m TradingSessionStatus) SetTradSesStatus(v enum.TradSesStatus) {
	m.Set(field.NewTradSesStatus(v))
}

// SetTradSesStartTime sets TradSesStartTime, Tag 341.
func (m TradingSessionStatus) SetTradSesStartTime(v time.Time) {
	m.Set(field.NewTradSesStartTime(v))
}

// SetTradSesOpenTime sets TradSesOpenTime, Tag 342.
func (m TradingSessionStatus) SetTradSesOpenTime(v time.Time) {
	m.Set(field.NewTradSesOpenTime(v))
}

// SetTradSesPreCloseTime sets TradSesPreCloseTime, Tag 343.
func (m TradingSessionStatus) SetTradSesPreCloseTime(v time.Time) {
	m.Set(field.NewTradSesPreCloseTime(v))
}

// SetTradSesCloseTime sets TradSesCloseTime, Tag 344.
func (m TradingSessionStatus) SetTradSesCloseTime(v time.Time) {
	m.Set(field.NewTradSesCloseTime(v))
}

// SetTradSesEndTime sets TradSesEndTime, Tag 345.
func (m TradingSessionStatus) SetTradSesEndTime(v time.Time) {
	m.Set(field.NewTradSesEndTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m TradingSessionStatus) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m TradingSessionStatus) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387.
func (m TradingSessionStatus) SetTotalVolumeTraded(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalVolumeTraded(value, scale))
}

// SetTradSesStatusRejReason sets TradSesStatusRejReason, Tag 567.
func (m TradingSessionStatus) SetTradSesStatusRejReason(v enum.TradSesStatusRejReason) {
	m.Set(field.NewTradSesStatusRejReason(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m TradingSessionStatus) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetFXCMNoParam sets FXCMNoParam, Tag 9016.
func (m TradingSessionStatus) SetFXCMNoParam(f FXCMNoParamRepeatingGroup) {
	m.SetGroup(f)
}

// SetFXCMServerTimeZone sets FXCMServerTimeZone, Tag 9019.
func (m TradingSessionStatus) SetFXCMServerTimeZone(v string) {
	m.Set(field.NewFXCMServerTimeZone(v))
}

// SetFXCMServerTimeZoneName sets FXCMServerTimeZoneName, Tag 9030.
func (m TradingSessionStatus) SetFXCMServerTimeZoneName(v string) {
	m.Set(field.NewFXCMServerTimeZoneName(v))
}

// GetText gets Text, Tag 58.
func (m TradingSessionStatus) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m TradingSessionStatus) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m TradingSessionStatus) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325.
func (m TradingSessionStatus) GetUnsolicitedIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.UnsolicitedIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesReqID gets TradSesReqID, Tag 335.
func (m TradingSessionStatus) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m TradingSessionStatus) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMethod gets TradSesMethod, Tag 338.
func (m TradingSessionStatus) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMode gets TradSesMode, Tag 339.
func (m TradingSessionStatus) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStatus gets TradSesStatus, Tag 340.
func (m TradingSessionStatus) GetTradSesStatus() (v enum.TradSesStatus, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStartTime gets TradSesStartTime, Tag 341.
func (m TradingSessionStatus) GetTradSesStartTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesOpenTime gets TradSesOpenTime, Tag 342.
func (m TradingSessionStatus) GetTradSesOpenTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesOpenTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343.
func (m TradingSessionStatus) GetTradSesPreCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesPreCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesCloseTime gets TradSesCloseTime, Tag 344.
func (m TradingSessionStatus) GetTradSesCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesEndTime gets TradSesEndTime, Tag 345.
func (m TradingSessionStatus) GetTradSesEndTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m TradingSessionStatus) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m TradingSessionStatus) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387.
func (m TradingSessionStatus) GetTotalVolumeTraded() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesStatusRejReason gets TradSesStatusRejReason, Tag 567.
func (m TradingSessionStatus) GetTradSesStatusRejReason() (v enum.TradSesStatusRejReason, err quickfix.MessageRejectError) {
	var f field.TradSesStatusRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m TradingSessionStatus) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMNoParam gets FXCMNoParam, Tag 9016.
func (m TradingSessionStatus) GetFXCMNoParam() (FXCMNoParamRepeatingGroup, quickfix.MessageRejectError) {
	f := NewFXCMNoParamRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetFXCMServerTimeZone gets FXCMServerTimeZone, Tag 9019.
func (m TradingSessionStatus) GetFXCMServerTimeZone() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMServerTimeZoneField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMServerTimeZoneName gets FXCMServerTimeZoneName, Tag 9030.
func (m TradingSessionStatus) GetFXCMServerTimeZoneName() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMServerTimeZoneNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m TradingSessionStatus) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m TradingSessionStatus) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m TradingSessionStatus) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325.
func (m TradingSessionStatus) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

// HasTradSesReqID returns true if TradSesReqID is present, Tag 335.
func (m TradingSessionStatus) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m TradingSessionStatus) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradSesMethod returns true if TradSesMethod is present, Tag 338.
func (m TradingSessionStatus) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

// HasTradSesMode returns true if TradSesMode is present, Tag 339.
func (m TradingSessionStatus) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

// HasTradSesStatus returns true if TradSesStatus is present, Tag 340.
func (m TradingSessionStatus) HasTradSesStatus() bool {
	return m.Has(tag.TradSesStatus)
}

// HasTradSesStartTime returns true if TradSesStartTime is present, Tag 341.
func (m TradingSessionStatus) HasTradSesStartTime() bool {
	return m.Has(tag.TradSesStartTime)
}

// HasTradSesOpenTime returns true if TradSesOpenTime is present, Tag 342.
func (m TradingSessionStatus) HasTradSesOpenTime() bool {
	return m.Has(tag.TradSesOpenTime)
}

// HasTradSesPreCloseTime returns true if TradSesPreCloseTime is present, Tag 343.
func (m TradingSessionStatus) HasTradSesPreCloseTime() bool {
	return m.Has(tag.TradSesPreCloseTime)
}

// HasTradSesCloseTime returns true if TradSesCloseTime is present, Tag 344.
func (m TradingSessionStatus) HasTradSesCloseTime() bool {
	return m.Has(tag.TradSesCloseTime)
}

// HasTradSesEndTime returns true if TradSesEndTime is present, Tag 345.
func (m TradingSessionStatus) HasTradSesEndTime() bool {
	return m.Has(tag.TradSesEndTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m TradingSessionStatus) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m TradingSessionStatus) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387.
func (m TradingSessionStatus) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
}

// HasTradSesStatusRejReason returns true if TradSesStatusRejReason is present, Tag 567.
func (m TradingSessionStatus) HasTradSesStatusRejReason() bool {
	return m.Has(tag.TradSesStatusRejReason)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m TradingSessionStatus) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasFXCMNoParam returns true if FXCMNoParam is present, Tag 9016.
func (m TradingSessionStatus) HasFXCMNoParam() bool {
	return m.Has(tag.FXCMNoParam)
}

// HasFXCMServerTimeZone returns true if FXCMServerTimeZone is present, Tag 9019.
func (m TradingSessionStatus) HasFXCMServerTimeZone() bool {
	return m.Has(tag.FXCMServerTimeZone)
}

// HasFXCMServerTimeZoneName returns true if FXCMServerTimeZoneName is present, Tag 9030.
func (m TradingSessionStatus) HasFXCMServerTimeZoneName() bool {
	return m.Has(tag.FXCMServerTimeZoneName)
}

// NoRelatedSym is a repeating group element, Tag 146.
type NoRelatedSym struct {
	*quickfix.Group
}

// SetSymbol sets Symbol, Tag 55.
func (m NoRelatedSym) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m NoRelatedSym) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoRelatedSym) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m NoRelatedSym) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m NoRelatedSym) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m NoRelatedSym) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m NoRelatedSym) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m NoRelatedSym) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m NoRelatedSym) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m NoRelatedSym) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m NoRelatedSym) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m NoRelatedSym) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m NoRelatedSym) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m NoRelatedSym) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m NoRelatedSym) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m NoRelatedSym) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m NoRelatedSym) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m NoRelatedSym) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m NoRelatedSym) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m NoRelatedSym) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m NoRelatedSym) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m NoRelatedSym) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m NoRelatedSym) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m NoRelatedSym) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m NoRelatedSym) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m NoRelatedSym) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m NoRelatedSym) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m NoRelatedSym) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m NoRelatedSym) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetPool sets Pool, Tag 691.
func (m NoRelatedSym) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m NoRelatedSym) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m NoRelatedSym) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m NoRelatedSym) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetNoEvents sets NoEvents, Tag 864.
func (m NoRelatedSym) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m NoRelatedSym) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m NoRelatedSym) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetFXCMSymID sets FXCMSymID, Tag 9000.
func (m NoRelatedSym) SetFXCMSymID(v string) {
	m.Set(field.NewFXCMSymID(v))
}

// SetFXCMSymPrecision sets FXCMSymPrecision, Tag 9001.
func (m NoRelatedSym) SetFXCMSymPrecision(v int) {
	m.Set(field.NewFXCMSymPrecision(v))
}

// SetFXCMSymPointSize sets FXCMSymPointSize, Tag 9002.
func (m NoRelatedSym) SetFXCMSymPointSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymPointSize(value, scale))
}

// SetFXCMSymSortOrder sets FXCMSymSortOrder, Tag 9005.
func (m NoRelatedSym) SetFXCMSymSortOrder(v int) {
	m.Set(field.NewFXCMSymSortOrder(v))
}

// SetFXCMProductID sets FXCMProductID, Tag 9080.
func (m NoRelatedSym) SetFXCMProductID(v int) {
	m.Set(field.NewFXCMProductID(v))
}

// SetFXCMSymMarginRatio sets FXCMSymMarginRatio, Tag 9006.
func (m NoRelatedSym) SetFXCMSymMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymMarginRatio(value, scale))
}

// SetCurrency sets Currency, Tag 15.
func (m NoRelatedSym) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetRoundLot sets RoundLot, Tag 561.
func (m NoRelatedSym) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

// SetFXCMSymInterestBuy sets FXCMSymInterestBuy, Tag 9003.
func (m NoRelatedSym) SetFXCMSymInterestBuy(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestBuy(value, scale))
}

// SetFXCMSymInterestSell sets FXCMSymInterestSell, Tag 9004.
func (m NoRelatedSym) SetFXCMSymInterestSell(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMSymInterestSell(value, scale))
}

// SetFXCMSubscriptionStatus sets FXCMSubscriptionStatus, Tag 9076.
func (m NoRelatedSym) SetFXCMSubscriptionStatus(v string) {
	m.Set(field.NewFXCMSubscriptionStatus(v))
}

// SetFXCMCondDistStop sets FXCMCondDistStop, Tag 9090.
func (m NoRelatedSym) SetFXCMCondDistStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistStop(value, scale))
}

// SetFXCMCondDistLimit sets FXCMCondDistLimit, Tag 9091.
func (m NoRelatedSym) SetFXCMCondDistLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistLimit(value, scale))
}

// SetFXCMCondDistEntryStop sets FXCMCondDistEntryStop, Tag 9092.
func (m NoRelatedSym) SetFXCMCondDistEntryStop(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryStop(value, scale))
}

// SetFXCMCondDistEntryLimit sets FXCMCondDistEntryLimit, Tag 9093.
func (m NoRelatedSym) SetFXCMCondDistEntryLimit(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMCondDistEntryLimit(value, scale))
}

// SetFXCMMaxQuantity sets FXCMMaxQuantity, Tag 9094.
func (m NoRelatedSym) SetFXCMMaxQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMaxQuantity(value, scale))
}

// SetFXCMMinQuantity sets FXCMMinQuantity, Tag 9095.
func (m NoRelatedSym) SetFXCMMinQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewFXCMMinQuantity(value, scale))
}

// SetFXCMTradingStatus sets FXCMTradingStatus, Tag 9096.
func (m NoRelatedSym) SetFXCMTradingStatus(v string) {
	m.Set(field.NewFXCMTradingStatus(v))
}

// GetSymbol gets Symbol, Tag 55.
func (m NoRelatedSym) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m NoRelatedSym) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoRelatedSym) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoRelatedSym) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m NoRelatedSym) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m NoRelatedSym) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m NoRelatedSym) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m NoRelatedSym) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m NoRelatedSym) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m NoRelatedSym) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m NoRelatedSym) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m NoRelatedSym) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m NoRelatedSym) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m NoRelatedSym) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m NoRelatedSym) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m NoRelatedSym) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m NoRelatedSym) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m NoRelatedSym) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m NoRelatedSym) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m NoRelatedSym) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m NoRelatedSym) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m NoRelatedSym) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m NoRelatedSym) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m NoRelatedSym) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m NoRelatedSym) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m NoRelatedSym) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m NoRelatedSym) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m NoRelatedSym) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m NoRelatedSym) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m NoRelatedSym) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m NoRelatedSym) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m NoRelatedSym) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m NoRelatedSym) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m NoRelatedSym) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m NoRelatedSym) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m NoRelatedSym) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m NoRelatedSym) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m NoRelatedSym) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m NoRelatedSym) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864.
func (m NoRelatedSym) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m NoRelatedSym) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m NoRelatedSym) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymID gets FXCMSymID, Tag 9000.
func (m NoRelatedSym) GetFXCMSymID() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSymIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPrecision gets FXCMSymPrecision, Tag 9001.
func (m NoRelatedSym) GetFXCMSymPrecision() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymPointSize gets FXCMSymPointSize, Tag 9002.
func (m NoRelatedSym) GetFXCMSymPointSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymPointSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymSortOrder gets FXCMSymSortOrder, Tag 9005.
func (m NoRelatedSym) GetFXCMSymSortOrder() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMSymSortOrderField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMProductID gets FXCMProductID, Tag 9080.
func (m NoRelatedSym) GetFXCMProductID() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMProductIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymMarginRatio gets FXCMSymMarginRatio, Tag 9006.
func (m NoRelatedSym) GetFXCMSymMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymMarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m NoRelatedSym) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoundLot gets RoundLot, Tag 561.
func (m NoRelatedSym) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestBuy gets FXCMSymInterestBuy, Tag 9003.
func (m NoRelatedSym) GetFXCMSymInterestBuy() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestBuyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSymInterestSell gets FXCMSymInterestSell, Tag 9004.
func (m NoRelatedSym) GetFXCMSymInterestSell() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMSymInterestSellField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMSubscriptionStatus gets FXCMSubscriptionStatus, Tag 9076.
func (m NoRelatedSym) GetFXCMSubscriptionStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMSubscriptionStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistStop gets FXCMCondDistStop, Tag 9090.
func (m NoRelatedSym) GetFXCMCondDistStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistLimit gets FXCMCondDistLimit, Tag 9091.
func (m NoRelatedSym) GetFXCMCondDistLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryStop gets FXCMCondDistEntryStop, Tag 9092.
func (m NoRelatedSym) GetFXCMCondDistEntryStop() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryStopField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMCondDistEntryLimit gets FXCMCondDistEntryLimit, Tag 9093.
func (m NoRelatedSym) GetFXCMCondDistEntryLimit() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMCondDistEntryLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMaxQuantity gets FXCMMaxQuantity, Tag 9094.
func (m NoRelatedSym) GetFXCMMaxQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMaxQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMMinQuantity gets FXCMMinQuantity, Tag 9095.
func (m NoRelatedSym) GetFXCMMinQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FXCMMinQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMTradingStatus gets FXCMTradingStatus, Tag 9096.
func (m NoRelatedSym) GetFXCMTradingStatus() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoRelatedSym) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m NoRelatedSym) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoRelatedSym) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m NoRelatedSym) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m NoRelatedSym) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m NoRelatedSym) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m NoRelatedSym) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m NoRelatedSym) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m NoRelatedSym) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m NoRelatedSym) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m NoRelatedSym) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m NoRelatedSym) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m NoRelatedSym) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m NoRelatedSym) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m NoRelatedSym) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m NoRelatedSym) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m NoRelatedSym) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m NoRelatedSym) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m NoRelatedSym) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m NoRelatedSym) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m NoRelatedSym) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m NoRelatedSym) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m NoRelatedSym) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m NoRelatedSym) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m NoRelatedSym) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m NoRelatedSym) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m NoRelatedSym) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m NoRelatedSym) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m NoRelatedSym) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasPool returns true if Pool is present, Tag 691.
func (m NoRelatedSym) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m NoRelatedSym) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m NoRelatedSym) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m NoRelatedSym) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m NoRelatedSym) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m NoRelatedSym) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m NoRelatedSym) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasFXCMSymID returns true if FXCMSymID is present, Tag 9000.
func (m NoRelatedSym) HasFXCMSymID() bool {
	return m.Has(tag.FXCMSymID)
}

// HasFXCMSymPrecision returns true if FXCMSymPrecision is present, Tag 9001.
func (m NoRelatedSym) HasFXCMSymPrecision() bool {
	return m.Has(tag.FXCMSymPrecision)
}

// HasFXCMSymPointSize returns true if FXCMSymPointSize is present, Tag 9002.
func (m NoRelatedSym) HasFXCMSymPointSize() bool {
	return m.Has(tag.FXCMSymPointSize)
}

// HasFXCMSymSortOrder returns true if FXCMSymSortOrder is present, Tag 9005.
func (m NoRelatedSym) HasFXCMSymSortOrder() bool {
	return m.Has(tag.FXCMSymSortOrder)
}

// HasFXCMProductID returns true if FXCMProductID is present, Tag 9080.
func (m NoRelatedSym) HasFXCMProductID() bool {
	return m.Has(tag.FXCMProductID)
}

// HasFXCMSymMarginRatio returns true if FXCMSymMarginRatio is present, Tag 9006.
func (m NoRelatedSym) HasFXCMSymMarginRatio() bool {
	return m.Has(tag.FXCMSymMarginRatio)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m NoRelatedSym) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasRoundLot returns true if RoundLot is present, Tag 561.
func (m NoRelatedSym) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

// HasFXCMSymInterestBuy returns true if FXCMSymInterestBuy is present, Tag 9003.
func (m NoRelatedSym) HasFXCMSymInterestBuy() bool {
	return m.Has(tag.FXCMSymInterestBuy)
}

// HasFXCMSymInterestSell returns true if FXCMSymInterestSell is present, Tag 9004.
func (m NoRelatedSym) HasFXCMSymInterestSell() bool {
	return m.Has(tag.FXCMSymInterestSell)
}

// HasFXCMSubscriptionStatus returns true if FXCMSubscriptionStatus is present, Tag 9076.
func (m NoRelatedSym) HasFXCMSubscriptionStatus() bool {
	return m.Has(tag.FXCMSubscriptionStatus)
}

// HasFXCMCondDistStop returns true if FXCMCondDistStop is present, Tag 9090.
func (m NoRelatedSym) HasFXCMCondDistStop() bool {
	return m.Has(tag.FXCMCondDistStop)
}

// HasFXCMCondDistLimit returns true if FXCMCondDistLimit is present, Tag 9091.
func (m NoRelatedSym) HasFXCMCondDistLimit() bool {
	return m.Has(tag.FXCMCondDistLimit)
}

// HasFXCMCondDistEntryStop returns true if FXCMCondDistEntryStop is present, Tag 9092.
func (m NoRelatedSym) HasFXCMCondDistEntryStop() bool {
	return m.Has(tag.FXCMCondDistEntryStop)
}

// HasFXCMCondDistEntryLimit returns true if FXCMCondDistEntryLimit is present, Tag 9093.
func (m NoRelatedSym) HasFXCMCondDistEntryLimit() bool {
	return m.Has(tag.FXCMCondDistEntryLimit)
}

// HasFXCMMaxQuantity returns true if FXCMMaxQuantity is present, Tag 9094.
func (m NoRelatedSym) HasFXCMMaxQuantity() bool {
	return m.Has(tag.FXCMMaxQuantity)
}

// HasFXCMMinQuantity returns true if FXCMMinQuantity is present, Tag 9095.
func (m NoRelatedSym) HasFXCMMinQuantity() bool {
	return m.Has(tag.FXCMMinQuantity)
}

// HasFXCMTradingStatus returns true if FXCMTradingStatus is present, Tag 9096.
func (m NoRelatedSym) HasFXCMTradingStatus() bool {
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

// NoRelatedSymRepeatingGroup is a repeating group, Tag 146.
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup.
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), NewNoEventsRepeatingGroup(), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.FXCMSymID), quickfix.GroupElement(tag.FXCMSymPrecision), quickfix.GroupElement(tag.FXCMSymPointSize), quickfix.GroupElement(tag.FXCMSymSortOrder), quickfix.GroupElement(tag.FXCMProductID), quickfix.GroupElement(tag.FXCMSymMarginRatio), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.RoundLot), quickfix.GroupElement(tag.FXCMSymInterestBuy), quickfix.GroupElement(tag.FXCMSymInterestSell), quickfix.GroupElement(tag.FXCMSubscriptionStatus), quickfix.GroupElement(tag.FXCMCondDistStop), quickfix.GroupElement(tag.FXCMCondDistLimit), quickfix.GroupElement(tag.FXCMCondDistEntryStop), quickfix.GroupElement(tag.FXCMCondDistEntryLimit), quickfix.GroupElement(tag.FXCMMaxQuantity), quickfix.GroupElement(tag.FXCMMinQuantity), quickfix.GroupElement(tag.FXCMTradingStatus)})}
}

// Add create and append a new NoRelatedSym to this group.
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

// Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup.
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}

// FXCMNoParam is a repeating group element, Tag 9016.
type FXCMNoParam struct {
	*quickfix.Group
}

// SetFXCMParamName sets FXCMParamName, Tag 9017.
func (m FXCMNoParam) SetFXCMParamName(v string) {
	m.Set(field.NewFXCMParamName(v))
}

// SetFXCMParamValue sets FXCMParamValue, Tag 9018.
func (m FXCMNoParam) SetFXCMParamValue(v string) {
	m.Set(field.NewFXCMParamValue(v))
}

// GetFXCMParamName gets FXCMParamName, Tag 9017.
func (m FXCMNoParam) GetFXCMParamName() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMParamNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMParamValue gets FXCMParamValue, Tag 9018.
func (m FXCMNoParam) GetFXCMParamValue() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMParamValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasFXCMParamName returns true if FXCMParamName is present, Tag 9017.
func (m FXCMNoParam) HasFXCMParamName() bool {
	return m.Has(tag.FXCMParamName)
}

// HasFXCMParamValue returns true if FXCMParamValue is present, Tag 9018.
func (m FXCMNoParam) HasFXCMParamValue() bool {
	return m.Has(tag.FXCMParamValue)
}

// FXCMNoParamRepeatingGroup is a repeating group, Tag 9016.
type FXCMNoParamRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewFXCMNoParamRepeatingGroup returns an initialized, FXCMNoParamRepeatingGroup.
func NewFXCMNoParamRepeatingGroup() FXCMNoParamRepeatingGroup {
	return FXCMNoParamRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.FXCMNoParam,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.FXCMParamName), quickfix.GroupElement(tag.FXCMParamValue)})}
}

// Add create and append a new FXCMNoParam to this group.
func (m FXCMNoParamRepeatingGroup) Add() FXCMNoParam {
	g := m.RepeatingGroup.Add()
	return FXCMNoParam{g}
}

// Get returns the ith FXCMNoParam in the FXCMNoParamRepeatinGroup.
func (m FXCMNoParamRepeatingGroup) Get(i int) FXCMNoParam {
	return FXCMNoParam{m.RepeatingGroup.Get(i)}
}
