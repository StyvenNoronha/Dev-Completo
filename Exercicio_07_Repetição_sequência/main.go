package main

import "fmt"

func main() {
	/*
	   Escreva uma função que aceite um inteiro ne uma string s como parâmetros e
	    retorne uma string de vezes sexatas repetidas n.

	   Exemplos (entrada -> saída)
	   6, "I"     -> "IIIIII"
	   5, "Hello" -> "HelloHelloHelloHelloHello"
	*/

	repeatNumber := 10
	word := "oi"
	var repeatString string
	for i := 0; i <= repeatNumber; i++ {
		repeatString += word
	}
	fmt.Println(repeatString)

}