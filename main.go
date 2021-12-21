package main

import (
	"fmt"
	"tetris/game"
	"tetris/grid"
	"tetris/shape"
)

func main() {
	game.Start()
	grid.CallClear()
	newShape := shape.O()
	newGrid := grid.InitializeNewGrid()
	grid.GetGridState(&newGrid)
	fmt.Println()
	grid.IntroduceShape(&newShape, &newGrid)
	grid.GetGridState(&newGrid)
}
