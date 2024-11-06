package main

import "fmt"

func soma(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

func subtracao(numberOne, numberTwo int) int {
	return numberOne - numberTwo
}

func name(nome,sobrenome string)string{
	return nome  + sobrenome
}

func main() {
	subtracaoDeValor := subtracao(20, 10)
	fmt.Println(subtracaoDeValor)
	
	fmt.Println(soma(2, 2))


	fmt.Println(name("styven","noronha"))

}
