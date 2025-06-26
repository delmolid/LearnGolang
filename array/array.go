package main

import "fmt"

func main() {
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])

	// fmt.Println("leng : ", len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// b = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// Déclaration avec taille explicite
	var arr [5]int

	// Déclaration avec initialisation
	arr := [5]int{1, 2, 3, 4, 5}

	// Taille automatique
	arr := [...]int{1, 2, 3, 4, 5}

	// Tableau 2D
	var matrix [2][3]int

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// twoD = [2][3]int{
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// }
	// fmt.Println("2d: ", twoD)

	// Accès
	element := arr[index]

	// Modification
	arr[index] = value

	// Longueur
	length := len(arr)

	// Parcours avec range
	for i, value := range arr {
		// i = index, value = valeur
	}

	// Parcours avec indices seulement
	for i := range arr {
		// i = index
	}
}
