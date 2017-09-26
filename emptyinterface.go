package main

import "fmt"

type Emp interface{
}

func main() {
	var e Emp
	e = 42
	fmt.Printf("%t %v", e, e)

}
