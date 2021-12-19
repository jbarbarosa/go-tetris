package main

import (
	"fmt"
	"tetris/grid"
	"tetris/shape"
)

func main() {
	newShape := shape.O()
	newGrid := grid.InitializeNewGrid()
	grid.GetGridState(&newGrid)
	fmt.Println()
	grid.IntroduceShape(&newShape, &newGrid)
	grid.GetGridState(&newGrid)
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )
// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Simple Shell")
// 	fmt.Println("---------------------")

// 	for {
// 		fmt.Print("-> ")
// 		text, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		text = strings.Replace(text, "\n", "", -1)
// 		fmt.Println(text)
// 	}

// }
