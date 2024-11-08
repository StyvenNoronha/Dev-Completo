package main

import "fmt"

func main() {

	/*
		Escreva um programa que encontre a soma de cada número de 1 a num.
		O número sempre será um inteiro positivo maior que 0. 
		Sua função só precisa retornar o resultado, o que é mostrado entre parênteses no exemplo abaixo
		é como você chega a esse resultado e não faz parte dele, veja os testes de exemplo.
		Por exemplo (Entrada -> Saída) :
		2 -> 3 (1 + 2)
		8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
	*/
	num := 8
	var soma int
	for i := 1; i <= num; i++ {
		soma = soma + i
	}
	fmt.Println(soma)
}
