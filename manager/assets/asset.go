package assets

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed "all:game-assets" 
var asset embed.FS

var PlayerSprite = LoadAsset("game-assets/Player/Player_Ship_Ant_00.png")
var PlayerBulletSprite = LoadAsset("game-assets/Player/Player_Shot_27.png")
var EnemySprite = LoadAsset("game-assets/Enemy/Enemy_Ship_Scout_21.png")

func LoadAsset(path string) *ebiten.Image {
	//read the file
	f, err := asset.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
