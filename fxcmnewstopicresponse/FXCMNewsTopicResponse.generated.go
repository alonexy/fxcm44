package fxcmnewstopicresponse

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// FXCMNewsTopicResponse is the fix44 FXCMNewsTopicResponse type, MsgType = U53.
type FXCMNewsTopicResponse struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a FXCMNewsTopicResponse from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) FXCMNewsTopicResponse {
	return FXCMNewsTopicResponse{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m FXCMNewsTopicResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a FXCMNewsTopicResponse initialized with the required fields for FXCMNewsTopicResponse.
func New(testreqid field.TestReqIDField, tradingsessionid field.TradingSessionIDField, tradingsessionsubid field.TradingSessionSubIDField) (m FXCMNewsTopicResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U53"))
	m.Set(testreqid)
	m.Set(tradingsessionid)
	m.Set(tradingsessionsubid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg FXCMNewsTopicResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "U53", r
}

// SetTestReqID sets TestReqID, Tag 112.
func (m FXCMNewsTopicResponse) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m FXCMNewsTopicResponse) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m FXCMNewsTopicResponse) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetFXCMPageIDNo sets FXCMPageIDNo, Tag 9026.
func (m FXCMNewsTopicResponse) SetFXCMPageIDNo(f FXCMPageIDNoRepeatingGroup) {
	m.SetGroup(f)
}

// GetTestReqID gets TestReqID, Tag 112.
func (m FXCMNewsTopicResponse) GetTestReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TestReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m FXCMNewsTopicResponse) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m FXCMNewsTopicResponse) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMPageIDNo gets FXCMPageIDNo, Tag 9026.
func (m FXCMNewsTopicResponse) GetFXCMPageIDNo() (FXCMPageIDNoRepeatingGroup, quickfix.MessageRejectError) {
	f := NewFXCMPageIDNoRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasTestReqID returns true if TestReqID is present, Tag 112.
func (m FXCMNewsTopicResponse) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m FXCMNewsTopicResponse) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m FXCMNewsTopicResponse) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasFXCMPageIDNo returns true if FXCMPageIDNo is present, Tag 9026.
func (m FXCMNewsTopicResponse) HasFXCMPageIDNo() bool {
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

// SetHeadline sets Headline, Tag 148.
func (m FXCMPageIDNo) SetHeadline(v string) {
	m.Set(field.NewHeadline(v))
}

// GetFXCMPageID gets FXCMPageID, Tag 9022.
func (m FXCMPageIDNo) GetFXCMPageID() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMPageIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHeadline gets Headline, Tag 148.
func (m FXCMPageIDNo) GetHeadline() (v string, err quickfix.MessageRejectError) {
	var f field.HeadlineField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasFXCMPageID returns true if FXCMPageID is present, Tag 9022.
func (m FXCMPageIDNo) HasFXCMPageID() bool {
	return m.Has(tag.FXCMPageID)
}

// HasHeadline returns true if Headline is present, Tag 148.
func (m FXCMPageIDNo) HasHeadline() bool {
	return m.Has(tag.Headline)
}

// FXCMPageIDNoRepeatingGroup is a repeating group, Tag 9026.
type FXCMPageIDNoRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewFXCMPageIDNoRepeatingGroup returns an initialized, FXCMPageIDNoRepeatingGroup.
func NewFXCMPageIDNoRepeatingGroup() FXCMPageIDNoRepeatingGroup {
	return FXCMPageIDNoRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.FXCMPageIDNo,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.FXCMPageID), quickfix.GroupElement(tag.Headline)})}
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
