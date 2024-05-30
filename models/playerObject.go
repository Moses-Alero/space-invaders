package models

import (
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/Moses-Alero/space-invaders/utils"
)


var speed int = 200 / ebiten.TPS()

type Player struct {
	GameObjectModel
	Bullet         Bullet
	Bullets        []*Bullet
	Health         int
	AttackTimer    *utils.Timer
	AttackCoolDown time.Duration
}

func (p *Player) Movement() {
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = float64(speed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y -= float64(speed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X -= float64(speed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = float64(speed)
	}

	//check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := float64(speed) / math.Sqrt(math.Pow(delta.X, 2) + math.Pow(delta.Y, 2))
		delta.X *= factor
		delta.Y *= factor
	}

	p.Position.X += delta.X
	p.Position.Y += delta.Y

}

func (p *Player) Attack() {
	p.AttackTimer.Update()

	if p.AttackTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.AttackTimer.Reset()
		p.SpawnBullet()
	}
}

func (p *Player) Controls() {
	p.Movement()
	p.Attack()
}

func (p *Player) SpawnBullet() {
	p.Bullet.Speed = 300 / ebiten.TPS()
	spawnPosRight := Vector{
		X: p.GetCenter().X,
		Y: p.GetCenter().Y - 20,
	}

	bRModel := p.Bullet.Spawn(spawnPosRight)
	bRModel.Name = fmt.Sprintf("Bullet%v", len(p.Bullets))
	bR := p.Bullet.New()
	bR.GameObjectModel = bRModel
	p.Bullets = append(p.Bullets, &bR)

}

