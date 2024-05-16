package models

// Enemy Region
type Enemy struct {
	GameObjectModel
	Bullet Bullet
}

// Barrier
type Barrier struct {
	GameObjectModel
}
