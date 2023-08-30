package main

import (
	"fmt"

	"github.com/TrevorMcCubbin/ConwaysGameOfLife/grid"
)

func main() {
	gameGrid := grid.GenerateGrid(10, 10)

	fmt.Println(grid.GetCell(3, 3, *gameGrid))
}
