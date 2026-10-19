package main

import (
	"fmt"
	"study/user"
)

func main() {
	u := user.NewUser("Alice", 30)
	fmt.Println(u, "u")

	u.SetName("Bob")
	u.SetAge(25)
	fmt.Println(u, "u")
	name := u.GetName()
	age := u.GetAge()
	fmt.Println(name, "name")
	fmt.Println(age, "age")
}
