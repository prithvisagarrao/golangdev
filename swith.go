package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on:")
	switch v:= runtime.GOARCH; v{
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s", v)
	}
}
