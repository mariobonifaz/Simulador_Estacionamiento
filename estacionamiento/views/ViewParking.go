package views

import (
	"Mario/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewParking struct {
	win            *pixelgl.Window
	utils          *utils.Utils
	spaces         [20]pixel.Vec
	questionSpaces [20]pixel.Vec
}

func NewViewParking(win *pixelgl.Window) *ViewParking {
	return &ViewParking{
		win: win,
		spaces: [20]pixel.Vec{
			pixel.V(935, 100), pixel.V(935, 245), pixel.V(935, 365), pixel.V(935, 625), pixel.V(935, 725),
			pixel.V(700, 725), pixel.V(700, 625), pixel.V(700, 365), pixel.V(700, 245), pixel.V(700, 100),
			pixel.V(85, 725), pixel.V(85, 625), pixel.V(85, 365), pixel.V(85, 245), pixel.V(85, 100),
			pixel.V(300, 100), pixel.V(300, 245), pixel.V(300, 365), pixel.V(300, 625), pixel.V(300, 725),
		},
	}
}

func (pw *ViewParking) PaintParking() {
	picParking, err := pw.utils.LoadPicture("./assets/Ecenario.png")
	if err != nil {
		panic(err)
	}

	parking := pw.utils.NewSprite(picParking, picParking.Bounds())

	matrix := pixel.IM
	matrix = pixel.IM.Moved(pixel.V(512, 469))
	parking.Draw(pw.win, matrix)
}

func (pw *ViewParking) PaintStreet() {
	picStreet, err := pw.utils.LoadPicture("./assets/Ecenario.png")
	if err != nil {
		panic(err)
	}

	street := pw.utils.NewSprite(picStreet, picStreet.Bounds())

	street.Draw(pw.win, pixel.IM.Moved(pixel.V(512, 85)))
}

func (pw *ViewParking) GetCoordinates(n int) pixel.Vec {
	return pw.spaces[n]
}
