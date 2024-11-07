package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	tipo string
}

func main() {

	fmt.Println(Pessoa{"styven", 25})
	fmt.Println(Pessoa{Nome: "styven", Idade: 25})

	p1 := Pessoa{Nome: "Styven", Idade: 25}

	//fmt.Println(p1)
	//fmt.Println(p1.Idade)
	//fmt.Println(p1.Nome)
	//p1.Nome = "cesar"
	//fmt.Println(p1.Nome)
	/*
		var alunos = map[string][]Pessoa{
			"programação": {{Nome: "styven", Idade: 20}},
			"Dev":         {{Nome: "styven", Idade: 20}},
		}

		fmt.Println(alunos)
	*/

	prof := Profissao{p1, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Idade)
}
