package simulation

import (
	_ "image/png"
	"Mario/controllers"
	"Mario/models"
	"Mario/utils"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Simulation struct {
	win             *pixelgl.Window
	carChannel      chan models.Car
	entranceChannel chan int
	winChannel      chan utils.ImgCar
	mu              *sync.Mutex
	parkingCtrl     *controllers.ControllerParking
	entranceCtrl    *controllers.EntranceController
	carCtrl         *controllers.ControllerCar
	carSprites      []utils.ImgCar
}

func NewSimulation() *Simulation {
	cfg := pixelgl.WindowConfig{
		Title:  "Estacionamiento",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	carChannel := make(chan models.Car, 100)
	entranceChannel := make(chan int)
	winChannel := make(chan utils.ImgCar)
	mu := &sync.Mutex{}

	return &Simulation{
		win:             win,
		carChannel:      carChannel,
		entranceChannel: entranceChannel,
		winChannel:      winChannel,
		mu:              mu,
		parkingCtrl:     controllers.NewControllerParking(win, mu),
		entranceCtrl:    controllers.NewEntranceController(win, mu),
		carCtrl:         controllers.NewControllerCar(win, mu),
	}
}

func (s *Simulation) Init() {
	s.carCtrl.LoadSprite()
	s.entranceCtrl.LoadStates()
}

func (s *Simulation) Run() {
	go s.parkingCtrl.Park(&s.carChannel, s.entranceCtrl, s.carCtrl, &s.entranceChannel, s.winChannel)
	go s.carCtrl.GenerateCars(100, &s.carChannel)

	for !s.win.Closed() {
		s.updateAndDraw()
	}
}

func (s *Simulation) updateAndDraw() {
	s.win.Clear(colornames.Black)
	s.parkingCtrl.PaintParking()
	s.parkingCtrl.PaintStreet()
	s.handleWinChannel()

	for _, value := range s.carSprites {
		s.drawCarSprite(value)
	}

	s.win.Update()
}

func (s *Simulation) handleWinChannel() {
	select {
	case val := <-s.winChannel:
		if val.IsEntering() {
			s.carSprites = append(s.carSprites, val)
		} else {
			s.removeCarSprite(val)
		}
	}
}

func (s *Simulation) removeCarSprite(val utils.ImgCar) {
	var arrAux []utils.ImgCar
	for _, value := range s.carSprites {
		if value.GetID() != val.GetID() {
			arrAux = append(arrAux, value)
		}
	}
	s.carSprites = s.carSprites[:0]
	s.carSprites = append(s.carSprites, arrAux...)
}

func (s *Simulation) drawCarSprite(value utils.ImgCar) {
	sprite := value.GetSprite()
	pos := value.GetPosition()
	sprite.Draw(s.win, pixel.IM.Moved(pos))
}