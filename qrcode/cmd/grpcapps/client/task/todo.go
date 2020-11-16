package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Printf("No subcommands provided\n")
		os.Exit(1)
	}

	var err error
	switch cmd := flag.Arg(0); cmd {
	case `add`:
		err = add(strings.Join(flag.Args()[1:], " "))
	case `list`:
		err = list()
	default:
		fmt.Printf("No subcommands\n")
	}

	if err != nil {
		log.Fatal("Error while executing subcommands\n")
	}
}

func add(subCmd string) error {
	return nil
}

func list() error {
	return nil
}
