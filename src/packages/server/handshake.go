package packageserver

import (
	"BattleArchive-server/src/packages"
)

type Handshake struct {
	Id byte
}

const (
	HandshakeId = byte(0x00)
)

func (p Handshake) Encode() []byte {
	writer := packages.NewWriter()
	writer.WriteByte(p.Id)
	return writer.EncodePackage(HandshakeId)
}
