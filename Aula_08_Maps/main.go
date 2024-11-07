package main


import "fmt"

func main(){
	/*
	idade := map[string]int{}
	idade["styven"] = 28
	idade["cesar"] = 4
	fmt.Println(idade)

	//pegar o valor dentro do map
	fmt.Println(idade["styven"])
	*/

	anoNasc := map[string]int{
		"styven":1999,
		"cesar":2000,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["styven"])

	//adicionar um valor novo para o map
	anoNasc["alberto"] = 2024
	fmt.Println(anoNasc)

	}
