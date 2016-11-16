package main

import (
	"fmt"
)

type Usuario struct {
	nombre    string
	apellidos string
	edad      int
}

type Animal struct {
	nombre string
	edad   int
}

//METODO en una FUNCION
func (users Usuario) nombre_completo() string {
	return users.nombre + " " + users.apellidos
}

func (this *Usuario) set_name(name string) {
	this.nombre = name
}

func main() {

	persona01 := Usuario{nombre: "alfredo", apellidos: "yupanqui", edad: 50}
	animal := Animal{nombre: "gato", edad: 2}
	animal2 := Animal{"perro", 5}

	fmt.Println(persona01)
	fmt.Println(animal)
	fmt.Println(animal2)

	fmt.Println(animal.nombre) //acceder a un atributo

	animal3 := new(Animal) //devuelve el puntero del struc
	fmt.Println(animal3)
	//podemos acceder y asignar con las siguiente sintaxis: animal3.nombre = "leon"

	//Metodo
	coder := new(Usuario)
	coder.nombre = "Alfredo"
	coder.apellidos = "Yupanqui"

	coder.set_name("Rafael")

	fmt.Println(coder.nombre_completo())
}
