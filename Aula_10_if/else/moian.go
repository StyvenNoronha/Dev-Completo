package main

import "fmt"

func main(){
	n1:= 107
	n2:=10
	if n1>n2{
		fmt.Println("o", n1,"e maior que ", n2 )
	}else if n1 == n2{
		fmt.Println("eles sao iguais")
	}else{
		fmt.Println("o", n2,"e maior que ", n1 )
	}
}