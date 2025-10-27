package main

import (
	"fmt"
)

type Vertex struct{ X, Y float64 }

func (v *Vertex) sum() float64 {

	return v.X + v.Y

}

func (v *Vertex) multipy() float64 {

	return v.X * v.Y
}

type allfunc interface {
	sum() float64
	multipy() float64
}

func test(i allfunc) {

	fmt.Println(i.sum())
	fmt.Println(i.multipy())
}

func main() {

	v := &Vertex{3, 3}
	test(v)

}
