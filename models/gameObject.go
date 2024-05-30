package models

import (
	"fmt"
	//"math"
	"time"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	
)

type Space struct {
	Position Vector
	Size     Vector
	Index    int
	Objects  map[string]*GameObjectModel
}


type Vector struct {
	X, Y float64
}

type GameObject interface {
	SetPosition(x float64, y float64)
	GetPosition() Vector
}

type GameObjectModel struct {
	Name              string
	Position          Vector
	Sprite            ebiten.Image
	CurrentWorldSpace *Space
}

func (gom *GameObjectModel) SetPosition(x float64, y float64) {
	gom.Position = Vector{x, y}
}
func (gom *GameObjectModel) GetPosition() Vector {
	return gom.Position
}

func (gom *GameObjectModel) GetSize() Vector {
	return Vector{
		X: float64(gom.Sprite.Bounds().Dx()),
		Y: float64(gom.Sprite.Bounds().Dy()),
	}
}

func (gom *GameObjectModel) SetCurrentSpacePosition(s *Space) {
	overlap, space := IsOverlapping(gom, s)
	if gom.Position.Y > 599 || gom.Position.X > 599 ||
	gom.Position.Y < 0 || gom.Position.X < 0 {
		delete(gom.CurrentWorldSpace.Objects, gom.Name)
		return
	}
	if overlap {
		gom.CurrentWorldSpace = space
		space.Objects[gom.Name] = gom
	} else {
		//check if player was in space then remove player from space
		_, ok := space.Objects[gom.Name]; 
		if ok && gom.CurrentWorldSpace != space {
			delete(space.Objects, gom.Name)
			return
		}
	}
}

func (gom *GameObjectModel) CheckCollision(cb func(collider *GameObjectModel)) {
	if len(gom.CurrentWorldSpace.Objects) < 2 {
		return
	}
	for _, object := range gom.CurrentWorldSpace.Objects {
		if object.Name != gom.Name {
			if IsColliding(gom, object) {
				fmt.Println(gom.Name, " Collided with ", object.Name, " at ", time.Now().String())
				cb(object)
			}
		}
	}
}

func (gom *GameObjectModel) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(gom.Position.X, gom.Position.Y)

	//drawPosition(screen, spaces, gom, color.RGBA{0,255,0,0})

	screen.DrawImage(&gom.Sprite, opts)

}

func (gom *GameObjectModel) Destroy(){
	delete(gom.CurrentWorldSpace.Objects, gom.Name)
	gom = nil
	return
}

func (gom *GameObjectModel) GetCenter() Vector {
	return Vector{
		X: gom.Position.X + float64(gom.Sprite.Bounds().Dx()/2),
		Y: gom.Position.Y + float64(gom.Sprite.Bounds().Dy()/2),
	}
}


func (gom *GameObjectModel) Spawn(pos Vector) GameObjectModel {
	vec := Vector{ X: pos.X,Y: pos.Y }
	return GameObjectModel{
		Position: vec, 
	 	Sprite: gom.Sprite,
	}
}




func DrawPosition(screen *ebiten.Image, spaces []*Space, object *GameObjectModel, color color.Color) {
	for _, s := range spaces {
		if ok, _ := IsOverlapping(object, s); ok {
			vector.DrawFilledRect(screen, float32(s.Position.X), float32(s.Position.Y), float32(s.Size.X), float32(s.Size.Y), color, false)
		}
	}
}

func IsOverlapping(gom *GameObjectModel, s *Space) (bool, *Space) {
	objPosX, objPosY := gom.GetCenter().X, gom.GetCenter().Y
	objBoundsX, objBoundsY := gom.Sprite.Bounds().Dx(), gom.Sprite.Bounds().Dy()

	spaceRight := s.Position.X + s.Size.X
	spaceBottom := s.Position.Y + s.Size.Y

	objRight := objPosX + float64(objBoundsX)
	objBottom := objPosY + float64(objBoundsY)

	if objPosX < spaceRight && objRight > s.Position.X &&
		objBottom > s.Position.Y && objPosY < spaceBottom {
		return true, s
	}
	return false, s
}

func IsColliding(a, b *GameObjectModel) bool {
	aBoundsX, bBoundsX := a.Sprite.Bounds().Dx(), b.Sprite.Bounds().Dx()
	aBoundsY, bBoundsY := a.Sprite.Bounds().Dy(), b.Sprite.Bounds().Dy()

	return a.Position.X < (b.Position.X+float64(bBoundsX)) &&
		b.Position.X < (a.Position.X+float64(aBoundsX)) &&
		a.Position.Y < (b.Position.X+float64(bBoundsY)) &&
		b.Position.Y < (a.Position.Y+float64(aBoundsY))

}
