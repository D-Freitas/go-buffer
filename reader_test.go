package gobuffer

import (
	"testing"
)

func TestReader_GetUint8(t *testing.T) {
	data := []byte{0x12, 0x34, 0x56}
	reader := NewReader(data, true)

	expected := uint8(0x12)
	actual := reader.GetUint8()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetInt8(t *testing.T) {
	data := []byte{0x12, 0xFE, 0x56}
	reader := NewReader(data, true)

	expected := int8(0x12)
	actual := reader.GetInt8()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetUint16(t *testing.T) {
	data := []byte{0x12, 0x34, 0x56}
	reader := NewReader(data, true)

	expected := uint16(0x3412)
	actual := reader.GetUint16()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetInt16(t *testing.T) {
	data := []byte{0x12, 0xFE, 0x56}
	reader := NewReader(data, true)

	expected := int16(-294)
	actual := reader.GetInt16()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetUint32(t *testing.T) {
	data := []byte{0x12, 0x34, 0x56, 0x78}
	reader := NewReader(data, true)

	expected := uint32(0x78563412)
	actual := reader.GetUint32()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetInt32(t *testing.T) {
	data := []byte{0x12, 0xFE, 0xDC, 0xBA}
	reader := NewReader(data, true)

	expected := int32(-1153433606)
	actual := reader.GetInt32()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestReader_GetFloat32(t *testing.T) {
	data := []byte{0xDB, 0x0F, 0x49, 0x40}
	reader := NewReader(data, true)

	expected := float32(3.14159265)
	actual := reader.GetFloat32()

	if actual != expected {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}

func TestReader_GetFloat64(t *testing.T) {
	data := []byte{0x18, 0x2D, 0x44, 0x54, 0xFB, 0x21, 0x09, 0x40}
	reader := NewReader(data, true)

	expected := float64(3.141592653589793)
	actual := reader.GetFloat64()

	if actual != expected {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}

func TestReader_GetStringUTF8(t *testing.T) {
	data := []byte("C++ is good!\x00ExtraData")
	reader := NewReader(data, true)

	expected := "C++ is good!"
	actual := reader.GetStringUTF8()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestReader_ReadUTF16String(t *testing.T) {
	data := []byte{0x48, 0x00, 0x65, 0x00, 0x6C, 0x00, 0x6C, 0x00, 0x6F, 0x00, 0x00, 0x00}
	reader := NewReader(data, true)

	expected := "Hello"
	actual := reader.ReadUTF16String()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
