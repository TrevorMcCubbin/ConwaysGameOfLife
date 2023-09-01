package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TrevorMcCubbin/ConwaysGameOfLife/grid"
)

const (
	width          = 40
	height         = 20
	sleepIteration = 1000
	ansiEscapeSeq  = "\033c\x0c"
)

func main() {
	clearScreen()
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	currentGrid := grid.GenerateGrid(width, height)
	nextGrid := grid.GenerateGrid(width, height)
	currentGrid.Seed()
	for {
		currentGrid.Display()
		grid.UpdateGrid(*currentGrid, *nextGrid)
		time.Sleep(sleepIteration * time.Millisecond)
		clearScreen()
	}
}

func clearScreen() {
	fmt.Printf(ansiEscapeSeq)
}
