package views

import (
	"Mario/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewEntrada struct {
	win    *pixelgl.Window
	utils  *utils.Utils
	states [1]pixel.Sprite
}

func NewViewEntrada(win *pixelgl.Window) *ViewEntrada {
	return &ViewEntrada{
		win:   win,
		
	}
}

func (ev *ViewEntrada) LoadStatesImages() [1]pixel.Sprite {
	openEntrance := ev.loadEntranceSprites()
	return [1]pixel.Sprite{openEntrance}
}

func (ev *ViewEntrada) loadEntranceSprites() (pixel.Sprite) {
	picEntranceOpen, _ := ev.utils.LoadPicture("./assets/carro.png")

	openEntrance := ev.utils.NewSprite(picEntranceOpen, picEntranceOpen.Bounds())

	return *openEntrance
}

func (ev *ViewEntrada) SetStateImages(imgs [1]pixel.Sprite) {
	ev.states = imgs
}

func (ev *ViewEntrada) PaintEntrance(img int) {
	entrancePos := pixel.V(920, 200)
	ev.states[img].Draw(ev.win, pixel.IM.Moved(entrancePos))
}
