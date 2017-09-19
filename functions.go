package main

import(
	"fmt"
)

func main(){
	fmt.Println(add(7,6))
}

func add(x, y int) (int, int){
	return y,x
}
