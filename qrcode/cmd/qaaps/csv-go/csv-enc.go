package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/kapjoshi/myinterest/qrcode/cmd/qaaps/types"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	//header := []string{`username`, `email`, `password`}
	////user1 := []string{`joe`, `joe@email.com`, `change it`}
	////user2 := []string{`jane`, `jane@email.com`, `please change it`}
	//f, err := os.Create(`../data/csv-enc.csv`)
	//if err != nil {
	//	log.Fatalln("error while creating file ", err)
	//}
	//defer f.Close()
	f, err := os.Open(`../data/csv-enc.csv`)
	if err != nil {
		log.Fatalln("error while creating file ", err)
	}
	//w := csv.NewWriter(f)
	//w.Write(header)
	//db := readJsonFile()
	//for _, u := range db.Users {
	//	s := u.EncodeAsString()
	//	w.Write(s)
	//}
	//now := time.Now()
	//w.Flush()
	//all := [][]string{header, user2, user1}
	//w.WriteAll(all)
	//if w.Error() != nil {
	//	log.Fatalln("Err in writing")
	//}
	//fmt.Println(time.Since(now))

	//Reader
	r := csv.NewReader(f)
	if err != nil {
		log.Fatalln("csv file reading error", err)
	}
	user := new(types.User)
	header, err := r.Read()
	log.Println("Header", header)
	for {
		record, err := r.Read()
		if err == nil {
			user.FromCSV(record)
			log.Println("User :", user)
		} else if io.EOF == err {
			break
		} else if err != nil {
			log.Fatalln("error on reading", err)
		}
	}
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
