package types

import "encoding/xml"

type (
	User struct {
		Username string `json:"username" xml:"secrets>username"`
		Password string `json:"password" xml:"secrets>password"`
		Email    string `json:"email" xml:"email"`
	}
	UsersDb struct {
		Id int `json:"id" xml:"id,attr"`
		XMLName xml.Name `json:"-" xml:"users"`
		Type  string `json:"type" xml:"type"`
		Users []User `json:"users" xml:"user"`
	}
)
