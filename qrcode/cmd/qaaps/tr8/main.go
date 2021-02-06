package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	"strings"
)

func main() {
	http.NewServeMux()
	runtime.NewServeMux()
	p := func(h string) bool {
		c := fmt.Sprintf("%c",[]rune(h)[1])
		return strings.Contains("aeiou", c)
	}

	rs := filter(p, "Iron Man","Batman", "Superman","Spider-man", "Wonder Woman","Iron Fist", "Daredevil","Supergirl", "Flash")
	fmt.Printf("Real heroes\n %v\n", rs)
}

func filter(predicate func(string) bool, s ...string, ) []string {
	r := make([]string, 0)
	for _, h := range s {
		b := predicate(h)
		if b{
			r = append(r, h)
		}
	}
	return r
}
