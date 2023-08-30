package main

import "fmt"

func main() {
	grid := GenerateGrid(10, 10)

	fmt.Println(GetCell(3, 3, *grid))
}
