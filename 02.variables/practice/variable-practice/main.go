package main

import "fmt"

func main() {
	name := "Sergey"
	age := 50
	language := "Go"
	experience := 8
	isLearning := true
	course := "Go for Beginners"

	fmt.Println("Course:", course)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Language:", language)
	fmt.Println("Experience:", experience, "years")
	fmt.Println("Is learning:", isLearning)

	experience += 1
	fmt.Println("Updated experience:", experience, "years")

}
