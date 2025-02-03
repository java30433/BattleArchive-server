package packages

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Writer struct {
	data   *bytes.Buffer
	writer io.Writer
}

func NewWriter() *Writer {
	data := &bytes.Buffer{}
	return &Writer{
		data:   data,
		writer: data,
	}
}

func (p *Writer) WriteByte(data byte) {
	binary.Write(p.writer, binary.NativeEndian, data)
}

func (p *Writer) WriteFloat32(data float32) {
	binary.Write(p.writer, binary.NativeEndian, data)
}

func (p *Writer) WriteBool(data bool) {
	binary.Write(p.writer, binary.NativeEndian, data)
}

func (p *Writer) WriteShort(data int16) {
	binary.Write(p.writer, binary.NativeEndian, data)
}

func (p *Writer) WriteVector3(data Vector3) {
	p.WriteFloat32(data.X)
	p.WriteFloat32(data.Y)
	p.WriteFloat32(data.Z)
}

func (p *Writer) EncodePackage(id byte) []byte {
	data := p.data.Bytes()
	var length = uint16(len(data) + 1)
	byte1 := byte(length & 0x00FF)
	byte2 := byte((length >> 8) & 0x00FF)
	result := make([]byte, 3+len(data))
	result[0] = byte1
	result[1] = byte2
	result[2] = id
	copy(result[3:], data)
	return result
}

func EncodePackageWithoutId(data []byte) []byte {
	var length = uint16(len(data))
	byte1 := byte(length & 0x00FF)
	byte2 := byte((length >> 8) & 0x00FF)
	result := make([]byte, 2+len(data))
	result[0] = byte1
	result[1] = byte2
	copy(result[2:], data)
	return result
}
