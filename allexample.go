package main

import (
	"fmt"
)

// func adder() func(int) int {
// 	sum := 0

// 	return func(x int) int {

// 		sum += x
// 		return sum
// 	}

// }

// func adder(m int) []int {

// 	s := make([]int, 0)
// 	for i := 0; i < m; i++ {

// 		s = append(s, i)
// 		fmt.Println("append the slice", s)

// 	}

// 	fmt.Printf("slice len: %d, slice cap: %d\n", len(s), cap(s))
// 	return s

// }
// func WordCounts(s string) map[string]int {

// 	count := make(map[string]int, 0)
// 	words := strings.Fields(s)

// 	for _, word := range words {

// 		count[word]++
// 	}

// 	fmt.Println(words, count)
// 	return count

// }
func apply(op func(int, int) int, a, b int) int {

	return op(a, b)

}

func sum(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(apply(sum, 3, 5))
	fmt.Println(apply(multiply, 3, 5))
	// WordCounts("test bir iki test")
	// pos := adder()
	// neg := adder()
	// fmt.Println(pos(6))
	// fmt.Println(pos(3))
	// fmt.Println(neg(6))
	// fmt.Println("use the panic")

	// 	row, cols := 2,3
	// matrix := make([][]int,row)

	// for i := range matrix{

	// 	matrix[i] = make([]int,cols)
	// }

	// for i := 0; i < row ; i ++{

	// 	for j := 0; j < cols ; j ++{

	// 		matrix[i][j] = (i+j)*2

	// 	}
	// }

	// for _,row := range matrix{

	// 	fmt.Println(row)
	// }

}
