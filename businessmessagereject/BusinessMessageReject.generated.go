package businessmessagereject

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// BusinessMessageReject is the fix44 BusinessMessageReject type, MsgType = j.
type BusinessMessageReject struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a BusinessMessageReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) BusinessMessageReject {
	return BusinessMessageReject{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m BusinessMessageReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a BusinessMessageReject initialized with the required fields for BusinessMessageReject.
func New(refmsgtype field.RefMsgTypeField, businessrejectreason field.BusinessRejectReasonField) (m BusinessMessageReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("j"))
	m.Set(refmsgtype)
	m.Set(businessrejectreason)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg BusinessMessageReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "j", r
}

// SetRefSeqNum sets RefSeqNum, Tag 45.
func (m BusinessMessageReject) SetRefSeqNum(v int) {
	m.Set(field.NewRefSeqNum(v))
}

// SetText sets Text, Tag 58.
func (m BusinessMessageReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m BusinessMessageReject) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m BusinessMessageReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m BusinessMessageReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetRefMsgType sets RefMsgType, Tag 372.
func (m BusinessMessageReject) SetRefMsgType(v string) {
	m.Set(field.NewRefMsgType(v))
}

// SetBusinessRejectRefID sets BusinessRejectRefID, Tag 379.
func (m BusinessMessageReject) SetBusinessRejectRefID(v string) {
	m.Set(field.NewBusinessRejectRefID(v))
}

// SetBusinessRejectReason sets BusinessRejectReason, Tag 380.
func (m BusinessMessageReject) SetBusinessRejectReason(v enum.BusinessRejectReason) {
	m.Set(field.NewBusinessRejectReason(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m BusinessMessageReject) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetFXCMRequestRejectReason sets FXCMRequestRejectReason, Tag 9025.
func (m BusinessMessageReject) SetFXCMRequestRejectReason(v enum.FXCMRequestRejectReason) {
	m.Set(field.NewFXCMRequestRejectReason(v))
}

// SetFXCMErrorDetails sets FXCMErrorDetails, Tag 9029.
func (m BusinessMessageReject) SetFXCMErrorDetails(v string) {
	m.Set(field.NewFXCMErrorDetails(v))
}

// GetRefSeqNum gets RefSeqNum, Tag 45.
func (m BusinessMessageReject) GetRefSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.RefSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m BusinessMessageReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m BusinessMessageReject) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m BusinessMessageReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m BusinessMessageReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRefMsgType gets RefMsgType, Tag 372.
func (m BusinessMessageReject) GetRefMsgType() (v string, err quickfix.MessageRejectError) {
	var f field.RefMsgTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBusinessRejectRefID gets BusinessRejectRefID, Tag 379.
func (m BusinessMessageReject) GetBusinessRejectRefID() (v string, err quickfix.MessageRejectError) {
	var f field.BusinessRejectRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBusinessRejectReason gets BusinessRejectReason, Tag 380.
func (m BusinessMessageReject) GetBusinessRejectReason() (v enum.BusinessRejectReason, err quickfix.MessageRejectError) {
	var f field.BusinessRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m BusinessMessageReject) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMRequestRejectReason gets FXCMRequestRejectReason, Tag 9025.
func (m BusinessMessageReject) GetFXCMRequestRejectReason() (v enum.FXCMRequestRejectReason, err quickfix.MessageRejectError) {
	var f field.FXCMRequestRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFXCMErrorDetails gets FXCMErrorDetails, Tag 9029.
func (m BusinessMessageReject) GetFXCMErrorDetails() (v string, err quickfix.MessageRejectError) {
	var f field.FXCMErrorDetailsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRefSeqNum returns true if RefSeqNum is present, Tag 45.
func (m BusinessMessageReject) HasRefSeqNum() bool {
	return m.Has(tag.RefSeqNum)
}

// HasText returns true if Text is present, Tag 58.
func (m BusinessMessageReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m BusinessMessageReject) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m BusinessMessageReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m BusinessMessageReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasRefMsgType returns true if RefMsgType is present, Tag 372.
func (m BusinessMessageReject) HasRefMsgType() bool {
	return m.Has(tag.RefMsgType)
}

// HasBusinessRejectRefID returns true if BusinessRejectRefID is present, Tag 379.
func (m BusinessMessageReject) HasBusinessRejectRefID() bool {
	return m.Has(tag.BusinessRejectRefID)
}

// HasBusinessRejectReason returns true if BusinessRejectReason is present, Tag 380.
func (m BusinessMessageReject) HasBusinessRejectReason() bool {
	return m.Has(tag.BusinessRejectReason)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m BusinessMessageReject) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasFXCMRequestRejectReason returns true if FXCMRequestRejectReason is present, Tag 9025.
func (m BusinessMessageReject) HasFXCMRequestRejectReason() bool {
	return m.Has(tag.FXCMRequestRejectReason)
}

// HasFXCMErrorDetails returns true if FXCMErrorDetails is present, Tag 9029.
func (m BusinessMessageReject) HasFXCMErrorDetails() bool {
	return m.Has(tag.FXCMErrorDetails)
}
