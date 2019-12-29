// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatums

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Result define sn response
type ResultT struct {
	Tid     int64
	Code    int64
	Message string
}

func ResultPack(builder *flatbuffers.Builder, t *ResultT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	messageOffset := builder.CreateString(t.Message)
	ResultStart(builder)
	ResultAddTid(builder, t.Tid)
	ResultAddCode(builder, t.Code)
	ResultAddMessage(builder, messageOffset)
	return ResultEnd(builder)
}

func (rcv *Result) UnPackTo(t *ResultT) {
	t.Tid = rcv.Tid()
	t.Code = rcv.Code()
	t.Message = string(rcv.Message())
}

func (rcv *Result) UnPack() *ResultT {
	if rcv == nil {
		return nil
	}
	t := &ResultT{}
	rcv.UnPackTo(t)
	return t
}

type Result struct {
	_tab flatbuffers.Table
}

func GetRootAsResult(buf []byte, offset flatbuffers.UOffsetT) *Result {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Result{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Result) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Result) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Result) Tid() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Result) MutateTid(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Result) Code() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Result) MutateCode(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *Result) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}

func ResultAddTid(builder *flatbuffers.Builder, tid int64) {
	builder.PrependInt64Slot(0, tid, 0)
}

func ResultAddCode(builder *flatbuffers.Builder, code int64) {
	builder.PrependInt64Slot(1, code, 0)
}

func ResultAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(message), 0)
}

func ResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
