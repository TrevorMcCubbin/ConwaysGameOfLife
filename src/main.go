package main

import (
	"fmt"

	"github.com/TrevorMcCubbin/ConwaysGameOfLife/grid"
)

func main() {
	grid := grid.GenerateGrid(10, 10)

	fmt.Println(grid.GetCell(3, 3, *grid))
}
