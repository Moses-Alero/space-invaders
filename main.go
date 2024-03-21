package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenX     = 600
	screenY     = 600
	gridRowSize = 5
	gridColSize = 6
)

var playerStartPos = Vector{float64(screenX / 2) + 50, float64(screenY / 2) + 50}
var enemyStartPos = Vector{float64(screenX / 2), float64(screenY / 2)}

var ship Player
var enemy Enemy
var speed int = 200 / ebiten.TPS()
var pBulletSpeed int = 300 / ebiten.TPS()
var spaces []Space
var world Vector = Vector{
	X: screenX,
	Y: screenY,
}
type Game struct {
}

type Space struct {
	Position Vector
	Size     Vector
	index    int
	objects  map[string]*GameObjectModel
}

func (g *Game) Update() error {
	for _, b := range ship.bullets {
		b.fire(pBulletSpeed)
	}
	for _, s := range spaces {
		ship.SpacePosition(s)
		enemy.SpacePosition(s)
		for _, b := range ship.bullets{
			b.SpacePosition(s)
		}
	}
	ship.Controls()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// So basically wo i put items we want to draw on the sreen in this area and
	// the items are drawn in a desecding order as they are placed
//	drawGrid(screen, gridRowSize, gridColSize)
	for _, b := range ship.bullets {
		b.Draw(screen)
	}

//	for _, s := range spaces {
//		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("POS: %v, %v", s.Position.X, s.Position.Y), int(s.Position.X), int(s.Position.Y))
//	}

	//	vector.StrokeLine(screen, float32(ship.center().X), float32(ship.center().Y), float32(ship.center().X + 4), float32(ship.center().Y), 4, color.White, false)

//	drawPosition(screen, spaces, &ship.GameObjectModel)
	ship.checkCollision(ship.drawBounds)
	ship.Draw(screen)
	enemy.Draw(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("POS: %v, %v", ship.GetCenter().X, ship.GetCenter().Y), 350, 350)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingMode() + 1)
	ebiten.SetWindowTitle("Space Invaders")

	spaces = setSpacePartition(world, gridRowSize, gridColSize)

	ship.Sprite = PlayerSprite
	ship.Name = "Player"
	ship.attackTimer = NewTimer(attackCoolDown)
	ship.SetPosition(playerStartPos.X, playerStartPos.Y)
	ship.bullet.Sprite = PlayerBulletSprite

	enemy.Sprite = EnemySprite
	enemy.Name = "Enemy"
	enemy.SetPosition(enemyStartPos.X, enemyStartPos.Y)
	enemy.bullet.Sprite = PlayerBulletSprite


	

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func drawGrid(dst *ebiten.Image, rowSize, colSize int) {
	row := dst.Bounds().Dx() / rowSize
	col := dst.Bounds().Dy() / colSize
	for i := 1; i < rowSize; i++ {
		ptX0 := float32(row * i)
		ptY0 := float32(0)
		ptX1 := ptX0
		ptY1 := float32(screenY)
		vector.StrokeLine(dst, ptX0, ptY0, ptX1, ptY1, 1, color.White, false)
	}

	for i := 1; i < colSize; i++ {
		ptX0 := float32(0)
		ptY0 := float32(col * i)
		ptX1 := float32(screenX)
		ptY1 := ptY0
		vector.StrokeLine(dst, ptX0, ptY0, ptX1, ptY1, 1, color.White, false)
	}
}

func setSpacePartition(worldSize Vector, row, col int) []Space {
	rowSize := int(worldSize.X) / row
	colSize := int(worldSize.Y) / col
	spaces := make([]Space, row*col)
	size := Vector{
		X: float64(rowSize),
		Y: float64(colSize),
	}

	//partition
	for i := 0; i < len(spaces); i++ {
		for x := 0; x < row; x++ {
			for y := 0; y < col; y++ {
				position := Vector{
					X: float64(x * rowSize),
					Y: float64(y * colSize),
				}
				spaces[i] = Space{
					Position: position,
					Size:     size,
					index:    i,
					objects: make(map[string]*GameObjectModel),
				}
				i++
			}
		}
	}

	return spaces
}

func drawPosition(screen *ebiten.Image, spaces []Space, object *GameObjectModel) {
	for _, s := range spaces {
		if ok, _ := isOverlapping(object, s); ok{
			vector.DrawFilledRect(screen, float32(s.Position.X), float32(s.Position.Y), float32(s.Size.X), float32(s.Size.Y), color.White, false)
		}
	}
}

//func setObjectsInSpace(spaces []Space)

func isOverlapping(gom *GameObjectModel, s Space ) (bool,  *Space){
	if gom.GetCenter().X <= (s.Position.X+s.Size.X) && gom.GetCenter().X >= s.Position.X || 
	 	(gom.GetCenter().X + float64(gom.Sprite.Bounds().Dx())) <= (s.Position.X + s.Size.X) && (gom.GetCenter().X + float64(gom.Sprite.Bounds().Dx()) >= s.Position.X){
			if gom.GetCenter().Y  >= s.Position.Y && gom.GetCenter().Y  <= (s.Position.Y+s.Size.Y) ||  
				(gom.GetCenter().Y + float64(gom.Sprite.Bounds().Dy())) <= (s.Position.Y + s.Size.Y) && (gom.GetCenter().Y + float64(gom.Sprite.Bounds().Dy()) >= s.Position.Y){
					return true, &s
				}
	
	}
	return false, &s
}

func isColliding(a, b *GameObjectModel) bool{
	if (a.Position.X + float64(a.Sprite.Bounds().Dx())) <= (b.Position.X + float64(b.Sprite.Bounds().Dx())) || 
	 	(a.Position.X + float64(a.Sprite.Bounds().Dx())) <= (b.Position.X + float64(b.Sprite.Bounds().Dx())) && (a.GetCenter().X + float64(a.Sprite.Bounds().Dx()) >= b.Position.X){
			if a.Position.Y  >= b.Position.Y && a.Position.Y  <= (b.Position.Y+ float64(b.Sprite.Bounds().Dy())) ||  
				(a.Position.Y + float64(a.Sprite.Bounds().Dy())) <= (b.Position.Y + float64(b.Sprite.Bounds().Dy())) && (a.Position.Y + float64(a.Sprite.Bounds().Dy()) >= b.Position.Y){
					return true
				}
	
	}
	return false

}


