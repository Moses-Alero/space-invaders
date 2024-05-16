package main

import (
	"fmt"
	"log"

	"github.com/Moses-Alero/space-invaders/models"
	"github.com/Moses-Alero/space-invaders/player"
	"github.com/Moses-Alero/space-invaders/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenX, screenY         = 600, 600
	gridRowSize, gridColSize = 4, 4
)

var playerStartPos = models.Vector{float64(screenX/2) + 50, float64(screenY/2) + 50}
var enemyStartPos = models.Vector{float64(screenX / 3), float64(screenY / 3)}

var ship models.Player
var enemy models.Enemy
var speed int = 200 / ebiten.TPS()
var spaces []*models.Space
var world models.Vector = models.Vector{
	X: screenX,
	Y: screenY,
}

type Game struct {
}


func (g *Game) Update() error {
	for _, s := range spaces {
		ship.SetCurrentSpacePosition(s)
		enemy.SetCurrentSpacePosition(s)
		for _, b := range ship.Bullets {
			b.SetCurrentSpacePosition(s)
		}
	}

	ship.CheckCollision(func() {
		var delta models.Vector
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			delta.Y = -float64(speed)
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			delta.Y = float64(speed)
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			delta.X = float64(speed)
		}

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			delta.X = -float64(speed)
		}

		ship.Position.X += delta.X
		ship.Position.Y += delta.Y

	})

	for _, b := range ship.Bullets {
		b.Fire()
		b.CheckCollision(func() {
			b.Speed *= -1
			b.Fire()
		})
	}

	enemy.CheckCollision(func() {
		//		enemy.Position.Y += 2
	})
	ship.Controls()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// So basically wo i put items we want to draw on the sreen in this area and
	// the items are drawn in a desecding order as they are placed
	utils.DrawGrid(screen, gridRowSize, gridColSize, screenY, screenX)
	for _, b := range ship.Bullets {
		b.Draw(screen)
	}

	for _, s := range spaces {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("POS: %v, %v", s.Position.X, s.Position.Y), int(s.Position.X), int(s.Position.Y))
	}

	ship.Draw(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("."), int(ship.Position.X), int(ship.Position.Y))

	enemy.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingMode() + 1)
	ebiten.SetWindowTitle("Space Invaders")

	spaces = models.CreateSpacePartition(world, gridRowSize, gridColSize)

	ship.Sprite = PlayerSprite
	ship.Name = "Player"
	ship.AttackTimer = utils.NewTimer(player.AttackCoolDown)
	ship.SetPosition(playerStartPos.X, playerStartPos.Y)
	ship.Bullet.Sprite = PlayerBulletSprite

	enemy.Sprite = EnemySprite
	enemy.Name = "Enemy"
	enemy.SetPosition(enemyStartPos.X, enemyStartPos.Y)
	enemy.Bullet.Sprite = PlayerBulletSprite

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

