package controllers

import (
	"Mario/models"
	"Mario/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrada
	view  *views.ViewEntrada
	mu    *sync.Mutex
}

func NewEntranceController(win *pixelgl.Window, mu *sync.Mutex) *EntranceController {
	return &EntranceController{
		model: models.NewEntrada(),
		view:  views.NewViewEntrada(win),
		mu:    mu,
	}
}

func (ec *EntranceController) LoadStates() {
	images := ec.view.LoadStatesImages()
	ec.view.SetStateImages(images)
}

func (ec *EntranceController) PaintEntrance(position int) {
	ec.view.PaintEntrance(position)
}
