package simulation
/*
	Cars:
		Position : x, y
		Direction : North , South , East , West
*/

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Car struct {
	X,Y int
	Direction Direction
}


func (c *Car) Move()(int , int){
	switch c.Direction{
		case North:
			return c.X,c.Y-1
	case South:
			return c.X,c.Y+1
	case East:
			return c.X+1,c.Y
	case West:
			return c.X-1,c.Y
	}
		return c.X,c.Y
}
