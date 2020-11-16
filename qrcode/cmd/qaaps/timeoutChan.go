package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		select {
		case r0 := <-search("q1"):
			fmt.Printf("Search0 time:%v\n", r0)
		case r1 := <-search("q2"):
			fmt.Printf("Search1 time:%v\n", r1)
		case r2 := <-search("q3"):
			fmt.Printf("Search2 time:%v\n", r2)
		case <-time.After(time.Duration(2) * time.Second):
			fmt.Println("Search timeout")
		}
	}
}

func search(query string) <-chan time.Time {
	num := rand.Intn(8)
	fmt.Println(`Random :`, num, time.Now())
	return time.After(time.Duration(num) * time.Second)
}
