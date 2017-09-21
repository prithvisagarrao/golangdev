package main

import (
	"fmt"
	"math"
)

func main(){
	var x,y = 2,3
	var f float64 = math.Sqrt(float64(x*y))
	var z uint = uint(f)
	fmt.Println(x,y,z,f)
}
