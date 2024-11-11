package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Raca  string
	Cor   string
	Patas int
}

type Animal interface {
	Barulho() string
	NumerosDePatas() int
}

func FazBarulho(animal Animal){
	fmt.Println(animal.Barulho())
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumerosDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau miau"
}

func (g Gato) NumerosDePatas() int {
	return g.Patas
}

func main() {
	dog := Cachorro{
		Raca:  "vira_lata",
		Cor:   "Azul",
		Patas: 6,
	}
	cat := Gato{
		Raca:  "vira_lata_top",
		Cor:   "Amarelo",
		Patas: 5,
	}
/*
	fmt.Println(cat.Barulho())
	fmt.Println(cat.NumerosDePatas())
	fmt.Println(dog.Barulho())
	fmt.Println(dog.NumerosDePatas())
*/
	FazBarulho(dog)
	FazBarulho(cat)
}