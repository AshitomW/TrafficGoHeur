package main

import (
	"ashitomw/trafficgoheur/simulation"
	"time"
)

func main(){
	road := simulation.NewRoad(30,4) // 30 cells wide road , 4 lanes

 var	lights []*simulation.TrafficLight
	light := &simulation.TrafficLight{X:5,Y:0,IsGreen: true,SwitchRate: 10}
	light_t := &simulation.TrafficLight{X:15,Y:1,IsGreen: true,SwitchRate: 10}
	light_three := &simulation.TrafficLight{X:23,Y:2,IsGreen: true,SwitchRate: 10}
	light_four := &simulation.TrafficLight{X:27,Y:3,IsGreen: true,SwitchRate: 10}

	lights = append(lights, light,light_t,light_three,light_four)
	road.Grid[light.Y][light.X].Type = simulation.TrafficLightGreen


	spawner:= simulation.NewSpawnner(road,0.3)


	s:= &simulation.Simulation{
		Road:road,
		Cars: []*simulation.Car{},
		Lights: lights,
		Spawner: spawner,
		TickDelay: 400 * time.Millisecond,
	}

	s.Run(50)


}
