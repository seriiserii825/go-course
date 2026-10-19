package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/09.modules-packages/user"
)

func main() {
	u := user.NewUser("Alice", 30)
	fmt.Println(u, "u")

	pp.Println(u)
}
