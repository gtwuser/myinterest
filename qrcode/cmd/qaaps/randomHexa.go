package main

import "fmt"

func main() {
	ch := make(chan string, 10)
	greetings(ch)
	for v := range ch {
		fmt.Print(v)
	}
	fmt.Println()
}

func greetings(ch chan<- string) {
	for i := 0; i < 29; i++ { // test both more than 10 and less than 10 observe
		select {
		case ch <- `a`:
		case ch <- `b`:
		case ch <- `c`:
		case ch <- `d`:
		case ch <- `e`:
		case ch <- `f`:
		case ch <- `0`:
		case ch <- `1`:
		case ch <- `2`:
		case ch <- `3`:
		case ch <- `4`:
		case ch <- `5`:
		}
	}
	close(ch)
}
