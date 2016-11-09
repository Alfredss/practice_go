package main 

import(
	"fmt"
)

type Usuario struct{
	nombre string
	apellidos string
	edad int
}

type Animal struct{
	nombre string
	edad int
}

func main() {
	persona01 := Usuario{nombre: "alfredo", apellidos: "yupanqui", edad : 50}
	animal := Animal{nombre:"gato", edad:2}
	animal2 := Animal{"perro", 5}

	fmt.Println(persona01)
	fmt.Println(animal)
	fmt.Println(animal2)

	animal3 := new(Animal) //devuelve el puntero del struc

	//podemos acceder y asignar con las siguiente sintaxis: animal3.nombre = "leon"
}