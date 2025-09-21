package simulation


import (
	"fmt"
	"time"
)

type Simulation struct{
	Road *Road
	Cars []*Car
	Lights []*TrafficLight
	TickDelay time.Duration
}

func (s *Simulation) Step(){
	for _, l:= range s.Lights{
		l.Update()
		if l.IsGreen{
			s.Road.Grid[l.Y][l.X].Type = TrafficLightGreen
		}else{
			s.Road.Grid[l.Y][l.X].Type = TrafficLightRed
		}
	}

	// move the cars 

	for _,c := range s.Cars{
		if ShouldMove(c,s.Road){
			s.Road.Grid[c.Y][c.X].Type = Empty // clear the old position
			c.X,c.Y = c.Move()
		}
		s.Road.Grid[c.Y][c.X].Type = CarCell
	}
}

func (s *Simulation) Render(){

	for y:=range s.Road.Height{
		for x:= range s.Road.Width{
			switch s.Road.Grid[y][x].Type{
			case Empty:
				fmt.Print(".")
			case CarCell:
				fmt.Print("C")
			case TrafficLightRed:
				fmt.Print("R")
			case TrafficLightGreen:
				fmt.Print("G")
			}
		}	
		fmt.Println()
	}
	fmt.Println()


}


func (s *Simulation) Run(steps int){
	for range steps{
		s.Step()
		s.Render()
		time.Sleep(s.TickDelay)
	}
}

