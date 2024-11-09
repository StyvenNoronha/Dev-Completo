/*
Dado um array de inteiros.

Retorna uma matriz, onde o primeiro elemento é a contagem de números positivos e o segundo elemento é a soma de números negativos. 0 não é positivo nem negativo.

Se a entrada for uma matriz vazia ou for nula, retorne uma matriz vazia.

Exemplo
Para entrada [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], você deve retornar [10, -65].
*/
package main

import "fmt"

func main(){
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	var sumNegativo int
	var positivo int
	for _, num := range numbers {
		if num <= 0 {
			sumNegativo = num + sumNegativo
		} else{
			positivo = num
		}

	}
	res:= []int{positivo,sumNegativo}
	fmt.Println(res)
}