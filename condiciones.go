package main 

import (
	"fmt"
)
func main() {
	x:=10
	y:=20
	if x>y {
		fmt.Printf("%d Es mayor que %d",x,y)
	}else if x<y{
		fmt.Printf("%d Es mayor que %d",y,x)
	}else{
		fmt.Printf("%d y %d son iguales", x, y)
	}
}