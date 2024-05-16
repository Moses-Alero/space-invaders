package utils

import(
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

)


func DrawGrid(dst *ebiten.Image, rowSize, colSize, screenY, screenX int) {
	row := dst.Bounds().Dx() / rowSize
	col := dst.Bounds().Dy() / colSize
	for i := 1; i < rowSize; i++ {
		ptX0 := float32(row * i)
		ptY0 := float32(0)
		ptX1 := ptX0
		ptY1 := float32(screenY)
		vector.StrokeLine(dst, ptX0, ptY0, ptX1, ptY1, 1, color.RGBA{0, 255, 5, 5}, false)
	}

	for i := 1; i < colSize; i++ {
		ptX0 := float32(0)
		ptY0 := float32(col * i)
		ptX1 := float32(screenX)
		ptY1 := ptY0
		vector.StrokeLine(dst, ptX0, ptY0, ptX1, ptY1, 1, color.RGBA{255, 0, 5, 5}, false)
	}
}


