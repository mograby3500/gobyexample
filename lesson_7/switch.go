package main

import (
	"fmt"
	"time"
)

func main() {

	i := 6
	fmt.Println("Write", i, "as: ")

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("IDK man")
	}

	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(x any) {
		switch t := x.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	printType := func(x any) {
		fmt.Printf("Type of x is: %T\n", x)
	}

	printType(122)
}
