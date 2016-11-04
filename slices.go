package main 

import(
	"fmt"
)

func main() {

	var array []int	//slice porque no especifica tamaño
	fmt.Println(array)

	//haciendo un slicing a un array (convirtiendo un array a un slice)

	array1 := [4] int {1,2,3,4}
	slice := array1[1:3]
	slice1 := array1[1:]
	slice2 := array1[:4]
	fmt.Println(slice) 
	fmt.Println(slice1)
	fmt.Println(slice2)
}