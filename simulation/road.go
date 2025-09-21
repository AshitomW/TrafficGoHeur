package simulation


type CellType int

const (
	Empty CellType = iota
	CarCell
	TrafficLightRed
	TrafficLightGreen
)


type Cell struct{
	Type CellType
}


type Road struct{
	Width int
	Height int 
	Grid [][]*Cell
}


func NewRoad(width,height int) *Road{
	grid := make([][] *Cell,height)
	for y:=range height{
		grid[y] = make([]*Cell,width)
		for x:=range width{
			grid[y][x] = &Cell{Type:Empty}
		}
	}
	return &Road{Width: width,Height: height,Grid: grid}
}
