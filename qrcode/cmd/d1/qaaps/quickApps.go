package main

import (
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("testing %s error", "dummy")
	log.Print(err)
}
