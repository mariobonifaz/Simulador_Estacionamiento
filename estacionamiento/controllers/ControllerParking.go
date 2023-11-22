package controllers

import (
	"Mario/models"
	"Mario/utils"
	"Mario/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type ControllerParking struct {
	model *models.Parking
	view  *views.ViewParking
	mu    *sync.Mutex
}

func NewControllerParking(win *pixelgl.Window, mu *sync.Mutex) *ControllerParking {
	model := models.NewParking()
	view := views.NewViewParking(win)
	return &ControllerParking{
		model: model,
		view:  view,
		mu:    mu,
	}
}

func (pc *ControllerParking) PaintParking() {
	pc.view.PaintParking()
}

func (pc *ControllerParking) PaintStreet() {
	pc.view.PaintStreet()
}

func (pc *ControllerParking) Park(carChannel *chan models.Car, entranceController *EntranceController, carController *ControllerCar, entranceChannel *chan int, winChannel chan utils.ImgCar) {
	go pc.ChangingState(entranceChannel, entranceController)

	for car := range *carChannel {
		position := pc.model.FindSpaces()
		if position != -1 {
			coordinates := pc.view.GetCoordinates(position)
			carController.view.SetSprite()
			sprite := carController.view.PaintCar(coordinates)

			state := entranceController.model.GetState()
			if state == "Parado" || state == "Entrando" {
				go car.Timer(position, pc.model, pc.mu, pc.model.GetAllSpaces(), entranceChannel, sprite, winChannel, coordinates)
			} else {
				*entranceChannel <- 0
				go car.Timer(position, pc.model, pc.mu, pc.model.GetAllSpaces(), entranceChannel, sprite, winChannel, coordinates)
			}
		}
	}
}

func (pc *ControllerParking) ChangingState(entranceChannel *chan int, entranceController *EntranceController) {
	for change := range *entranceChannel {
		entranceController.model.SetState(change)
	}
}
