package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.After(1 * time.Nanosecond)
	for i := 0; i < 5; i++ {
		select {
		case s := <-a:
			fmt.Println("Time now ", s)
		default:
			fmt.Println("Default use case")
		}
	}
}
