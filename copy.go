package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice), cap(slice)*2) //duplicamos la capacidad para que se cree un buen slice

	copy(copia, slice) //copy(destino, fuente)
	fmt.Println(cap(copia))
	fmt.Println(slice)
	fmt.Println(copia)

	/*
		copy copia el minimo de elementos en ambos arreglos
	*/
}
