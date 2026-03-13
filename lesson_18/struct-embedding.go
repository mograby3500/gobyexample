package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func (c container) String() string {
	return fmt.Sprintf("containerr with str=%s and num=%v", c.str, c.num)
}

func main() {
	co := container{
		base: base{
			num: 6,
		},
		str: "Hello",
	}

	fmt.Println(co)
	fmt.Println(co.base.num)
	fmt.Println(co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
