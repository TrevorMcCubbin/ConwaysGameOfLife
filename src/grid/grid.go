package grid

import (
	"fmt"
	"math/rand"
)

const (
	aliveCell = "1"
	deadCell  = "0"
)

type Grid struct {
	gameGrid [][]bool
	width    int
	height   int
}

func GenerateGrid(width int, height int) *Grid {
	// create 2d array of cells
	newGrid := make([][]bool, height)

	// create second dimension of cells
	for i := range newGrid {
		newGrid[i] = make([]bool, width)
	}

	return &Grid{
		gameGrid: newGrid,
		width:    width,
		height:   height,
	}
}

// Create a random grid, should be used
func (grid Grid) Seed() {
	for _, row := range grid.gameGrid {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (grid Grid) Display() {
	for _, row := range grid.gameGrid {
		for _, cell := range row {
			if cell {
				fmt.Printf(aliveCell)
			} else {
				fmt.Printf(deadCell)
			}
		}
		fmt.Printf("\n")
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
