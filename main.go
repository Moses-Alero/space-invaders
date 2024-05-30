package main

import (
	"fmt"
	"log"

	"github.com/Moses-Alero/space-invaders/enemy"
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

var playerStartPos = models.Vector{X: float64(screenX/2) + 50, Y: float64(screenY/2) + 50}
var enemyStartPos = models.Vector{X: float64(screenX / 3), Y: float64(screenY / 3)}

var p *models.Player
var e *models.Enemy

var spaces []*models.Space
var world models.Vector = models.Vector{
	X: screenX,
	Y: screenY,
}

type Game struct {
}

func setup() {
	p = player.New()
	e = enemy.New()

	p.Position = playerStartPos
	e.Position = enemyStartPos

}

func (g *Game) Update() error {
	player.Update(p, spaces)
	enemy.Update(e, spaces)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// So basically wo i put items we want to draw on the sreen in this area and
	// the items are drawn in a desecding order as they are placed
	player.Draw(screen, p)

	utils.DrawGrid(screen, gridRowSize, gridColSize, screenY, screenX)
	enemy.Draw(screen, e)
	for _, s := range spaces {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v", len(s.Objects)), int(s.Position.X), int(s.Position.Y))
	}

	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("."), int(ship.Position.X), int(ship.Position.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingMode())
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowTitle("Space Invaders")

	spaces = CreateSpacePartition(world, gridRowSize, gridColSize)
	setup()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func CreateSpacePartition(worldSize models.Vector, row, col int) []*models.Space {
	rowSize := int(worldSize.X) / row
	colSize := int(worldSize.Y) / col
	spaces := []*models.Space{}
	size := models.Vector{
		X: float64(rowSize),
		Y: float64(colSize),
	}

	//partition
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			position := models.Vector{
				X: float64(x * rowSize),
				Y: float64(y * colSize),
			}
			space := &models.Space{
				Position: position,
				Size:     size,
				Index:    x*col + y,
				Objects:  make(map[string]*models.GameObjectModel),
			}
			spaces = append(spaces, space)
		}
	}

	return spaces
}
