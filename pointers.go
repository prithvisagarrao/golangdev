package main

import "fmt"

func main() {
	i := 42
	p := &i
	fmt.Println(i)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}
