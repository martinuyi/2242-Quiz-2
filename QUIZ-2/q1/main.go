//Filename:main.go
//To Demonstrate Interfaces: named collections of signature methods

package main

import (
	"fmt"
)

// A type Geometry interface
type Geometry interface {
	area() float64
	perimeter() float64
}

// A type Rectangle struct
type Rectangle struct {
	length float64
	width  float64
}

// A type Square struct
type Square struct {
	length float64
}

// Implementing area() method on Rectangle struct(a type Geometry)
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Implementing area() method on Square struct(a type Geometry)
func (s Square) area() float64 {
	return s.length * s.length
}

// Implementing perimeter() method on Rectangle struct(a type Geometry)
func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Implementing perimeter() method on Square struct(a type Geometry)
func (s Square) perimeter() float64 {
	return 2 * (s.length + s.length)
}

func calculate(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := Rectangle{
		length: 4,
		width:  3.5,
	}

	s := Square{
		length: 5.0,
	}

	calculate(r)
	calculate(s)
}
