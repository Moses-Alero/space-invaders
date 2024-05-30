package models

import (

	"github.com/Moses-Alero/space-invaders/utils"
)

// Enemy Region
type Enemy struct {
	GameObjectModel
	Health int
	Bullet Bullet
	AttackTimer *utils.Timer
	Bullets []Bullet
}

// Barrier
type Barrier struct {
	GameObjectModel
}
