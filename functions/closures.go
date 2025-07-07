package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeAccumulator() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15

	acc := makeAccumulator()
	fmt.Println(acc(1)) // 1
	fmt.Println(acc(2)) // 3
	fmt.Println(acc(3)) // 6
}
