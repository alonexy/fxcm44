package fxcmnewstopicrequest

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// FXCMNewsTopicRequest is the fix44 FXCMNewsTopicRequest type, MsgType = U51.
type FXCMNewsTopicRequest struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a FXCMNewsTopicRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) FXCMNewsTopicRequest {
	return FXCMNewsTopicRequest{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m FXCMNewsTopicRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a FXCMNewsTopicRequest initialized with the required fields for FXCMNewsTopicRequest.
func New(testreqid field.TestReqIDField, tradingsessionid field.TradingSessionIDField, tradingsessionsubid field.TradingSessionSubIDField) (m FXCMNewsTopicRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U51"))
	m.Set(testreqid)
	m.Set(tradingsessionid)
	m.Set(tradingsessionsubid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg FXCMNewsTopicRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "U51", r
}

// SetTestReqID sets TestReqID, Tag 112.
func (m FXCMNewsTopicRequest) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m FXCMNewsTopicRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m FXCMNewsTopicRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetTestReqID gets TestReqID, Tag 112.
func (m FXCMNewsTopicRequest) GetTestReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TestReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m FXCMNewsTopicRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m FXCMNewsTopicRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTestReqID returns true if TestReqID is present, Tag 112.
func (m FXCMNewsTopicRequest) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m FXCMNewsTopicRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m FXCMNewsTopicRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
