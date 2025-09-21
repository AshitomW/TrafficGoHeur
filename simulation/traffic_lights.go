package simulation


type TrafficLight struct{
	X,Y int
	IsGreen bool
	SwitchRate int // time / ticks before switching lights
	Ticks int 
}

func (l *TrafficLight) Update(){
	l.Ticks++
	if l.Ticks >= l.SwitchRate {
		l.IsGreen = !l.IsGreen
		l.Ticks = 0;
	}
}
