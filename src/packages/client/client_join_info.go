package packageclient

import (
	"BattleArchive-server/src/game"
	"BattleArchive-server/src/packages"
)

type ClientJoinInfo struct {
	PlayerId byte
	Name     string
}

const (
	ClientJoinInfoId = byte(0x01)
)

func (p *ClientJoinInfo) Read(data []byte) {
	var reader = packages.NewReader(data)
	p.PlayerId = reader.ReadByte()
	p.Name = reader.ReadString()
	game.GetPlayer(p.PlayerId).Name = p.Name
}
