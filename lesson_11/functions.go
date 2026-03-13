package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}

func addn(nums ...int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}

	return sum
}

func vals() (int, int) {
	return 3, 7
}

// clousers
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion and dp
var dp [100]int

func solve(n int) int {
	if n < 2 {
		return n
	}
	if dp[n] != 0 {
		return dp[n]
	}
	// fmt.Println("calculating fib", n)
	dp[n] = solve(n-1) + solve(n-2)
	return dp[n]
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add3(1, 2, 3))
	fmt.Println(vals())

	s := []int{1, 2, 3}

	a, b := vals()
	s = append(s, a, b)
	fmt.Println(s)

	fmt.Println(addn(1, 2, 3))
	fmt.Println(addn(1, 2, 3, add(1, 2)))

	nums := []int{1, 1, 1}
	fmt.Println(addn(nums...))

	fmt.Println("---------\nClosuers")
	//closuers
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//recursion
	for i := range 10 {
		fmt.Print(solve(i), " ")
	}
	fmt.Println()
}
