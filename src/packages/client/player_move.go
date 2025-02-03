package packageclient

import (
	"BattleArchive-server/src/game"
	"BattleArchive-server/src/model"
	"BattleArchive-server/src/packages"
)

const (
	PlayerMoveId = byte(0x03)
)

func ReadPlayerMove(data []byte) *model.Player {
	var reader = packages.NewReader(data)
	var p = game.GetPlayer(reader.ReadByte())
	p.Position = reader.ReadVector3()
	p.RotationY = reader.ReadFloat32()
	p.Speed = reader.ReadFloat32()
	p.IsAiming = reader.ReadBool()
	p.IsReloading = reader.ReadBool()
	return p
}
