package main

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		"1", "木兮", "123456",
	},
	"2": {
		"2", "兮木", "654321",
	},
}
