package main

import (
	"fmt"
	foo "hello-shapes-go/foo"
	shapes "hello-shapes-go/shapes"
	"reflect"
)

type rectangle struct {
	base   float64
	height float64
}

func (x rectangle) Area() float64 {
	return x.base * x.height
}

func getArea(x shapes.Shape) {
	fmt.Printf("x.TypeOf(): %s\n", reflect.TypeOf(x))
	fmt.Printf("x.Area(): %f\n", x.Area())
}

func main() {
	fmt.Printf("hello, shapes\n")

	foo.Bar()

	// import constructs from external module
	// i.e., structs, functions
	square := shapes.Square{Side: 4.0}
	circle := shapes.Circle{Radius: 3.0}
	triangle := shapes.Triangle{Base: 3.0, Height: 5.0}

	// local scope
	// i.e,. struct and functions
	rectangle := rectangle{base: 3.0, height: 5.0}

	// NOTE: Functional thinking pattern
	// - FP: function applied to data structure
	// instead of
	// - OOP: object instance + behavior/method call requiring state
	// to be passed to private fields via constructor
	getArea(square)
	getArea(circle)
	getArea(triangle)
	getArea(rectangle)
}
