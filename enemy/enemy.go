package enemy

import (
	"fmt"
	"strings"
	"time"

	"github.com/Moses-Alero/space-invaders/manager/assets"
	"github.com/Moses-Alero/space-invaders/models"
	"github.com/Moses-Alero/space-invaders/utils"
	"github.com/hajimehoshi/ebiten/v2"
)
var AttackCoolDown time.Duration = time.Millisecond * 500


func New() *models.Enemy{
	enemy := models.Enemy {
		Health: 100,
		AttackTimer: utils.NewTimer(AttackCoolDown),
	}
	enemy.Bullet.Sprite = assets.PlayerBulletSprite
	enemy.Sprite = assets.EnemySprite
	enemy.Name = "Enemy"

	return &enemy
}

func Update( e *models.Enemy, s []*models.Space){
	setWorldPos(s, e)

	e.CheckCollision(func(gom *models.GameObjectModel){
		if strings.Contains(gom.Name, "Bullet"){
			e.Health -= 50
			gom.Destroy()
			fmt.Println(e.Health)
			if e.Health < 1{
				e.Destroy()
				return
			}
		}
	})
}


func Draw(screen *ebiten.Image, enemy *models.Enemy){
	enemy.Draw(screen)	
}


func setWorldPos(spaces []*models.Space, enemy *models.Enemy){
	for _, s := range spaces {
		enemy.SetCurrentSpacePosition(s)
		for _, b := range enemy.Bullets {
			b.SetCurrentSpacePosition(s)
		}
	}
}

