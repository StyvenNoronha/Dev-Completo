package main

import "fmt"

func main() {
	yearsPet := [3]int{1, 15, 15}
	years := 7
	if years == 1 {
		fmt.Println(yearsPet)
	}else if years == 2{
		yearsPet := [3]int{2, 24, 24}
		fmt.Println(yearsPet)
	}else{
		catYears:= 24 + (years - 2 ) * 4
		dogYears:= 24 + (years - 2 ) * 5
		yearsPet:= [3]int{years,catYears,dogYears}
		fmt.Println(yearsPet)
	}

}
/*

Eu tenho um gato e um cachorro.

Eu os peguei na mesma época que o gatinho/cachorrinho. Isso foi human Years anos atrás.

Retorne suas respectivas idades agora como [ humanYears, catYears, dogYears]

NOTAS:

humanYears>= 1
 apenas números inteiros
Anos de gato
15anos de gato para o primeiro ano
+9anos de gato para o segundo ano
+4anos de gato para cada ano depois disso
Anos de cão
15anos de cachorro para o primeiro ano
+9anos de cachorro para o segundo ano
+5anos de cachorro para cada ano depois disso
*/