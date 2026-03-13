package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type Element[T any] struct {
	next *Element[T]
	val  T
}

type List[T any] struct {
	head, tail *Element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &Element[T]{
			val: v,
		}
		list.tail = list.head
	} else {
		list.tail.next = &Element[T]{
			val: v,
		}
	}
}

func (lst List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if a == 3 || !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	next, stop := iter.Pull(genFib())
	defer stop()

	val, ok := next() // first value: 0
	fmt.Println(val, ok)

	val, ok = next() // second value: 1
	fmt.Println(val, ok)

	val, ok = next() // 1
	fmt.Println(val, ok)

	val, ok = next() // 2
	fmt.Println(val, ok)

	val, ok = next() // val=0, ok=false !!!
	fmt.Println(val, ok)

	val, ok = next() // val=0, ok=false !!!
	fmt.Println(val, ok)

	for n := range genFib() {
		if n >= 100 {
			break
		}
		fmt.Printf("%v ", n)
	}

	fmt.Println()
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Printf("%v ", e)
	}
	fmt.Println()
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}
}
