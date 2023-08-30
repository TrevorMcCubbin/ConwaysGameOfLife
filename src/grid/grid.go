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
		Width:  width,
		Height: height,
		cells:  cells,
	}
}

func GetCell(x int, y int, grid Grid) *cell.Cell {
	return &grid.cells[x][y]
}

func Update(xPosition int, yPosition int, gameGrid Grid) {
	cellsAlive := 0

	// loop through all neighbours
	for x := xPosition - 1; x <= xPosition+1; x++ {
		for y := yPosition - 1; y <= yPosition+1; y++ {
			// skip the cell you are currently checking
			if x == xPosition && y == yPosition {
				continue
			}

			if GetCell(x, y, gameGrid).IsAlive {
				cellsAlive++
			}
		}
	}

	currentCell := GetCell(xPosition, yPosition, gameGrid)

	currentCell.NextAliveState = currentCell.IsAlive

	if currentCell.IsAlive && cellsAlive == 3 {
		currentCell.NextAliveState = true
	} else {
		currentCell.NextAliveState = false
	}

}
