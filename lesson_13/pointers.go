package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func updateSum(n int, sum *int) {
	if n == 0 {
		return
	}
	*sum += n * n
	updateSum(n-1, sum)
}

func main() {
	i := 5
	fmt.Println("Init:", i)

	zeroval(i)
	fmt.Println("Zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	sum := 0
	updateSum(3, &sum)
	fmt.Println("sum of first 3 squares:", sum)
}
