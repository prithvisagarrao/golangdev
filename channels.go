package main

import "fmt"

func sums (s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum	
}

func main() {
	s := []int{2,3,4,2,43,-9,-3}
	c := make(chan int)
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
	_, ok := <-ch
	fmt.Println(ok)	
	go sums(s[:len(s)/2], c)
	go sums(s[len(s)/2:], c)
	x,y := <-c,<-c
	//fmt.Println(<-c,<-c)
	fmt.Println(x,y,x+y)
	//chann := make(chan int)
	//fmt.Println(<-chann)
}
