package main

import (
	"fmt"

	"github.com/fatih/color"
)

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func NewUser(name string, age int, phoneNumber string, isClose bool, rating float64) User {
	if name == "" || age <= 0 || age >= 100 || phoneNumber == "" || rating < 0 || rating > 10 {
		return User{}
	} else {
		return User{
			Name:        name,
			Age:         age,
			PhoneNumber: phoneNumber,
			IsClose:     isClose,
			Rating:      rating,
		}
	}
}

func (user User) Greeting() {
	fmt.Println("My name is", user.Name)
	fmt.Println("I am", user.Age, "years old")
	fmt.Println("My phone number is", user.PhoneNumber)
}

func (user *User) ChangeName(name string) {
	if name != "" {
		user.Name = name
	} else {
		color.Red("Name is empty")
	}
}

func (user *User) ChangeAge(age int) {
	if age > 0 && age < 100 {
		user.Age = age
	} else {
		color.Red("Age is not valid")
	}
}

func (user *User) ChangePhoneNumber(phoneNumber string) {
	if phoneNumber != "" {
		user.PhoneNumber = phoneNumber
	} else {
		color.Red("Phone number is empty")
	}
}

func (user *User) CloseAccount() {
	user.IsClose = true
}
func (user *User) OpenAccount() {
	user.IsClose = false
}

func (user *User) RatingUp(rating float64) {
	if rating >= 0 && user.Rating+rating <= 10 {
		user.Rating = rating
	} else {
		color.Red("Rating is not valid, it should be between 0 and 10")
	}
}

func (user *User) Ratingdown(rating float64) {
	if rating >= 0 && user.Rating-rating >= 0 {
		user.Rating = rating
	} else {
		color.Red("Rating is not valid, it should be between 0 and 10")
	}
}

func main() {
	user := NewUser("John", 25, "123-456-7890", false, 8.5)
	fmt.Println("user", user)
	user2 := NewUser("Jane", 30, "098-765-4321", false, 3.0)
	fmt.Println("user2", user2)
}
