package fxcmrequestreject

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// FXCMRequestReject is the fix44 FXCMRequestReject type, MsgType = U52.
type FXCMRequestReject struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a FXCMRequestReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) FXCMRequestReject {
	return FXCMRequestReject{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m FXCMRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a FXCMRequestReject initialized with the required fields for FXCMRequestReject.
func New(testreqid field.TestReqIDField, fxcmrequestrejectreason field.FXCMRequestRejectReasonField) (m FXCMRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("U52"))
	m.Set(testreqid)
	m.Set(fxcmrequestrejectreason)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg FXCMRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "U52", r
}

// SetText sets Text, Tag 58.
func (m FXCMRequestReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTestReqID sets TestReqID, Tag 112.
func (m FXCMRequestReject) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m FXCMRequestReject) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m FXCMRequestReject) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetFXCMRequestRejectReason sets FXCMRequestRejectReason, Tag 9025.
func (m FXCMRequestReject) SetFXCMRequestRejectReason(v enum.FXCMRequestRejectReason) {
	m.Set(field.NewFXCMRequestRejectReason(v))
}

// SetFXCMErrorDetails sets FXCMErrorDetails, Tag 9029.
func (m FXCMRequestReject) SetFXCMErrorDetails(v string) {
	m.Set(field.NewFXCMErrorDetails(v))
}

// GetText gets Text, Tag 58.
func (m FXCMRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTestReqID gets TestReqID, Tag 112.
func (m FXCMRequestReject) GetTestReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TestReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m FXCMRequestReject) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m FXCMRequestReject) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMRequestRejectReason gets FXCMRequestRejectReason, Tag 9025.
func (m FXCMRequestReject) GetFXCMRequestRejectReason() (v enum.FXCMRequestRejectReason, err quickfix.MessageRejectError) {
	var f field.FXCMRequestRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMErrorDetails gets FXCMErrorDetails, Tag 9029.
func (m FXCMRequestReject) GetFXCMErrorDetails() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMErrorDetailsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m FXCMRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasTestReqID returns true if TestReqID is present, Tag 112.
func (m FXCMRequestReject) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m FXCMRequestReject) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m FXCMRequestReject) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasFXCMRequestRejectReason returns true if FXCMRequestRejectReason is present, Tag 9025.
func (m FXCMRequestReject) HasFXCMRequestRejectReason() bool {
	return m.Has(tag.FXCMRequestRejectReason)
}

// HasFXCMErrorDetails returns true if FXCMErrorDetails is present, Tag 9029.
func (m FXCMRequestReject) HasFXCMErrorDetails() bool {
	return m.Has(tag.FXCMErrorDetails)
}
