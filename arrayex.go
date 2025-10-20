package main

import (
	"fmt"
	// "strings"
	"golang.org/x/tour/pic"
)

func main() {
	// fmt.Println("Array and Slice Example")
	// board := [][]string{
	// 	[]string{"_", "_", "_", "_"},
	// 	[]string{"_", "_", "_", "_"},
	// 	[]string{"_", "_", "_", "_"},
	// }
	// board[0][3] = "X"
	// board[1][1] = "X"
	// board[2][2] = "0"

	// fmt.Printf("Current Board length: %d\n", len(board))
	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], "zonzur"))

	// }
	fmt.Println("pic example")
	pic.Show(Pic)

}

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		arr[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			arr[i][j] = uint8((i + j) / 2)
		}
	}
	return arr

}
