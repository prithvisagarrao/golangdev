package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hi"
	a[1] = "user"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := []int{223,33,44}
	fmt.Println(b)
}
