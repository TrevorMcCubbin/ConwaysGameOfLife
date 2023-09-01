package grid

import "math/rand"

type Grid [][]bool

func GenerateGrid(width int, height int) *Grid {
	// create 2d array of cells
	newGrid := make(Grid, height)

	// create second dimension of cells
	for i := range newGrid {
		newGrid[i] = make([]bool, width)
	}

	return &newGrid
}

// Create a random grid, should be used
func (grid Grid) Seed() {
	for _, row := range grid {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
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
