package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"database/sql"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

var s string
var db  *sql.DB
func main() {
	tempFile, err := ioutil.TempFile("/tmp", "example")
	s = `blank`
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO accounts SET hash=?, email=?")
	if err != nil {
		panic(err)
	}
	_, err := stmt.ExecContext(context.Background(), `hashedPassword`, `email`)
	defer log.Info().Msg(`1`)
	if err != nil {
		log.Err(err).Msg("Error while creating")
		return
	}
	dir := os.TempDir()
	log.Info().Msgf("temp loc %v", dir)
	defer log.Info().Msg(`2`)
	n, err := tempFile.Write([]byte(`Hello World`))
	if err != nil {
		log.Err(err).Msg("Error while writing")
		return
	}
	defer log.Info().Msg(`3`)
	log.Info().Msgf(`bytes written %v`, n)
	log.Info().Msgf("filename %v", tempFile.Name())
	defer tempFile.Close()
	demo()
	log.Info().Msgf(`temp dir is %v`, dir)
	ioutil.WriteFile(filepath.Join(dir, "sample.txt"), []byte(`YoYo`), 0777)
	sample()
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func demo() {
	var s time.Duration = 3 * time.Second
	for i := 0; i < 3; i++ {
		jitter := time.Duration(rand.Int63n(int64(s)))
		s = s + jitter/2
		log.Info().Msg("time " + s.String())
	}
}

func sample() {
	log.Info().Msg(s)
}
