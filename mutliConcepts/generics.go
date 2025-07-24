package main

import "fmt"

// func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
// 	for i := range s {
// 		if v == s[i] {
// 			return i
// 		}
// 	}
// 	return -1
// }

// type List[T any] struct {
// 	head, tail *element[T]
// }
// type element[T any] struct {
// 	next *element[T]
// 	val  T
// }

// func (lst *List[T]) Push(v T) {
// 	if lst.tail == nil {
// 		lst.head = &element[T]{val: v}
// 		lst.tail = lst.head
// 	} else {
// 		lst.tail.next = &element[T]{val: v}
// 		lst.tail = lst.tail.next
// 	}
// }

// func (lst *List[T]) AllElements() []T {
// 	var elems []T
// 	for e := lst.head; e != nil; e = e.next {
// 		elems = append(elems, e.val)
// 	}
// 	return elems
// }

// func PrintInt(s []int) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// func PrintString(s []string) {
// 	for _, v := range s {
// 		fmt.Print(v)
// 	}
// }

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

type Person interface {
	Work()
}
type worker string

func (w worker) Work() {
	fmt.Printf("%s is working\n", w)
}

func DoWork[T Person](things []T) {
	for _, v := range things {
		v.Work()
	}
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

type Number interface {
	int | float64
}

func MultiplyTen[T Number](a T) T {
	return a * 10
}

type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, v := range g {
		fmt.Println(v)
	}
}

func PrintNew[T any](g GenericSlice[T]) {
	for _, v := range g {
		fmt.Println(v)
	}
}

func main() {
	// var s = []string{"foo", "bar", "zoo"}

	// fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// _ = SlicesIndex[[]string, string](s, "zoo")
	// lst := List[int]{}
	// lst.Push(10)
	// lst.Push(13)
	// lst.Push(23)
	// fmt.Println("list:", lst.AllElements())

	// Print([]int{1, 2, 10})
	// Print([]string{"Hello", " World "})

	var a, b, c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})

	Equal("a", "a")

	fmt.Println(MultiplyTen(10))
	fmt.Println(MultiplyTen(5.55))

	g := GenericSlice[int]{1, 2, 3}

	g.Print()   //1 2 3
	PrintNew(g) //1 2 3

}
