package main

import "fmt"

func main() {

	var n int = 9

	if n%2 == 0 {
		fmt.Println("n is even")
	} else {
		fmt.Println("n is odd")
	}

	if m := 10; m < 0 || m == 10 {
		fmt.Println("m less than 0 or equal to 10")
	} else if m < 10 {
		fmt.Println("m is less than 10")
	} else {
		fmt.Println("m is greater than 10")
	}
}
