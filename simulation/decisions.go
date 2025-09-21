package simulation


// Should the car move ? 

func ShouldMove(car *Car, road *Road) bool {
	nextX, nextY := car.Move()

	if nextX < 0 || nextY <0 || nextX >= road.Width || nextY >= road.Height {
		return false 
	}

	nextCell := road.Grid[nextY][nextX]


	if nextCell.Type == CarCell{
		return false;
	}

	if nextCell.Type == TrafficLightRed {
		return false
	}

	// avoiding collisions -> look 2 steps ahead
	if nextCell.Type == Empty{
		farX, farY := nextX, nextY 
		switch car.Direction{
		case North:
			farY--
		case South:
			farY++ 
		case East:
			farX++
		case West:
			farX--
		}

		if farX >=0 && farY >=0  && farX <road.Width && farY < road.Height {
			if road.Grid[farY][farX].Type == CarCell{
				return false
			}
		}

	}


	return true

}
