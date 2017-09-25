package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	fmt.Printf("%T\n", pos)
	fmt.Println(&pos)
	fmt.Println(adder)
	poss := adder()
	j := poss(3)
	fmt.Println(j)
	for i := 0; i< 10;  i++{
	fmt.Println(
		pos(i),
		neg(2*i),
		)
	}
}
