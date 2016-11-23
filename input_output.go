package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//OUTPUT
	//imprimir con salto de lineea
	number := 5
	fmt.Println(number)

	//imprimir con formato, introducimos verbos en las cadenas
	nombre := "Alfredo"
	fmt.Printf("mi nombre es %v\n", nombre)

	numero := 4
	fmt.Printf("El numero es %d\n ", numero)
	/*
		%v para valores de variables por default
		%d para numeros
		para conocer mas:
			https://golang.org/pkg/fmt/
	*/

	//INPUT
	var edad int
	fmt.Println("Cual es tu Edad: ")
	fmt.Scanf("%d\n", &edad) //necesario agregar \n para que interprete el salto de linea
	fmt.Printf("tu edad es %d\n", edad)

	//USANDO BUFIO
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ingrese su nombre")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}
