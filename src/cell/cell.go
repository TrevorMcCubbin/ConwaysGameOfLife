package cell

type Cell struct {
	IsAlive        bool
	nextAliveState bool
}

func NewCell() *Cell {
	return &Cell{
		IsAlive:        false,
		nextAliveState: false,
	}
}
