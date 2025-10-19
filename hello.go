package main

import (
	"math"
)

func swap(x, y string) (string, string) {

	return y, x
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(z float64) float64 {
	for i := 1; i < 10; i++ {
		z = z - (z*z-2)/(2*z)
	}
	return z
}

// func main() {

// 	fmt.Print(Sqrt(()))

// 	fmt.Println(pow(3, 2, 10))
// 	//Usage swap function but use Println to print the result
// 	fmt.Println(swap("alkaya", "emre"))
// 	//second way to use swap function
// 	a, b := swap("emre", "alkaya")
// 	fmt.Println(a, b)
// 	var isim string
// 	fmt.Scan(&isim)
// 	ptr := &isim
// 	fmt.Println("merhbaar", *ptr)

// 	*ptr = "ahmet"
// 	fmt.Println("hello", isim)

// 	//this is a for loop example
// 	// var i int = 5
// 	// for i < 10 {
// 	// 	i++
// 	// 	fmt.Println(i)
// 	// }

// 	//this is another for loop example

// }
