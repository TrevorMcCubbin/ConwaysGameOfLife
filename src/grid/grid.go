package grid

import "github.com/TrevorMcCubbin/ConwaysGameOfLife/cell"

type Grid struct {
	Width  int
	Height int
	cells  [][]cell.Cell
}

func GenerateGrid(width int, height int) *Grid {
	// create 2d array of cells
	cells := make([][]cell.Cell, width)

	// create second dimension of cells
	for i := range cells {
		cells[i] = make([]cell.Cell, height)
	}

	// create a new grid with the values provided
	return &Grid{
		width:  width,
		height: height,
		cells:  cells,
	}
}

func GetCell(x int, y int, grid Grid) *cell.Cell {
	return &grid.Cells[x][y]
}
