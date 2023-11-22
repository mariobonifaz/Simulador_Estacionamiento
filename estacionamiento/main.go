package main

import (
    "github.com/faiface/pixel/pixelgl"
    "Mario/Simulation"
)

func main() {
    pixelgl.Run(func() {
        sim := simulation.NewSimulation()
        sim.Init()
        sim.Run()
    })
}
