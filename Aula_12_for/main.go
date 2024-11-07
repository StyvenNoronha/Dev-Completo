package main

import "fmt"

func main() {
	/*
			soma:=0
			for i:=0; i<10; i++{
				soma = soma + 1
				fmt.Println(soma)
			}

		sub := 10
		for i := 10; i >= 1; i-- {
			sub -= 1
			fmt.Println("de ", sub, "ate ", i)
		}
	*/

	lista := []string{"a", "b", "c", "d"}

	for _, aparece := range lista {
		fmt.Println(aparece)
	}

}
