package game

import (
	"BattleArchive-server/src/model"
)

var players = make(map[byte]*model.Player)

func GetPlayer(id byte) *model.Player {
	if _, ok := players[id]; !ok {
		players[id] = &model.Player{}
		players[id].Health = 100
	}
	return players[id]
}

func GetPlayers() map[byte]*model.Player {
	return players
}
