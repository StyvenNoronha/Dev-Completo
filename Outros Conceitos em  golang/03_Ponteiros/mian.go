package main

import "fmt"

/*
Ponteiros: um ponteiro nada mais é do que uma variável
que ao invés de armazenar um valor (ex 1, string),
armazena um end de memoria
*/

func main(){
	/*
	Memoria => essa memoria tem um endereço -> esse endereço guarda um valor
	*/

	i:= 1
	fmt.Println(i)
	fmt.Println(&i)

	a:= &i

	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(&a)
}