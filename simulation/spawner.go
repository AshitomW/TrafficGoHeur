package simulation

import (
	"math/rand"
	"time"
)

type Spawner struct{
	SpawnRate float64 
	Road *Road
}


func NewSpawnner(road *Road, rate float64) *Spawner{
	rand.Seed(time.Now().UnixNano())
	return &Spawner{Road:road, SpawnRate: rate}
}


func (s *Spawner) TrySpawn(cars *[]*Car){
	
	for y:=0;y<s.Road.Height;y++{
		if rand.Float64() < s.SpawnRate{
			if s.Road.Grid[y][0].Type == Empty{
				newCar:=&Car{X:0,Y:y,Direction: East}
				*cars = append(*cars, newCar)
				s.Road.Grid[y][0].Type = CarCell
			}
		}
	}

}
