package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(`/kundalini`, func(writer http.ResponseWriter, request *http.Request) {
		content, err := ioutil.ReadFile(`/Users/kapjoshi/go/src/github.com/gtwuser/myinterest/myfile.html`)
		if err != nil {
			log.Fatal("File not found : ", err)
		}
		writer.Write(content)
	})

	http.ListenAndServe("192.168.1.2:8081", nil)
}
