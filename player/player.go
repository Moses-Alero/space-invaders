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
	player := new(models.Player)
	player.Health = 100
	player.AttackTimer = utils.NewTimer(AttackCoolDown)
	player.Bullet.Sprite = assets.PlayerBulletSprite
	player.Name = "Player"
	player.Sprite = assets.PlayerSprite

	return player
}


func Update(player *models.Player, s []*models.Space){
	
	setPlayerWorldPos(s, player)
	player.CheckCollision(func(gom *models.GameObjectModel) {

	})

	if player.Position.X > (600 - player.GetSize().X * 2) {
		player.Position.X = 600 - player.GetSize().X * 2
	}
	if player.Position.X <= 0 {
		player.Position.X = 0
	}
	if player.Position.Y > (600 - player.GetSize().Y * 2) {
		player.Position.Y = 600 - player.GetSize().Y * 2
	}
	if player.Position.Y <= 0 {
		player.Position.Y = 0
	}


	for i, b := range player.Bullets {
		b.Fire()
		b.CheckCollision(func(gom *models.GameObjectModel) {
			if len(player.Bullets) < 1 {
				return
			} 
			player.Bullets = append(player.Bullets[:i], player.Bullets[i+1:]...)
		})
		if b != nil && b.Position.Y <= -10 {
			player.Bullets = append(player.Bullets[:i], player.Bullets[i+1:]...)
		}
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

