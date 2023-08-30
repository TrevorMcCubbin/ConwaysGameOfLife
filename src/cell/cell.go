package cell

type Cell struct {
	IsAlive        bool
	NextAliveState bool
}

func NewCell() *Cell {
	return &Cell{
		IsAlive:        false,
		NextAliveState: false,
	}
}
