package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, c := range "hello-world" {
		fmt.Println(i, string(c))
	}
}
