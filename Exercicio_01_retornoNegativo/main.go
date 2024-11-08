package main

import "fmt"

func main() {
	/*
	Nesta tarefa simples, você recebe um número e tem que torná-lo negativo. Mas talvez o número já seja negativo?
	makeNegative(1);    // return -1
	makeNegative(-5);   // return -5
	makeNegative(0);    // return 0
	makeNegative(0.12); // return -0.12
	*/
	num := 0
	if num <= 0 {
		fmt.Println(num)
	} else {
		fmt.Println(num * -1)
	}
}
