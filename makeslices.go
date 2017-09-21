package main

import "fmt"

func main() {
	a := make([]int,0, 5)
	fmt.Println(a)
	a = append(a, 1,2,3,4)
	fmt.Println(a)
	for i, v := range a{
	fmt.Println(i,v)
}
}
