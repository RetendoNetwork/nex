package nex

import (
	gobuffer "github.com/RetendoNetwork/GoBuffer"
)

type StreamIn struct {
	*gobuffer.GoBuffer
	server *Server
}

func (stream *StreamIn) ReadUInt8() uint8 {
	return uint8(stream.ReadByteNext())
}

func (stream *StreamIn) ReadUInt16LE() uint16 {
	return stream.ReadU16LENext(1)[0]
}

func (stream *StreamIn) ReadUInt32LE() uint32 {
	return stream.ReadU32LENext(1)[0]
}

func (stream *StreamIn) ReadInt32LE() int32 {
	return int32(stream.ReadU32LENext(1)[0])
}

func (stream *StreamIn) ReadUInt64LE() uint64 {
	return stream.ReadU64LENext(1)[0]
}

func NewStreamIn(data []byte, server *Server) *StreamIn {
	return &StreamIn{
		GoBuffer: gobuffer.NewGoBuffer(),
		server:   server,
	}
}

type StreamOut struct {
	*gobuffer.GoBuffer
	server *Server
}

func (stream *StreamOut) WriteUInt8(u8 uint8) {
	stream.Grow(1)
	stream.WriteByteNext(byte(u8))
}

func (stream *StreamOut) WriteUInt16LE(u16 uint16) {
	stream.Grow(2)
	stream.WriteU16LENext([]uint16{u16})
}

func (stream *StreamOut) WriteUInt32LE(u32 uint32) {
	stream.Grow(4)
	stream.WriteU32LENext([]uint32{u32})
}

func (stream *StreamOut) WriteInt32LE(s32 int32) {
	stream.Grow(4)
	stream.WriteU32LENext([]uint32{uint32(s32)})
}

func (stream *StreamOut) WriteUInt64LE(u64 uint64) {
	stream.Grow(8)
	stream.WriteU64LENext([]uint64{u64})
}

func (stream *StreamOut) WriteInt64LE(s64 int64) {
	stream.Grow(8)
	stream.WriteU64LENext([]uint64{uint64(s64)})
}

func NewStreamOut(server *Server) *StreamOut {
	return &StreamOut{
		GoBuffer: gobuffer.NewGoBuffer(),
		server:   server,
	}
}
