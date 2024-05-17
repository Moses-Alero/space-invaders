package player

import (
	"time"

	"github.com/Moses-Alero/space-invaders/manager/assets"
	"github.com/Moses-Alero/space-invaders/models"
	"github.com/Moses-Alero/space-invaders/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

var AttackCoolDown time.Duration = time.Millisecond * 500
var speed = 200/ebiten.TPS()

func New() *models.Player{
	player := models.Player {
		Name: "Player",
		Health: 100,
		AttackTimer: utils.NewTimer(AttackCoolDown),
	}
	player.Bullet.Sprite = assets.PlayerBulletSprite
	player.Sprite = assets.PlayerSprite

	return &player
}


func Update(player *models.Player, s []*models.Space){
	
	setPlayerWorldPos(s, player)
	player.CheckCollision(func() {
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

		player.Position.X += delta.X
		player.Position.Y += delta.Y

	})

	for _, b := range player.Bullets {
		b.Fire()
		b.CheckCollision(func() {
			b.Speed *= -1
			b.Fire()
		})
	}

	player.Movement()
	player.Attack()

}

func Draw(screen *ebiten.Image,  player *models.Player){
	for _, b := range player.Bullets {
		b.Draw(screen)
	}
	player.Draw(screen)

}


func setPlayerWorldPos(spaces []*models.Space, player *models.Player){
	for _, s := range spaces {
		player.SetCurrentSpacePosition(s)
		for _, b := range player.Bullets {
			b.SetCurrentSpacePosition(s)
		}
	}


}

