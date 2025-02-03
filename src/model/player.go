package model

import "BattleArchive-server/src/packages"

type Player struct {
	Position    packages.Vector3
	RotationY   float32
	Speed       float32
	IsAiming    bool
	Name        string
	IsReloading bool
	Health      int16
}
