package gobuffer

import (
	"bytes"
	"encoding/binary"
	"math"
)

type Writer struct {
	buf          bytes.Buffer
	tmpBuf       []byte
	littleEndian bool
}

func NewWriter(littleEndian bool) *Writer {
	w := &Writer{
		littleEndian: littleEndian,
		tmpBuf:       make([]byte, 8),
	}
	w.reset(littleEndian)
	return w
}

func (w *Writer) reset(littleEndian bool) {
	w.littleEndian = littleEndian
	w.buf.Reset()
	w.tmpBuf = make([]byte, 8)
}

func (w *Writer) WriteUint8(value uint8) *Writer {
	w.buf.WriteByte(value)
	return w
}

func (w *Writer) WriteInt8(value int8) *Writer {
	w.buf.WriteByte(byte(value))
	return w
}

func (w *Writer) WriteUint16(value uint16) *Writer {
	binary.LittleEndian.PutUint16(w.tmpBuf, value)
	w.writeTempBuffer(2)
	return w
}

func (w *Writer) WriteInt16(value int16) *Writer {
	binary.LittleEndian.PutUint16(w.tmpBuf, uint16(value))
	w.writeTempBuffer(2)
	return w
}

func (w *Writer) WriteUint32(value uint32) *Writer {
	binary.LittleEndian.PutUint32(w.tmpBuf, value)
	w.writeTempBuffer(4)
	return w
}

func (w *Writer) WriteInt32(value int32) *Writer {
	binary.LittleEndian.PutUint32(w.tmpBuf, uint32(value))
	w.writeTempBuffer(4)
	return w
}

func (w *Writer) WriteFloat32(value float32) *Writer {
	bits := math.Float32bits(value)
	binary.LittleEndian.PutUint32(w.tmpBuf, bits)
	w.writeTempBuffer(4)
	return w
}

func (w *Writer) WriteFloat64(value float64) *Writer {
	bits := math.Float64bits(value)
	binary.LittleEndian.PutUint64(w.tmpBuf, bits)
	w.writeTempBuffer(8)
	return w
}

func (w *Writer) writeTempBuffer(size int) {
	w.buf.Write(w.tmpBuf[:size])
}

func (w *Writer) Build() []byte {
	return w.buf.Bytes()
}
