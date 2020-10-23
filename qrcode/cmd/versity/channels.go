package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 2)
	ch <- 23
	ch <- 32
	var v = <-ch
	fmt.Println(v)
}
