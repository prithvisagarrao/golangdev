package main

import "fmt"

func main() {
	primes := [6]int{1,2,3,4,5,6}
	var s []int = primes[2:5]
	fmt.Println(s)

	names := [3]string{"jon", "dan", "ron"}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	
	b[0] = "AAA"
	fmt.Println(a, b)
	fmt.Println(names)

	def := primes[1:5]
	def2 := def[:3]
	def3 := def2[1:]
	fmt.Println(def)
	fmt.Println(def2)
	fmt.Println(def3)
}
