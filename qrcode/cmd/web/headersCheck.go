package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(`GET`, `https://godoc.org/github.com/sirupsen/logrus#pkg-examples`, nil)
	req.Header.Set(`Authorization`, "Bearer some-token")
	req.Header.Set(`Content-Type`, "application/json")

	logrus.Info("request URL ", req.URL)
}
