package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Convirtiendo una cadena a un int
	number := 25   //asignaciom
	cadena := "02" //asignacion

	cadena_int, _ := strconv.Atoi(cadena) //como devuelve varios parametros ponemos el operador "_"

	fmt.Println(number + cadena_int) //concatenamos e imprimimos

	//convirtiendo un int a una cadena
	numero := 25 //asignacion

	numero_str := strconv.Itoa(numero) //

	fmt.Println("el numero es: " + numero_str) //concatenamos
}
