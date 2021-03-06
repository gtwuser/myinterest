package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

type myApp struct {
}

func newMyApp() *myApp {
	return &myApp{}
}

func init() {
	os.Setenv("MYTEXT", "GOLANG")
}
var (
	app      = kingpin.New("chat", "A command-line chat application.")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()
	serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	register     = app.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "Nickname for user.").Required().String()
	registerName = register.Arg("name", "Name of user.").Required().String()


	post        = app.Command("post", "Post a message to a channel.")
	postImage   = post.Flag("image", "Image to post.").File()
	postChannel = post.Arg("channel", "Channel to post to.").Required().String()
	postText    = post.Arg("text", "Text to post.").Envar("MYTEXT").Strings()
)

func main() {

	text := strings.Join(*postText, " ")
	println("Pre Parse Post:", text)

	parse, err := app.Parse(os.Args[1:])
	fmt.Println("Parsed vals ", parse)
	mustParse := kingpin.MustParse(parse, err)
	fmt.Println("Parsed full commands", mustParse)

	switch mustParse {
	// Register user
	case register.FullCommand():
		println(*registerNick)

	// Post message
	case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
	}
}