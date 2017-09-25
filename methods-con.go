package main

import (
	"fmt"
	"math"
)

type Vertex float64

type Test int

func (x Test) Subs() int{
	return int(x+10)
}


func (v Vertex) Abs() float64 {
	if v < 0 {
		return float64(-v)
		}
	return float64(v+10.1)
}

func main() {
	f := Vertex(math.Sqrt2)
	g := Test(3)
	fmt.Println(g)
	fmt.Println(g.Subs())
	//fmt.Println(7.Subs())
	fmt.Println(f.Abs())
}
