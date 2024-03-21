package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Vector struct {
	X, Y float64
}

type GameObject interface {
	SetPosition(x float64, y float64)
	GetPosition() Vector
}

type GameObjectModel struct {
	Name     string
	Position Vector
	Sprite   *ebiten.Image
	currentWorldSpace Space
}

func (gom *GameObjectModel) SetPosition(x float64, y float64) {
	gom.Position = Vector{x, y}
}
func (gom *GameObjectModel) GetPosition() Vector {
	return gom.Position
}

func (gom *GameObjectModel) GetSize() Vector{
	return Vector{
		X: float64(gom.Sprite.Bounds().Dx()),
		Y: float64(gom.Sprite.Bounds().Dy()),
	}
}

func (gom *GameObjectModel) SpacePosition(s Space){
	overlap, space := isOverlapping(gom, s)
	if  overlap{
		gom.currentWorldSpace = *space
		space.objects[gom.Name] = gom
	}else if !overlap{
		//check if player was in space then remove player from space	
		if _, ok := space.objects[gom.Name]; ok{
			delete(space.objects, gom.Name)
		}
	}
}

func (gom *GameObjectModel) checkCollision(cb func()){
	if len(gom.currentWorldSpace.objects) < 2 {
		return
	}
	for _, object := range gom.currentWorldSpace.objects{
		if object != gom{
			if isColliding(gom, object){
				fmt.Println("Collision detected")
				cb()
			}else{
				return
			}
		} 
	} 
}

func (gom *GameObjectModel) drawBounds(){
	drawGrid(gom.Sprite, 2,2)		
}


func (gom *GameObjectModel) GetCenter() Vector {
	return Vector{
		X: gom.Position.X + float64(gom.Sprite.Bounds().Dx()),
		Y: gom.Position.Y + float64(gom.Sprite.Bounds().Dy()),
	}
}

//Player Region

var attackCoolDown time.Duration = time.Millisecond * 500

type Player struct {
	GameObjectModel
	bullet         Bullet
	bullets        []*Bullet
	health         int
	attackTimer    *Timer
	attackCoolDown time.Duration
}

func (p *Player) movement() {
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
		factor := float64(speed) / math.Sqrt(math.Pow(delta.X, 2)+math.Pow(delta.Y, 2))
		delta.X *= factor
		delta.Y *= factor
	}

	p.Position.X += delta.X
	p.Position.Y += delta.Y

}


func (p *Player) attack() {
	p.attackTimer.Update()

	if p.attackTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.attackTimer.Reset()
		p.SpawnBullet()
	}
}

func (p *Player) Controls() {
	p.movement()
	p.attack()
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.Sprite, opts)

}

func (p *Player) SpawnBullet() {
	spawnPosRight := Vector{
		X: p.Position.X + 7,
		Y: p.Position.Y + 2,
	}

	spawnPosLeft := Vector{
		X: p.Position.X - 7,
		Y: p.Position.Y - 2,
	}
	bulletR := p.bullet.Spawn(spawnPosRight)
	bulletR.Name = fmt.Sprintf("Bullet%v", len(p.bullets))
	p.bullets = append(p.bullets, &bulletR)

	bulletL := p.bullet.Spawn(spawnPosLeft)
	bulletL.Name = fmt.Sprintf("Bullet%v", len(p.bullets))
	p.bullets = append(p.bullets, &bulletL)

}

type Bullet struct {
	GameObjectModel
}

func (b *Bullet) Spawn(pos Vector) Bullet {
	b.Position = Vector{
		X: pos.X,
		Y: pos.Y,
	}

	gom := GameObjectModel{
		Position: b.Position,
		Sprite:   b.Sprite,
	}

	return Bullet{
		gom,
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(b.Position.X, b.Position.Y)
	screen.DrawImage(b.Sprite, opts)
}

func (b *Bullet) fire(speed int) {
	var delta Vector
	delta.Y -= float64(speed)
	b.Position.Y += delta.Y
}

// Enemy Region
type Enemy struct {
	GameObjectModel
	bullet Bullet
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(e.Position.X, e.Position.Y)

	screen.DrawImage(e.Sprite, opts)

}


// Barrier
type Barrier struct {
	GameObjectModel
}
