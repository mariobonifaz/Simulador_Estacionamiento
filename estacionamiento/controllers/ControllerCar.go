package controllers

import (
	"Mario/models"
	"Mario/views"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ControllerCar struct {
	model *models.Car
	view  *views.ViewCar
	mu    *sync.Mutex
}

func NewControllerCar(win *pixelgl.Window, mu *sync.Mutex) *ControllerCar {
	return &ControllerCar{
		model: models.NewCar(),
		view:  views.NewViewCar(win),
		mu:    mu,
	}
}

func (cc *ControllerCar) GenerateCars(n int, chCar *chan models.Car) {
	cc.model.GenerateCars(n, *chCar)
}

func (cc *ControllerCar) LoadSprite() {
	cc.view.SetSprite()
}

func (cc *ControllerCar) PaintCar(pos pixel.Vec) {
	cc.view.PaintCar(pos)
}