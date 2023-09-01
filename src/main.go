package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TrevorMcCubbin/ConwaysGameOfLife/grid"
)

const (
	width          = 50
	height         = 25
	sleepIteration = 20
	ansiEscapeSeq  = "\033c\x0c"
)

func main() {
	clearScreen()
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	currentGrid := grid.GenerateGrid(width, height)
	currentGrid.Seed()
	for {
		currentGrid.Display()
		grid.UpdateGrid(*currentGrid)
		time.Sleep(sleepIteration * time.Millisecond)
		clearScreen()
	}
}

func clearScreen() {
	fmt.Printf(ansiEscapeSeq)
}
