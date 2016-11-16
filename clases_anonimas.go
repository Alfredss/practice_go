package main

import "fmt"

type Humano struct {
	name string
}

func (this Humano) hablar() string {
	return "Bla Bla Bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Humano
	Dummy
}

func (this Tutor) hablar() string {
	return "metodo sobrescribido"
} //Sobrescribimos el metodo de Humano

func main() {
	alfredo := Tutor{Humano{name: "Alfredo"}, Dummy{name: "Dummy"}}
	fmt.Println(alfredo.Humano.name)
	fmt.Println(alfredo.hablar())
}
