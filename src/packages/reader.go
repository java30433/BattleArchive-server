package packages

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Reader struct {
	reader io.Reader
}

func NewReader(data []byte) *Reader {
	reader := bytes.NewReader(data)
	// 跳过第一个字节（包ID）
	reader.ReadByte()
	return &Reader{
		reader: reader,
	}
}

func (p *Reader) ReadString() string {
	var length uint16
	binary.Read(p.reader, binary.NativeEndian, &length)
	var data = make([]byte, length)
	p.reader.Read(data)
	return string(data)
}

func (p *Reader) ReadByte() byte {
	var data byte
	binary.Read(p.reader, binary.NativeEndian, &data)
	return data
}

func (p *Reader) ReadBool() bool {
	var data bool
	binary.Read(p.reader, binary.NativeEndian, &data)
	return data
}

func (p *Reader) ReadFloat32() float32 {
	var data float32
	binary.Read(p.reader, binary.NativeEndian, &data)
	return data
}

func (p *Reader) ReadVector3() Vector3 {
	var data Vector3
	data.X = p.ReadFloat32()
	data.Y = p.ReadFloat32()
	data.Z = p.ReadFloat32()
	return data
}
