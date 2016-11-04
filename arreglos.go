package main

import(
	"fmt"
)

func main() {
	var array [] int
	fmt.Println(array) //  salida : []


	var array1 [5] int
	fmt.Println(array1) //[0,0,0,0,0]


	array2 := [3] int {4, 5, 6}
	fmt.Println(array2) //[4,5,6]


	array3 := [3] int {1,2}
	fmt.Println(array3) //[1,2,0]


	array4:= [4]string {"naranjas", "manzanas", "peras"}
	array4[3] = "uvas"
	for i := 0; i < len(array4); i++ {
		fmt.Println(array4[i])
	}


	var matriz [3][3] int
	matriz[0][2] = 5
	fmt.Println(matriz) 
}