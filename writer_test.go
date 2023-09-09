package gobuffer

import (
	"bytes"
	"testing"
)

func TestWriter_WriteUint8(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x12}
	w.WriteUint8(0x12)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteInt8(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x12}
	w.WriteInt8(0x12)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteUint16(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x34, 0x12}
	w.WriteUint16(0x1234)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteInt16(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x34, 0x12}
	w.WriteInt16(0x1234)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteUint32(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x78, 0x56, 0x34, 0x12}
	w.WriteUint32(0x12345678)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteInt32(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x78, 0x56, 0x34, 0x12}
	w.WriteInt32(0x12345678)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteFloat32(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0xDB, 0x0F, 0x49, 0x40}
	w.WriteFloat32(3.14159265)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}

func TestWriter_WriteFloat64(t *testing.T) {
	w := NewWriter(true)

	expected := []byte{0x18, 0x2D, 0x44, 0x54, 0xFB, 0x21, 0x09, 0x40}
	w.WriteFloat64(3.141592653589793)

	if !bytes.Equal(w.Build(), expected) {
		t.Errorf("Expected %v, but got %v", expected, w.Build())
	}
}
