package main

import "fmt"

type user struct {
	uname string
	age   int
}

func main() {
	u := user{
		uname: `kapil`,
		age:   34,
	}

	fmt.Println(u)
	fmt.Println(&u)
}

func (u *user) String() string {
	return fmt.Sprintf("User name is %v", u.uname)
}
