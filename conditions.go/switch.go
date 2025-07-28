package main

import (
	"fmt"
)

func main() {
	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// }
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's weekday")
	// }
	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("Is's after noon")
	// }
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't Know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
