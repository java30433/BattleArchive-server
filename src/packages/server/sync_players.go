package packageserver

import (
	"BattleArchive-server/src/game"
	"BattleArchive-server/src/packages"
)

const (
	SyncPlayersId = byte(0x02)
)

func EncodeSyncPlayers() []byte {
	writer := packages.NewWriter()
	var players = game.GetPlayers()
	writer.WriteByte(byte(len(players)))
	for id, player := range players {
		writer.WriteByte(id)
		writer.WriteVector3(player.Position)
		writer.WriteFloat32(player.RotationY)
		writer.WriteFloat32(player.Speed)
		writer.WriteBool(player.IsAiming)
		writer.WriteBool(player.IsReloading)
		writer.WriteShort(player.Health)
	}
	return writer.EncodePackage(SyncPlayersId)
}
