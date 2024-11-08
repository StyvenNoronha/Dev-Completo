package main

import "fmt"

func main() {
	/*
		Você obtém uma matriz de números e retorna a soma de todos os positivos.

		Exemplo [1,-4,7,12]=>1 + 7 + 12 = 20
				{1, 2, 3, 4, 5})).To(Equal(15)


		Observação: se não houver nada para somar, a soma padrão será 0.
	*/
	numbers := []int{-1, 2, 3, 4, -5}
	var soma int
	for _, num := range numbers {
		if num >= 0 {
			soma = num + soma
		}

	}
	fmt.Println(soma)

}
