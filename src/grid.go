package main

type Grid struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func GenerateGrid(width int, height int) *Grid {
	// create 2d array of cells
	cells := make([][]Cell, width)

	// create second dimension of cells
	for i := range cells {
		cells[i] = make([]Cell, height)
	}

	// create a new grid with the values provided
	return &Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}
