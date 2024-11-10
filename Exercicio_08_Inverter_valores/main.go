package main

import "fmt"

func main() {
	/*
	   Dado um conjunto de números, retorne o inverso aditivo de cada um.
	    Cada positivo se torna negativo, e os negativos se tornam positivos.

	   [1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
	   [1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
	   [] --> []
	*/

	list := []int{1, -2, 3, -4, 5}
	var convertedList []int
	for _, lists := range list {
		convertedList = append(convertedList,lists * (-1))
	}
	fmt.Println(convertedList)
}