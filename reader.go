package gobuffer

import (
	"encoding/binary"
	"math"
	"strings"
)

type Reader struct {
	view   []byte
	offset int
}

func NewReader(value []byte, ab bool) *Reader {
	r := &Reader{}
	if value != nil {
		if ab {
			r.view = value
		} else {
			r.view = value[:]
		}
		r.offset = 0
	}
	return r
}

func (r *Reader) SetView(view []byte) {
	r.offset = 0
	r.view = view
}

func (r *Reader) EndOfBuffer() bool {
	return r.offset >= len(r.view)
}

func (r *Reader) ShiftOffset(number int) {
	r.offset += number
}

func (r *Reader) GetUint8() uint8 {
	val := r.view[r.offset]
	r.offset++
	return val
}

func (r *Reader) GetInt8() int8 {
	val := int8(r.view[r.offset])
	r.offset++
	return val
}

func (r *Reader) GetUint16() uint16 {
	val := binary.LittleEndian.Uint16(r.view[r.offset : r.offset+2])
	r.offset += 2
	return val
}

func (r *Reader) GetInt16() int16 {
	val := int16(binary.LittleEndian.Uint16(r.view[r.offset : r.offset+2]))
	r.offset += 2
	return val
}

func (r *Reader) GetUint32() uint32 {
	val := binary.LittleEndian.Uint32(r.view[r.offset : r.offset+4])
	r.offset += 4
	return val
}

func (r *Reader) GetInt32() int32 {
	val := int32(binary.LittleEndian.Uint32(r.view[r.offset : r.offset+4]))
	r.offset += 4
	return val
}

func (r *Reader) GetFloat32() float32 {
	bits := binary.LittleEndian.Uint32(r.view[r.offset : r.offset+4])
	val := math.Float32frombits(bits)
	r.offset += 4
	return val
}

func (r *Reader) GetFloat64() float64 {
	bits := binary.LittleEndian.Uint64(r.view[r.offset : r.offset+8])
	val := math.Float64frombits(bits)
	r.offset += 8
	return val
}

func (r *Reader) GetStringUTF8() string {
	var sb strings.Builder
	for {
		b := r.GetUint8()
		if b == 0 {
			break
		}
		sb.WriteByte(b)
	}
	return sb.String()
}

func (r *Reader) ReadUTF16String() string {
	var sb strings.Builder
	for !r.EndOfBuffer() {
		fe := r.GetUint16()
		if fe == 0 {
			break
		}
		sb.WriteRune(rune(fe))
	}
	return sb.String()
}
