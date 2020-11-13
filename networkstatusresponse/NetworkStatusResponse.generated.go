package networkstatusresponse

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// NetworkStatusResponse is the fix44 NetworkStatusResponse type, MsgType = BD.
type NetworkStatusResponse struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a NetworkStatusResponse from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) NetworkStatusResponse {
	return NetworkStatusResponse{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m NetworkStatusResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a NetworkStatusResponse initialized with the required fields for NetworkStatusResponse.
func New(networkstatusresponsetype field.NetworkStatusResponseTypeField) (m NetworkStatusResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BD"))
	m.Set(networkstatusresponsetype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg NetworkStatusResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BD", r
}

// SetNetworkResponseID sets NetworkResponseID, Tag 932.
func (m NetworkStatusResponse) SetNetworkResponseID(v string) {
	m.Set(field.NewNetworkResponseID(v))
}

// SetNetworkRequestID sets NetworkRequestID, Tag 933.
func (m NetworkStatusResponse) SetNetworkRequestID(v string) {
	m.Set(field.NewNetworkRequestID(v))
}

// SetLastNetworkResponseID sets LastNetworkResponseID, Tag 934.
func (m NetworkStatusResponse) SetLastNetworkResponseID(v string) {
	m.Set(field.NewLastNetworkResponseID(v))
}

// SetNoCompIDs sets NoCompIDs, Tag 936.
func (m NetworkStatusResponse) SetNoCompIDs(f NoCompIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNetworkStatusResponseType sets NetworkStatusResponseType, Tag 937.
func (m NetworkStatusResponse) SetNetworkStatusResponseType(v enum.NetworkStatusResponseType) {
	m.Set(field.NewNetworkStatusResponseType(v))
}

// GetNetworkResponseID gets NetworkResponseID, Tag 932.
func (m NetworkStatusResponse) GetNetworkResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.NetworkResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNetworkRequestID gets NetworkRequestID, Tag 933.
func (m NetworkStatusResponse) GetNetworkRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.NetworkRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastNetworkResponseID gets LastNetworkResponseID, Tag 934.
func (m NetworkStatusResponse) GetLastNetworkResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.LastNetworkResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoCompIDs gets NoCompIDs, Tag 936.
func (m NetworkStatusResponse) GetNoCompIDs() (NoCompIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCompIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNetworkStatusResponseType gets NetworkStatusResponseType, Tag 937.
func (m NetworkStatusResponse) GetNetworkStatusResponseType() (v enum.NetworkStatusResponseType, err quickfix.MessageRejectError) {
	var f field.NetworkStatusResponseTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNetworkResponseID returns true if NetworkResponseID is present, Tag 932.
func (m NetworkStatusResponse) HasNetworkResponseID() bool {
	return m.Has(tag.NetworkResponseID)
}

// HasNetworkRequestID returns true if NetworkRequestID is present, Tag 933.
func (m NetworkStatusResponse) HasNetworkRequestID() bool {
	return m.Has(tag.NetworkRequestID)
}

// HasLastNetworkResponseID returns true if LastNetworkResponseID is present, Tag 934.
func (m NetworkStatusResponse) HasLastNetworkResponseID() bool {
	return m.Has(tag.LastNetworkResponseID)
}

// HasNoCompIDs returns true if NoCompIDs is present, Tag 936.
func (m NetworkStatusResponse) HasNoCompIDs() bool {
	return m.Has(tag.NoCompIDs)
}

// HasNetworkStatusResponseType returns true if NetworkStatusResponseType is present, Tag 937.
func (m NetworkStatusResponse) HasNetworkStatusResponseType() bool {
	return m.Has(tag.NetworkStatusResponseType)
}

// NoCompIDs is a repeating group element, Tag 936.
type NoCompIDs struct {
	*quickfix.Group
}

// SetRefCompID sets RefCompID, Tag 930.
func (m NoCompIDs) SetRefCompID(v string) {
	m.Set(field.NewRefCompID(v))
}

// SetRefSubID sets RefSubID, Tag 931.
func (m NoCompIDs) SetRefSubID(v string) {
	m.Set(field.NewRefSubID(v))
}

// SetLocationID sets LocationID, Tag 283.
func (m NoCompIDs) SetLocationID(v string) {
	m.Set(field.NewLocationID(v))
}

// SetDeskID sets DeskID, Tag 284.
func (m NoCompIDs) SetDeskID(v string) {
	m.Set(field.NewDeskID(v))
}

// SetStatusValue sets StatusValue, Tag 928.
func (m NoCompIDs) SetStatusValue(v enum.StatusValue) {
	m.Set(field.NewStatusValue(v))
}

// SetStatusText sets StatusText, Tag 929.
func (m NoCompIDs) SetStatusText(v string) {
	m.Set(field.NewStatusText(v))
}

// GetRefCompID gets RefCompID, Tag 930.
func (m NoCompIDs) GetRefCompID() (v string, err quickfix.MessageRejectError) {
	var f field.RefCompIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRefSubID gets RefSubID, Tag 931.
func (m NoCompIDs) GetRefSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RefSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocationID gets LocationID, Tag 283.
func (m NoCompIDs) GetLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.LocationIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeskID gets DeskID, Tag 284.
func (m NoCompIDs) GetDeskID() (v string, err quickfix.MessageRejectError) {
	var f field.DeskIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStatusValue gets StatusValue, Tag 928.
func (m NoCompIDs) GetStatusValue() (v enum.StatusValue, err quickfix.MessageRejectError) {
	var f field.StatusValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStatusText gets StatusText, Tag 929.
func (m NoCompIDs) GetStatusText() (v string, err quickfix.MessageRejectError) {
	var f field.StatusTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRefCompID returns true if RefCompID is present, Tag 930.
func (m NoCompIDs) HasRefCompID() bool {
	return m.Has(tag.RefCompID)
}

// HasRefSubID returns true if RefSubID is present, Tag 931.
func (m NoCompIDs) HasRefSubID() bool {
	return m.Has(tag.RefSubID)
}

// HasLocationID returns true if LocationID is present, Tag 283.
func (m NoCompIDs) HasLocationID() bool {
	return m.Has(tag.LocationID)
}

// HasDeskID returns true if DeskID is present, Tag 284.
func (m NoCompIDs) HasDeskID() bool {
	return m.Has(tag.DeskID)
}

// HasStatusValue returns true if StatusValue is present, Tag 928.
func (m NoCompIDs) HasStatusValue() bool {
	return m.Has(tag.StatusValue)
}

// HasStatusText returns true if StatusText is present, Tag 929.
func (m NoCompIDs) HasStatusText() bool {
	return m.Has(tag.StatusText)
}

// NoCompIDsRepeatingGroup is a repeating group, Tag 936.
type NoCompIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoCompIDsRepeatingGroup returns an initialized, NoCompIDsRepeatingGroup.
func NewNoCompIDsRepeatingGroup() NoCompIDsRepeatingGroup {
	return NoCompIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoCompIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefCompID), quickfix.GroupElement(tag.RefSubID), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID), quickfix.GroupElement(tag.StatusValue), quickfix.GroupElement(tag.StatusText)})}
}

// Add create and append a new NoCompIDs to this group.
func (m NoCompIDsRepeatingGroup) Add() NoCompIDs {
	g := m.RepeatingGroup.Add()
	return NoCompIDs{g}
}

// Get returns the ith NoCompIDs in the NoCompIDsRepeatinGroup.
func (m NoCompIDsRepeatingGroup) Get(i int) NoCompIDs {
	return NoCompIDs{m.RepeatingGroup.Get(i)}
}
