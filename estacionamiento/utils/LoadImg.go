package utils

import (
	"image"
	"os"
	"github.com/faiface/pixel"
)

type Utils struct{}

type ImgCar struct {
	sprite     *pixel.Sprite
	ID         int
	entering   bool
	position   pixel.Vec
}

func NewImgCar(sprite *pixel.Sprite, ID int, state bool, position pixel.Vec) *ImgCar {
	return &ImgCar{
		sprite:   sprite,
		ID:       ID,
		entering: state,
		position: position,
	}
}

func (ic *ImgCar) GetSprite() *pixel.Sprite {
	return ic.sprite
}

func (ic *ImgCar) GetPosition() pixel.Vec {
	return ic.position
}

func (ic *ImgCar) GetID() int {
	return ic.ID
}

func (ic *ImgCar) IsEntering() bool {
	return ic.entering
}

func (ic *ImgCar) GetData() *ImgCar {
	return ic
}

func (u *Utils) LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (u *Utils) NewSprite(picture pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(picture, form)
}
