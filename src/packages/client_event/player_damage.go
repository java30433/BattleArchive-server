package clientevent

import (
	"BattleArchive-server/src/game"
	"BattleArchive-server/src/packages"
	"fmt"
)

const (
	PlayerDamageId = byte(0x05)
)

func ApplyPlayerDamage(data []byte) {
	var reader = packages.NewReader(data)
	_ = reader.ReadByte() //senderId
	targetId := reader.ReadByte()
	_ = reader.ReadVector3() //pos
	damage := reader.ReadByte()
	game.GetPlayer(targetId).Health -= int16(damage)
	fmt.Println("玩家受伤：", targetId, damage)
}
