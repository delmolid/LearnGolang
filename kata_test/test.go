package main

import (
	"fmt"
	"strings"
)

func main() {

	domaine := "ovhcloud.fr.com"

	array := strings.Split(domaine, ".")

	extension := array[1:]

	fmt.Println(extension)

	// for i, values := range array {
	// 	fmt.Println(i, values)
	// 	fmt.Println(array[len(array)-1])

	// }

}
