package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {

	return y, x
}

func main() {

	//Usage swap function but use Println to print the result
	fmt.Println(swap("alkaya", "emre"))
	//second way to use swap function
	a, b := swap("emre", "alkaya")
	fmt.Println(a, b)
	var isim string
	fmt.Scan(&isim)
	ptr := &isim
	fmt.Println("merhbaar", *ptr)

	*ptr = "ahmet"
	fmt.Println("merhbaar", isim)
}
