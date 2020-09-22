package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	afChan := time.After(2 * time.Second)
	//done := make(chan bool, 1)
	signal.Notify(sig)
	for {
		select {
		case s := <-sig:
			fmt.Println("Got a signal", s)
			os.Exit(1)
		case t := <-afChan:
			fmt.Println("timed out 😲", t)
		}
	}
	//go func() {
	//	<- afChan
	//	log.Fatal("Exiting 🤪")
	//}()
	//s := <- sig
}
