package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func isCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	} else {
		fmt.Println("Not a circle")
	}
}

func (r rect) String() string {
	return fmt.Sprintf("This is a rectangle of width: %g and height: %g", r.width, r.height)
}

func (c circle) String() string {
	return fmt.Sprintf("This is a circle of radius: %g", c.radius)
}

func main() {
	r := rect{3, 4}
	c := circle{2.5}

	fmt.Println(r)

	measure(r)
	measure(c)

	isCircle(r)
	isCircle(c)
}
