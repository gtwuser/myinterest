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
	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		imgPath := flag.String("path", "", "abs path to image file")
		flag.Parse()
		f, _ := os.Open(*imgPath)
		rd := bufio.NewReader(f)
		all, _ := ioutil.ReadAll(rd)

		writer.Header().Set(`Content-Type`, `image/jpeg`)
		writer.Write(all)
	})

	log.Fatal(http.ListenAndServe(`:8080`, nil))
}
