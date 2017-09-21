package main

import(
	"fmt"
	"math"
)

func sqrt(x, n, lim  float64) float64 {

	if v := math.Pow(x, n); v<lim {
		return v
}	else {
		fmt.Printf("%g > %g\n", v, lim)
}
	return lim
}

func main() {
	fmt.Println(sqrt(3, 2, 20))
	fmt.Println(sqrt(3, 3, 20))
}
