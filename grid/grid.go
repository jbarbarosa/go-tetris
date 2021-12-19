package grid

import (
	"fmt"
	"tetris/shape"
)

type Grid struct {
	Rows [24][10]byte
}

func InitializeNewGrid() Grid {
	grid := Grid{}
	rows := [24][10]byte{}
	for i := 0; i < 24; i++ {
		rows[i] = getEmptyRow()
	}
	grid.Rows = rows
	return grid
}

func getEmptyRow() [10]byte {
	return [10]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func GetGridState(grid *Grid) {
	fmt.Println("----------")
	for j := 0; j < len(grid.Rows); j++ {
		for i := 0; i < len(grid.Rows[j]); i++ {
			cell := grid.Rows[j][i]
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(1)
			}
		}
		fmt.Println()
	}
	fmt.Println("----------")
}

func IntroduceShape(shape *shape.Shape, grid *Grid) *Grid {
	grid.Rows[0][4] = shape.Rows[0][0]
	grid.Rows[0][5] = shape.Rows[0][1]
	grid.Rows[0][6] = shape.Rows[0][2]
	grid.Rows[0][7] = shape.Rows[0][3]

	grid.Rows[1][4] = shape.Rows[1][0]
	grid.Rows[1][5] = shape.Rows[1][1]
	grid.Rows[1][6] = shape.Rows[1][2]
	grid.Rows[1][7] = shape.Rows[1][3]

	grid.Rows[2][4] = shape.Rows[2][0]
	grid.Rows[2][5] = shape.Rows[2][1]
	grid.Rows[2][6] = shape.Rows[2][2]
	grid.Rows[2][7] = shape.Rows[2][3]

	grid.Rows[3][4] = shape.Rows[3][0]
	grid.Rows[3][5] = shape.Rows[3][1]
	grid.Rows[3][6] = shape.Rows[3][2]
	grid.Rows[3][7] = shape.Rows[3][3]
	return grid
}
