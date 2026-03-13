package main

import "fmt"

func FindIndex[seq ~[]T, T comparable](s seq, v T) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type Point struct {
	x float32
	y float32
}

type Polygon struct {
	pts []Point
}

type Element[T any] struct {
	nxt, prev *Element[T]
	val       T
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
		var el = Element[T]{
			val:  v,
			prev: list.tail,
		}
		list.tail.nxt = &el
		list.tail = list.tail.nxt
	}
}

func (list *List[T]) Pop() {
	if list.tail == nil {
		return
	}
	list.tail = list.tail.prev
	list.tail.nxt = nil
}

func (list List[T]) String() string {
	res := ""
	first := true
	for node := list.head; node != nil; node = node.nxt {
		if !first {
			res += " ==> "
		}
		res += fmt.Sprintf("(%v, %v)", &node, node.val)
		first = false
	}
	return res
}

func FindIndexInLinkedList[T comparable](head *Element[T], v T) int {
	i := 0
	for node := head; node != nil; node = node.nxt {
		if node.val == v {
			return i
		}
		i++
	}
	return -1
}

func main() {
	var pts = []Point{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	var pt = Point{
		2, 3,
	}
	fmt.Println("Infdex found at:", FindIndex(pts, pt))

	var polygon1 = Polygon{
		pts: []Point{
			{1, 3},
			{3, 3},
			{4, 4},
		},
	}
	var polygon2 = Polygon{
		pts: []Point{
			{6, 3},
			{8, 3},
			{10, 4},
		},
	}
	var polygon3 = Polygon{
		pts: []Point{
			{15, 3},
			{17, 3},
			{19, 4},
		},
	}

	var polygonPointers = []*Polygon{&polygon1, &polygon2, &polygon3}
	fmt.Println("Infdex found at:", FindIndex(polygonPointers, &polygon3))

	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	list.Push(6)
	fmt.Println(list)
	list.Pop()
	fmt.Println(list)
	list.Pop()
	fmt.Println(list)
	list.Push(7)
	list.Push(8)
	fmt.Println(list)

	fmt.Println("Linked list node position:", FindIndexInLinkedList(list.head, 7))
}
