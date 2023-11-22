package views

import (
	"Mario/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewCar struct {
	win    *pixelgl.Window
	utils  *utils.Utils
	sprite *pixel.Sprite
}

type SpriteCar struct {
	img *pixel.Sprite
	Id  int
}

func NewViewCar(win *pixelgl.Window) *ViewCar {
	return &ViewCar{
		win:   win,
		
	}
}

func (cv *ViewCar) SetSprite() {
	carSprite := cv.loadCarSprite()
	cv.sprite = carSprite
}

func (cv *ViewCar) PaintCar(pos pixel.Vec) *pixel.Sprite {
	cv.sprite.Draw(cv.win, pixel.IM.Moved(pos))
	return cv.sprite
}

func (cv *ViewCar) loadCarSprite() *pixel.Sprite {
	picCar, _ := cv.utils.LoadPicture("./assets/carro.png")
	return cv.utils.NewSprite(picCar, picCar.Bounds())
}

func NewImgCar(spr *pixel.Sprite, Id int) *SpriteCar {
	return &SpriteCar{
		img: spr,
		Id:  Id,
	}
}
