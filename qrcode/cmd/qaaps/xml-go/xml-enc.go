package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/gtwuser/myinterest/qrcode/cmd/qaaps/types"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	usersDb := readJsonFile()

	var buf = new(bytes.Buffer)
	enc := xml.NewEncoder(buf)
	err := enc.Encode(usersDb)
	if err != nil {
		log.Fatalln("error in encoding", err)
	}

	xmlFile := "../data/users.xml"
	f, err := os.Create(xmlFile)
	if err != nil {
		log.Fatalln("error in creating", err)
	}
	io.Copy(f, buf)
	err = f.Close()
	if err != nil {
		log.Fatalln("error in closing file", err)
	} else {
		log.Println("Closed")
	}
	uf, err := os.Open(xmlFile)
	if err != nil {
		log.Fatalln("error in opening", err)
	}
	defer uf.Close()
	usersDec := new(types.UsersDb)
	dec := xml.NewDecoder(uf)
	err = dec.Decode(usersDec)
	if err != nil {
		log.Fatalln("decoding failed", err)
	}
	log.Println(usersDec)
}

func readJsonFile() *types.UsersDb {
	fu, err := os.Open(`../data/user.json`)
	if err != nil {
		log.Fatalln("error in opening file ", err)
	}
	defer fu.Close()
	usersDb := new(types.UsersDb)
	dec := json.NewDecoder(fu)
	err = dec.Decode(usersDb)
	if err != nil {
		log.Fatalln("error in decoding", err)
	}
	return usersDb
}
