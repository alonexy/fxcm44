package userresponse

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// UserResponse is the fix44 UserResponse type, MsgType = BF.
type UserResponse struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a UserResponse from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) UserResponse {
	return UserResponse{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m UserResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a UserResponse initialized with the required fields for UserResponse.
func New(userrequestid field.UserRequestIDField, username field.UsernameField) (m UserResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BF"))
	m.Set(userrequestid)
	m.Set(username)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BF", r
}

// SetUsername sets Username, Tag 553.
func (m UserResponse) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

// SetUserRequestID sets UserRequestID, Tag 923.
func (m UserResponse) SetUserRequestID(v string) {
	m.Set(field.NewUserRequestID(v))
}

// SetUserStatus sets UserStatus, Tag 926.
func (m UserResponse) SetUserStatus(v enum.UserStatus) {
	m.Set(field.NewUserStatus(v))
}

// SetUserStatusText sets UserStatusText, Tag 927.
func (m UserResponse) SetUserStatusText(v string) {
	m.Set(field.NewUserStatusText(v))
}

// GetUsername gets Username, Tag 553.
func (m UserResponse) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserRequestID gets UserRequestID, Tag 923.
func (m UserResponse) GetUserRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.UserRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatus gets UserStatus, Tag 926.
func (m UserResponse) GetUserStatus() (v enum.UserStatus, err quickfix.MessageRejectError) {
	var f field.UserStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatusText gets UserStatusText, Tag 927.
func (m UserResponse) GetUserStatusText() (v string, err quickfix.MessageRejectError) {
	var f field.UserStatusTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUsername returns true if Username is present, Tag 553.
func (m UserResponse) HasUsername() bool {
	return m.Has(tag.Username)
}

// HasUserRequestID returns true if UserRequestID is present, Tag 923.
func (m UserResponse) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

// HasUserStatus returns true if UserStatus is present, Tag 926.
func (m UserResponse) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

// HasUserStatusText returns true if UserStatusText is present, Tag 927.
func (m UserResponse) HasUserStatusText() bool {
	return m.Has(tag.UserStatusText)
}
UserResponse) GetFXCMNoParam() (FXCMNoParamRepeatingGroup, quickfix.MessageRejectError) {
	f := NewFXCMNoParamRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386.
func (m UserResponse) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

// HasUsername returns true if Username is present, Tag 553.
func (m UserResponse) HasUsername() bool {
	return m.Has(tag.Username)
}

// HasUserRequestID returns true if UserRequestID is present, Tag 923.
func (m UserResponse) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

// HasUserStatus returns true if UserStatus is present, Tag 926.
func (m UserResponse) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

// HasUserStatusText returns true if UserStatusText is present, Tag 927.
func (m UserResponse) HasUserStatusText() bool {
	return m.Has(tag.UserStatusText)
}

// HasFXCMNoParam returns true if FXCMNoParam is present, Tag 9016.
func (m UserResponse) HasFXCMNoParam() bool {
	return m.Has(tag.FXCMNoParam)
}

// NoTradingSessions is a repeating group element, Tag 386.
type NoTradingSessions struct {
	*quickfix.Group
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m NoTradingSessions) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m NoTradingSessions) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m NoTradingSessions) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m NoTradingSessions) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// NoTradingSessionsRepeatingGroup is a repeating group, Tag 386.
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup.
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID)})}
}

// Add create and append a new NoTradingSessions to this group.
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

// Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup.
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
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
