// package main

// import "fmt"

// func main() {
// 	// array := [5]int{1, 2, 3, 4, 5}
// 	// array[2] = 100
// 	// length := len(array)
// 	// fmt.Println(array)
// 	// fmt.Println(length)

// 	// arr := [3]int{10, 20, 30}

// 	// for i, value := range arr {
// 	// 	fmt.Printf("Index: %d, Valeur: %d\n", i, value)
// 	// }

// 	array := [3]int{10, 20, 30}

// 	for i := range array {
// 		fmt.Printf("Index: %d, Valeur: %d\n", i, array[i])
// 	}

// }

// package main

// import "fmt"

// type Person struct {
// 	Name   string
// 	Age    int
// 	Gender string
// }

// func main() {
// 	// Using new() to allocate memory for a Person struct
// 	p := new(Person)

// 	// Initializing the fields
// 	p.Name = "John Doe"
// 	p.Age = 30
// 	p.Gender = "Male"

// 	fmt.Println(p)
// }

package main

import "fmt"

func main() {
	// Création avec capacité spécifique
	s := make([]int, 10, 15)

	// Initialisation
	for i := 0; i < 10; i++ {
		s[i] = i + 1
	}

	fmt.Printf("Initial - Len: %d, Cap: %d, Slice: %v\n", len(s), cap(s), s)

	// Ajout dans la capacité existante
	s = append(s, 11, 12, 13, 14, 15)
	fmt.Printf("Dans la capacité - Len: %d, Cap: %d\n", len(s), cap(s))

	// Dépassement de la capacité
	s = append(s, 16, 17, 18)
	fmt.Printf("Après dépassement - Len: %d, Cap: %d\n", len(s), cap(s))

	// Vérifier l'espace libre
	espaceLibre := cap(s) - len(s)
	fmt.Printf("Espace libre: %d éléments\n", espaceLibre)
}
