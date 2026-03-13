package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 26
	return &p
}

func newPersonByValue() person {
	return person{
		"ahmed", 26,
	}
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2 * (r.height + r.width)
}

func main() {
	a := person{"ahmed", 26}
	fmt.Println(a)

	b := person{name: "seif"}
	fmt.Println(b)

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	jon := newPerson("jon")
	fmt.Println(jon.name, jon.age)

	max := &jon
	(*max).name = "max"

	fmt.Println(jon, *max)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex", true,
	}
	fmt.Println(dog)

	x := newPersonByValue()
	x.age = 27
	fmt.Println(x)

	r1 := rect{4, 5}
	fmt.Println("Area of r1", r1.area())
	fmt.Println("Perimeter of r1", r1.perim())

}
