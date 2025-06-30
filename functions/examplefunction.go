package main

import "fmt"

func add(num1 int, num2 int) int {

	return num1 + num2
}

func plusadd(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func vals() (int, int) {
	return 3, 7
}

func main() {

	// res := add(1, 2)
	// fmt.Println("1+2 = ", res)
	// res = plusadd(1, 2, 3)
	// fmt.Println("Total", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
