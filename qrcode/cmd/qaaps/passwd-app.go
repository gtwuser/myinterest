package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/textproto"
)

type MD map[string][]string

func main() {
	md := MD{}
	var key string
	key = "accept"
	fmt.Println(md)
	md[key] = append(md[key], "s")
	fmt.Println(md[key])
	md[key] = append(md[key], "y")
	fmt.Println(md[key])
	md[key] = append(md[key], "z")
	fmt.Println(md[key])
	//fmt.Println(md)
	//passwd := generatePasswd()
	//fmt.Println(passwd)
	//canonicalForm := generateCanonical()
	//fmt.Println(canonicalForm)
}

func generateCanonical() string {
	resp := textproto.CanonicalMIMEHeaderKey("you-are-too-good")
	return resp
}

func generatePasswd() string {
	adminPass, _ := bcrypt.GenerateFromPassword(
		[]byte("Cosco123"), bcrypt.DefaultCost,
	)
	defer func() {
		fmt.Println(string(adminPass))
	}()
	return string(adminPass)
}
