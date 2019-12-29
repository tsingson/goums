// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatums

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// TerminalProfile 认证成功后返回用户详细档案
type TerminalProfileT struct {
	UserID            int64
	ActiveStatus      bool
	ActiveDate        int64
	MaxActiveSession  int64
	ServiceStatus     int64
	ServiceExpiration int64
	SerialNumber      string
	ActiveCode        string
	AccessRole        string
	Operation         string
}

func TerminalProfilePack(builder *flatbuffers.Builder, t *TerminalProfileT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	serialNumberOffset := builder.CreateString(t.SerialNumber)
	activeCodeOffset := builder.CreateString(t.ActiveCode)
	accessRoleOffset := builder.CreateString(t.AccessRole)
	operationOffset := builder.CreateString(t.Operation)
	TerminalProfileStart(builder)
	TerminalProfileAddUserID(builder, t.UserID)
	TerminalProfileAddActiveStatus(builder, t.ActiveStatus)
	TerminalProfileAddActiveDate(builder, t.ActiveDate)
	TerminalProfileAddMaxActiveSession(builder, t.MaxActiveSession)
	TerminalProfileAddServiceStatus(builder, t.ServiceStatus)
	TerminalProfileAddServiceExpiration(builder, t.ServiceExpiration)
	TerminalProfileAddSerialNumber(builder, serialNumberOffset)
	TerminalProfileAddActiveCode(builder, activeCodeOffset)
	TerminalProfileAddAccessRole(builder, accessRoleOffset)
	TerminalProfileAddOperation(builder, operationOffset)
	return TerminalProfileEnd(builder)
}

func (rcv *TerminalProfile) UnPackTo(t *TerminalProfileT) {
	t.UserID = rcv.UserID()
	t.ActiveStatus = rcv.ActiveStatus()
	t.ActiveDate = rcv.ActiveDate()
	t.MaxActiveSession = rcv.MaxActiveSession()
	t.ServiceStatus = rcv.ServiceStatus()
	t.ServiceExpiration = rcv.ServiceExpiration()
	t.SerialNumber = string(rcv.SerialNumber())
	t.ActiveCode = string(rcv.ActiveCode())
	t.AccessRole = string(rcv.AccessRole())
	t.Operation = string(rcv.Operation())
}

func (rcv *TerminalProfile) UnPack() *TerminalProfileT {
	if rcv == nil {
		return nil
	}
	t := &TerminalProfileT{}
	rcv.UnPackTo(t)
	return t
}

type TerminalProfile struct {
	_tab flatbuffers.Table
}

func GetRootAsTerminalProfile(buf []byte, offset flatbuffers.UOffsetT) *TerminalProfile {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TerminalProfile{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TerminalProfile) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TerminalProfile) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TerminalProfile) UserID() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalProfile) MutateUserID(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *TerminalProfile) ActiveStatus() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *TerminalProfile) MutateActiveStatus(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *TerminalProfile) ActiveDate() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalProfile) MutateActiveDate(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *TerminalProfile) MaxActiveSession() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalProfile) MutateMaxActiveSession(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *TerminalProfile) ServiceStatus() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalProfile) MutateServiceStatus(n int64) bool {
	return rcv._tab.MutateInt64Slot(12, n)
}

func (rcv *TerminalProfile) ServiceExpiration() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalProfile) MutateServiceExpiration(n int64) bool {
	return rcv._tab.MutateInt64Slot(14, n)
}

func (rcv *TerminalProfile) SerialNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalProfile) ActiveCode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalProfile) AccessRole() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalProfile) Operation() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TerminalProfileStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}

func TerminalProfileAddUserID(builder *flatbuffers.Builder, userID int64) {
	builder.PrependInt64Slot(0, userID, 0)
}

func TerminalProfileAddActiveStatus(builder *flatbuffers.Builder, activeStatus bool) {
	builder.PrependBoolSlot(1, activeStatus, false)
}

func TerminalProfileAddActiveDate(builder *flatbuffers.Builder, activeDate int64) {
	builder.PrependInt64Slot(2, activeDate, 0)
}

func TerminalProfileAddMaxActiveSession(builder *flatbuffers.Builder, maxActiveSession int64) {
	builder.PrependInt64Slot(3, maxActiveSession, 0)
}

func TerminalProfileAddServiceStatus(builder *flatbuffers.Builder, serviceStatus int64) {
	builder.PrependInt64Slot(4, serviceStatus, 0)
}

func TerminalProfileAddServiceExpiration(builder *flatbuffers.Builder, serviceExpiration int64) {
	builder.PrependInt64Slot(5, serviceExpiration, 0)
}

func TerminalProfileAddSerialNumber(builder *flatbuffers.Builder, serialNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(serialNumber), 0)
}

func TerminalProfileAddActiveCode(builder *flatbuffers.Builder, activeCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(activeCode), 0)
}

func TerminalProfileAddAccessRole(builder *flatbuffers.Builder, accessRole flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(accessRole), 0)
}

func TerminalProfileAddOperation(builder *flatbuffers.Builder, operation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(operation), 0)
}

func TerminalProfileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
