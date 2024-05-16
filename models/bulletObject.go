package models

type Bullet struct {
	GameObjectModel
	Speed  int
	Damage int
}

func (b *Bullet) Spawn(pos Vector) Bullet {
	b.Position = Vector{
		X: pos.X,
		Y: pos.Y,
	}

	gom := GameObjectModel{
		Position: b.Position,
		Sprite:   b.Sprite,
	}

	bullet := new(Bullet)
	bullet.GameObjectModel = gom
	return *bullet
}

func (b *Bullet) Fire() {
	b.Position.Y += float64(-1 * b.Speed)
}

