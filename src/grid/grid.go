package grid

import (
	"fmt"
	"math/rand"
)

const (
	aliveCell = "\xF0\x9F\x9F\xA9"
	deadCell  = "\xF0\x9F\x9F\xAB"
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
			if rand.Intn(grid.height/3) == 1 {
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

func (grid Grid) GetNeighbors(x, y int) int {
	var neighbors int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if grid.CellIsAlive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (grid Grid) CellIsAlive(x, y int) bool {
	y = (grid.height + y) % grid.height
	x = (grid.width + x) % grid.width
	return grid.gameGrid[y][x]
}

func (grid Grid) UpdateCell(x, y int) bool {
	neighbours := grid.GetNeighbors(x, y)
	isAlive := grid.CellIsAlive(x, y)

	if neighbours < 4 && neighbours > 1 && isAlive {
		return true
	} else if neighbours == 3 && !isAlive {
		return true
	} else {
		return false
	}
}

func UpdateGrid(current Grid) {
	for i := 0; i < current.height; i++ {
		for j := 0; j < current.width; j++ {
			current.gameGrid[i][j] = current.UpdateCell(j, i)
		}
	}
}
