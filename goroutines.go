package main

import "fmt"

func say(s string) {
	{
	fmt.Println(s)
	}
}

func main() {
	go say ("hello")
	var input string
	fmt.Scanln(&input)
	say("world")
}

/*package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
*/
