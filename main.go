package main

import (
	"ashitomw/trafficgoheur/simulation"
	"time"
)

func main(){
	road := simulation.NewRoad(20,10)
	
	c1 := &simulation.Car{X: 0, Y: 5, Direction: simulation.East}
	c2 := &simulation.Car{X: 2, Y: 5, Direction: simulation.East}
	c3 := &simulation.Car{X: 10, Y: 2, Direction: simulation.South}


	light := &simulation.TrafficLight{X:8,Y:5,IsGreen: true, SwitchRate: 5}


	road.Grid[c1.Y][c1.X].Type = simulation.CarCell
	road.Grid[c2.Y][c2.X].Type = simulation.CarCell
	road.Grid[c3.Y][c3.X].Type = simulation.CarCell

	road.Grid[light.Y][light.X].Type = simulation.TrafficLightGreen


	s := &simulation.Simulation{
		Road: road,
		Cars: []*simulation.Car{c1,c2,c3},
		Lights: []*simulation.TrafficLight{light},
		TickDelay: 500 * time.Millisecond,
	}

	s.Run(30)

}
