package pointers

import "fmt"

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}
func Email(u User) string {
	return u.email
}

func pointers() {
	user := User{
		email:    "email@email.com",
		username: "username",
		age:      1,
	}
	fmt.Println(user.Email())
}
