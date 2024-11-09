package main

import (
	"fmt"
	"strconv"
)
/*
Precisamos de uma função que possa transformar uma string em um número.
Quais maneiras de conseguir isso você conhece?
Observação: não se preocupe, todas as entradas serão strings e cada string é 
uma representação perfeitamente válida de um número integral.
*/
func main(){
	number:= "123"
	i,_:=strconv.Atoi(number)
	fmt.Printf("%T",i)
}