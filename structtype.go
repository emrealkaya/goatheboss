package main

import (
	"fmt"
)

// type Vertex struct {
// 	X int
// 	Y int
// }

type Person struct {
	Name string
	Age  int
}

func changeName(person Person) {
	person.Name = "Ahmet"
	fmt.Println("Inside changeName:", person.Name)
}

func main() {

	// p := &Person{Name: "test", Age: 30}
	// changeName(*p)
	// fmt.Println("name:", p.Name)

	// pa := [6]int{1, 2, 3, 4, 5, 6}

	// var s []int = pa[1:4]
	// fmt.Println(pa)
	// fmt.Println(s)

	// name := [4]string{
	// 	"emre",
	// 	"alkaya",
	// 	"test",
	// 	"try"}

	// fmt.Println(name)

	// a := name[0:2]
	// b := name[1:3]

	// fmt.Println(a, b)
	// b[0] = "changed"
	// fmt.Println(a, b)
	// fmt.Println(name)

	// numbbers := []int{1, 2, 3, 4, 5}
	// fmt.Println(numbbers)

	// sl := numbbers[2:4]
	// fmt.Println(sl)

	// sl[0] = 100
	// fmt.Println(sl)
	// fmt.Println(numbbers)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	test := []int{2, 3, 5, 7, 11, 13}

	test = test[:0]
	fmt.Println(test)
	test = test[1:4]
	fmt.Println(test)

	test = test[:2]
	fmt.Println(test)

	test = test[1:]

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len:%d, cap:%d %v\n", s, len(x), cap(x), x)
}
