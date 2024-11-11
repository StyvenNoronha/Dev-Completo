package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

func (p Pessoa) ListaNome(pessoa Pessoa) string {
	return fmt.Sprintf("%s tem %d anos",p.Nome,p.Idade)
}

/*
func ListaNome(nome string, idade int) string {
	return fmt.Sprintf("%s tem %d anos",nome,idade)
	fmt.Println(ListaNome(pessoa.Nome,pessoa.Idade))
}
*/

func main() {
	pessoa:= Pessoa{
		Nome: "Styven",
		Idade: 25,
		Profissao: "dev",

	}
	pessoa1:= Pessoa{
		Nome: "Otávio",
		Idade: 25,
		Profissao: "dev",

	}
	fmt.Println(pessoa.ListaNome(pessoa))
	fmt.Println(pessoa1.ListaNome(pessoa))
}
