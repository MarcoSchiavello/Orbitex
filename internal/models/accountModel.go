package models

import "encoding/json"

type User struct {
	UserId    int    `json:"userId"`
	Username  string `json:"username"`
	Pwd       string `json:"pwd"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Avatar    string `json:"avatar"`
	Mobile    string `json:"mobile"`
	Favorites []int  `json:"favorites"`
}

func (user *User) toJSON() (string, error) {
	str, err := json.Marshal(user)

	return string(str), err
}

func UserFromJSON(userJson string) (User, error) {
	var user User
	err := json.Unmarshal(([]byte)(userJson), &user)

	return user, err
}

func AllUsers() {

}
