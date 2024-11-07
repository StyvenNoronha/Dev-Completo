package main

import "fmt"

func main(){
	//Arrays
	/*
	var array [2]string
	array[0] = "styven"
	array[1] = "Cesar"
	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])

	numPrimo := [6]int{2,3,5,7,11,13}
	fmt.Println(numPrimo)
	*/

	//Slice
	nomes:= []string{"styven","cesar","noronha"}
	fmt.Println(nomes)

	slice := make([]string,5)
	slice[0] = "styven"
	slice[1] = "styven"
	fmt.Println(slice)
	fmt.Println(len(nomes)) //mostrar o tamanho do slice
	nomes = append(nomes, "joão") //adiciona mais um valor
	fmt.Println(nomes)

}