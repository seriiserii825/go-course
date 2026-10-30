package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

type User struct {
	Name     string
	Ballance int
}

func NewUser(name string, balance int) User {
	return User{
		Name:     name,
		Ballance: balance,
	}
}

func Pay(user *User, usd int) bool {
	if user.Ballance-usd >= 0 {
		user.Ballance -= usd
		return true
	}
	return false
}

func main() {

	user := NewUser("Oleg", 10)

	pp.Println(user)
	payed := Pay(&user, 50)
	if payed {
		pp.Println(user)
	} else {
		fmt.Println("No payed")
	}
}
