package fizzbuzz

import (
	"fmt"
	"strconv"
	"time"
)

func Fizzbuzz() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
			fmt.Println(time.Weekday(3))
		default:
			fmt.Println(strconv.Itoa(i))
		}
	}
}
