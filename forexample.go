package main

import (
	"fmt"
	"time"
)

// func Sqrt(x float64) float64 {
// 	z := x / 2
// 	for i := 1; i < 10; i++ {
// 		prev := z
// 		z -= (z*z - x) / (2 * z)

// 		if math.Abs(z-prev) < 1e-10 {
// 			break
// 		}
// 		fmt.Println(z)
// 	}
// 	return z
// }

func main() {

	fmt.Println("when's saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today is saturday")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	case today + 6:
		fmt.Println("in six days")
	default:
		fmt.Println("too far away")
	}

	// // fmt.Println("test")
	// // fmt.Println(Sqrt(10))
	// // fmt.Println("kütüphane", math.Sqrt(10))
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("macOS")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Println("os:", os)
	// }

}
