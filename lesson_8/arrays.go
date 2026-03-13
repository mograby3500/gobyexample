package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Println("Empty: ", arr)

	arr[4] = 100

	// assigning arrays copies them
	var a = arr
	a[4] = 200
	fmt.Println("arr ", arr)
	fmt.Println("a: ", a)

	fmt.Println("Length: ", len(arr))

	fmt.Println("Arr rotated by 2: ")
	L := len(arr)
	for i := range L {
		j := (i + 2) % L
		fmt.Print(arr[j], " ")
	}
	fmt.Println("\n------------------")

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	fmt.Println("2d array: ", twoD)

	for i := range len(twoD) {
		for j := range len(twoD[i]) {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d array: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("new assigned 2d array: ")
	for i := range 2 {
		for j := range 3 {
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Println()
	}
}
