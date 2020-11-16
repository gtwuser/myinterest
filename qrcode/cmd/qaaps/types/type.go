package types

import "encoding/xml"

type (
	User struct {
		Username string `json:"username" xml:"secrets>username"`
		Password string `json:"password" xml:"secrets>password"`
		Email    string `json:"email" xml:"email"`
	}
	UsersDb struct {
		Id      int      `json:"id" xml:"id,attr"`
		XMLName xml.Name `json:"-" xml:"users"`
		Type    string   `json:"type" xml:"type"`
		Users   []User   `json:"users" xml:"user"`
	}
)

func (u *User) EncodeAsString() []string {
	us := make([]string, 3)
	us[0] = u.Email
	us[1] = u.Password
	us[2] = u.Username
	return us
}

func (u *User) FromCSV(ss []string)  {
	u.Email = ss[0]
	u.Password = ss[1]
	u.Username = ss[2]
}