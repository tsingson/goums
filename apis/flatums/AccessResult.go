// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatums

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// AccessResult  result
type AccessResultT struct {
	Me    *AccessProfileT
	Token string
}

func AccessResultPack(builder *flatbuffers.Builder, t *AccessResultT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	MeOffset := AccessProfilePack(builder, t.Me)
	TokenOffset := builder.CreateString(t.Token)
	AccessResultStart(builder)
	AccessResultAddMe(builder, MeOffset)
	AccessResultAddToken(builder, TokenOffset)
	return AccessResultEnd(builder)
}

func (rcv *AccessResult) UnPackTo(t *AccessResultT) {
	t.Me = rcv.Me(nil).UnPack()
	t.Token = string(rcv.Token())
}

func (rcv *AccessResult) UnPack() *AccessResultT {
	if rcv == nil {
		return nil
	}
	t := &AccessResultT{}
	rcv.UnPackTo(t)
	return t
}

type AccessResult struct {
	_tab flatbuffers.Table
}

func GetRootAsAccessResult(buf []byte, offset flatbuffers.UOffsetT) *AccessResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AccessResult{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AccessResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccessResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AccessResult) Me(obj *AccessProfile) *AccessProfile {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AccessProfile)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *AccessResult) Token() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AccessResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}

func AccessResultAddMe(builder *flatbuffers.Builder, Me flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Me), 0)
}

func AccessResultAddToken(builder *flatbuffers.Builder, Token flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Token), 0)
}

func AccessResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
