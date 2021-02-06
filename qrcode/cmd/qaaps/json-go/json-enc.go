package main

import (
	"bytes"
	"encoding/json"
	"github.com/gtwuser/myinterest/qrcode/cmd/qaaps/types"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	createJsonFile()
	var users = new(types.UsersDb)
	f, err := os.Open("../data/user.json")
	if err != nil {
		log.Fatalln("error opening file", err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(users)
	if err != nil {
		log.Fatalln("error decoding", err)
	}
	log.Println(*users)
}

func createJsonFile() {
	user1 := types.User{
		Username: "joe",
		Email:    "joe@email.com",
		Password: "change it",
	}
	user2 := types.User{
		Username: "jane",
		Email:    "jane@email.com",
		Password: "please change it",
	}

	users := types.UsersDb{
		Users: []types.User{user1, user2},
		Type:  `Simple`,
	}

	var buf = new(bytes.Buffer)
	f, err := os.Create("../data/user.json")
	if err != nil {
		log.Fatalln("error opening a file", err)
	}
	defer func() {
		f.Close()
	}()
	enc := json.NewEncoder(buf)
	err = enc.Encode(users)
	if err != nil {
		log.Fatalln("encoding failed", err)
	}
	io.Copy(f, buf)
}
