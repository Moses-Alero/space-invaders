package models

type Bullet struct {
	GameObjectModel
	Speed  int
	Damage int
}

func (b *Bullet) New() Bullet{
	return Bullet{
		Speed: b.Speed,
	}
}

func (b *Bullet) Fire() {
	b.Position.Y += float64(-1 * b.Speed)
}

