package main

import "fmt"
func main() {
	//variables de la forma tipada
	
	var numero int
	var cadena string
	var booleano bool
	var array[2] string

	array[0] = "hello" //añadimos elementos a la lista
	array[1] = "world"

	numero = 05
	cadena = "hola mundo"
	booleano = true

	fmt.Println(numero)
	fmt.Println(cadena)
	fmt.Println(booleano)
	fmt.Println(array[0], array[1])

	//variables con asignacion dinamica

	number := 25
	str:="hola"
	boo:=false

	fmt.Println(number)
	fmt.Println(str)
	fmt.Println(boo)
}