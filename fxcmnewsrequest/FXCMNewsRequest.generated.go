package fxcmnewsrequest

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// FXCMNewsRequest is the fix44 FXCMNewsRequest type, MsgType = U50.
type FXCMNewsRequest struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a FXCMNewsRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) FXCMNewsRequest {
	return FXCMNewsRequest{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m FXCMNewsRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a FXCMNewsRequest initialized with the required fields for FXCMNewsRequest.
func New(testreqid field.TestReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField, tradingsessionid field.TradingSessionIDField, tradingsessionsubid field.TradingSessionSubIDField) (m FXCMNewsRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U50"))
	m.Set(testreqid)
	m.Set(subscriptionrequesttype)
	m.Set(tradingsessionid)
	m.Set(tradingsessionsubid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg FXCMNewsRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "U50", r
}

// SetTestReqID sets TestReqID, Tag 112.
func (m FXCMNewsRequest) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m FXCMNewsRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m FXCMNewsRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m FXCMNewsRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetFXCMPageIDNo sets FXCMPageIDNo, Tag 9026.
func (m FXCMNewsRequest) SetFXCMPageIDNo(f FXCMPageIDNoRepeatingGroup) {
	m.SetGroup(f)
}

// GetTestReqID gets TestReqID, Tag 112.
func (m FXCMNewsRequest) GetTestReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TestReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m FXCMNewsRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m FXCMNewsRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m FXCMNewsRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMPageIDNo gets FXCMPageIDNo, Tag 9026.
func (m FXCMNewsRequest) GetFXCMPageIDNo() (FXCMPageIDNoRepeatingGroup, quickfix.MessageRejectError) {
	f := NewFXCMPageIDNoRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasTestReqID returns true if TestReqID is present, Tag 112.
func (m FXCMNewsRequest) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m FXCMNewsRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m FXCMNewsRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m FXCMNewsRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasFXCMPageIDNo returns true if FXCMPageIDNo is present, Tag 9026.
func (m FXCMNewsRequest) HasFXCMPageIDNo() bool {
	return m.Has(tag.FXCMPageIDNo)
}

// FXCMPageIDNo is a repeating group element, Tag 9026.
type FXCMPageIDNo struct {
	*quickfix.Group
}

// SetFXCMPageID sets FXCMPageID, Tag 9022.
func (m FXCMPageIDNo) SetFXCMPageID(v string) {
	m.Set(field.NewFXCMPageID(v))
}

// SetFXCMPageviewID sets FXCMPageviewID, Tag 9023.
func (m FXCMPageIDNo) SetFXCMPageviewID(v int) {
	m.Set(field.NewFXCMPageviewID(v))
}

// SetFXCMPageviewLifetime sets FXCMPageviewLifetime, Tag 9024.
func (m FXCMPageIDNo) SetFXCMPageviewLifetime(v int) {
	m.Set(field.NewFXCMPageviewLifetime(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m FXCMPageIDNo) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetFXCMNoSnapshot sets FXCMNoSnapshot, Tag 9021.
func (m FXCMPageIDNo) SetFXCMNoSnapshot(v int) {
	m.Set(field.NewFXCMNoSnapshot(v))
}

// GetFXCMPageID gets FXCMPageID, Tag 9022.
func (m FXCMPageIDNo) GetFXCMPageID() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMPageIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMPageviewID gets FXCMPageviewID, Tag 9023.
func (m FXCMPageIDNo) GetFXCMPageviewID() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMPageviewIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMPageviewLifetime gets FXCMPageviewLifetime, Tag 9024.
func (m FXCMPageIDNo) GetFXCMPageviewLifetime() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMPageviewLifetimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m FXCMPageIDNo) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMNoSnapshot gets FXCMNoSnapshot, Tag 9021.
func (m FXCMPageIDNo) GetFXCMNoSnapshot() (v int, err quickfix.MessageRejectError) {
	var f field.FXCMNoSnapshotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasFXCMPageID returns true if FXCMPageID is present, Tag 9022.
func (m FXCMPageIDNo) HasFXCMPageID() bool {
	return m.Has(tag.FXCMPageID)
}

// HasFXCMPageviewID returns true if FXCMPageviewID is present, Tag 9023.
func (m FXCMPageIDNo) HasFXCMPageviewID() bool {
	return m.Has(tag.FXCMPageviewID)
}

// HasFXCMPageviewLifetime returns true if FXCMPageviewLifetime is present, Tag 9024.
func (m FXCMPageIDNo) HasFXCMPageviewLifetime() bool {
	return m.Has(tag.FXCMPageviewLifetime)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m FXCMPageIDNo) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasFXCMNoSnapshot returns true if FXCMNoSnapshot is present, Tag 9021.
func (m FXCMPageIDNo) HasFXCMNoSnapshot() bool {
	return m.Has(tag.FXCMNoSnapshot)
}

// FXCMPageIDNoRepeatingGroup is a repeating group, Tag 9026.
type FXCMPageIDNoRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewFXCMPageIDNoRepeatingGroup returns an initialized, FXCMPageIDNoRepeatingGroup.
func NewFXCMPageIDNoRepeatingGroup() FXCMPageIDNoRepeatingGroup {
	return FXCMPageIDNoRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.FXCMPageIDNo,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.FXCMPageID), quickfix.GroupElement(tag.FXCMPageviewID), quickfix.GroupElement(tag.FXCMPageviewLifetime), quickfix.GroupElement(tag.TransactTime), quickfix.GroupElement(tag.FXCMNoSnapshot)})}
}

// Add create and append a new FXCMPageIDNo to this group.
func (m FXCMPageIDNoRepeatingGroup) Add() FXCMPageIDNo {
	g := m.RepeatingGroup.Add()
	return FXCMPageIDNo{g}
}

// Get returns the ith FXCMPageIDNo in the FXCMPageIDNoRepeatinGroup.
func (m FXCMPageIDNoRepeatingGroup) Get(i int) FXCMPageIDNo {
	return FXCMPageIDNo{m.RepeatingGroup.Get(i)}
}
