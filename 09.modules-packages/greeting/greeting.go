package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello, World!")

	j := GiveMeInt()
	fmt.Println("j", j)
}
