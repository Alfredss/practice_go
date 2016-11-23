package main

import (
	"fmt"
	"strconv"
)

func main() {
	//FOR
	for i := 0; i < 10; i++ {
		number_str := strconv.Itoa(i)
		fmt.Printf("Number %v \n", number_str)
	}

	//simulando WHILE

	j := 0

	for j < 10 {
		fmt.Println(j)
		j++
	}

	//ciclos infinitos uso de BREAK y CONTINUE
	k := 0
	for {
		if k == 2 {
			k++
			continue
		}
		fmt.Println(k)
		if k > 10 {
			break
		}
		k++
	}
}
