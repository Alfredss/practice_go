package main

import (
	"fmt"
)

func main() {
	var x, y *int //declaramos los punteros que almacenan direcciones
	entero := 5   //almacenamos 5 en la variable entero

	x = &entero //almacenamos la direccion de entero en el puntero x
	y = &entero

	*x = 7 //accedemos al valor del puntero x con * y le asignamos el nuevo valor 6

	//como cambiamos el valor del puntero x entonces y y entero tambien se ven
	//afectadas al apuntar estas a la misma direccion
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(entero)

}
