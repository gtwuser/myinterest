package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/image", img())

	log.Fatal(http.ListenAndServe(`:8080`, nil))
}

func img() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		imgPath := flag.String("path", "", "abs path to image file")
		flag.Parse()
		f, _ := os.Open(*imgPath)
		rd := bufio.NewReader(f)
		all, _ := ioutil.ReadAll(rd)

		writer.Header().Set(`Content-Type`, `image/jpeg`)
		writer.Write(all)
	}
}
