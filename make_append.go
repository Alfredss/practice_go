package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 5) //crea un array de tipo slice
	slice = append(slice, 8)   //agregamos elementos a un slice
	fmt.Println(slice)         //[0,0,0]
	fmt.Println(len(slice))    //tamaño 3
	fmt.Println(cap(slice))    //capacidad 5
}

/*
	make(tipo, longitud, capacidad)
	append(lista, elemento)
*/
