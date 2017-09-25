package main

import(
	"fmt"
	_ "math"
)

type Area interface{
	area() float64
	peri() float64
}

type rect struct{
	X, Y float64
}

type MyFloat float64

func (r rect) area() float64 {
	return r.X*r.Y
}

func (r rect) peri() float64 {
	return 2*r.X + 2*r.Y
}

func (f MyFloat) area() float64 {
	return float64(f*f)
}

func (f MyFloat) peri() float64 {
	return float64(2*f)
}

func measure(a Area){
fmt.Println(a)
fmt.Println(a.area())
fmt.Println(a.peri())
}

func main() {
	var a Area
	r := rect{3,4}
	a = r
	fmt.Println(a.area())

	measure(r)
	
	f := MyFloat(5)
	a = f
	fmt.Println(a)
	fmt.Println(a.area())
	fmt.Println(a.peri())

	d := rect{}
	a = d
	fmt.Println(a.area())
	fmt.Printf("%T %+v\n", d, d)

	var i interface{} = "hello"
	fmt.Printf("%T, %v\n", i, i)
		
}
