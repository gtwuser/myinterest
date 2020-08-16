package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "www.google.com:80")
	errMsg(err, "Error connecting to baidu")
	if err != nil {
		return
	}

	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.google.com\r\n"
	data += "connection: close\r\n"
	data += "\r\n\r\n"

	_, err = io.WriteString(conn, data)
	errMsg(err, "Error in writing data")
	if err != nil{
		return
	}

	//buffer := bytes.Buffer{}
	var buffer [1024]byte
	for{
		n, err := conn.Read(buffer[:])
		errMsg(err, "error in reading content")
		if err != nil {
			return
		}
		ioutil.WriteFile("./baidu.com.html", buffer[:n], os.ModePerm)
	}
}

func errMsg(err error, msg ...string) {
	if err != nil {
		log.Print(msg, err)
	}
}
