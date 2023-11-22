package models

import (
	"Mario/utils"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/faiface/pixel"
)

type Car struct {
	ParkingTime int
	Id          int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano())
	parkingTime := rand.Intn(10) + 15
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch <- *car
		randomSleep(1, 2)
	}
	close(ch)
}

func (c *Car) Timer(pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	enterParkingLot(pos, pc, mu, chEntrance, sprite, chWin, coo)
	parkCar(c, pos, pc, mu, spaces, chEntrance, sprite, chWin, coo)
	leaveParkingLot(c, pos, pc, mu, spaces, chEntrance, sprite, chWin, coo)
}

func randomSleep(minSeconds, maxSeconds int) {
	rand.Seed(time.Now().UnixNano())
	newTime := rand.Intn(maxSeconds-minSeconds+1) + minSeconds
	time.Sleep(time.Second * time.Duration(newTime))
}

func enterParkingLot(pos int, pc *Parking, mu *sync.Mutex, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	mu.Lock()
	data := utils.NewImgCar(sprite, pos, true, coo)
	chWin <- *data
	*chEntrance <- 0
	mu.Unlock()
}

func parkCar(c *Car, pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	mu.Lock()
	pc.nSpaces--
	fmt.Printf("Generados. %d: %d\n", c.Id, pos)
	fmt.Printf("Disponibles: %d\n", pc.nSpaces)
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))
}

func leaveParkingLot(c *Car, pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	fmt.Printf("Salida: %d\n", c.Id)

	mu.Lock()
	data := utils.NewImgCar(sprite, pos, false, coo)
	chWin <- *data
	pc.nSpaces++
	spaces[pos] = true
	fmt.Printf("Disponibles: %d\n", pc.nSpaces)
	mu.Unlock()

	mu.Lock()
	*chEntrance <- 1
	mu.Unlock()
}
