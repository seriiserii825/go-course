package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func (user User) Greeting() {
	fmt.Println("My name is", user.Name)
	fmt.Println("I am", user.Age, "years old")
	fmt.Println("My phone number is", user.PhoneNumber)
}

func (user *User) SetName(name string) {
	if name != "" {
		user.Name = name
	} else {
		fmt.Println("Name is empty")
	}
}

func (user *User) SetAge(age int) {
	if age > 0 && age < 100 {
		user.Age = age
	} else {
		fmt.Println("Age is not valid")
	}
}

func (user *User) SetPhoneNumber(phoneNumber string) {
	if phoneNumber != "" {
		user.PhoneNumber = phoneNumber
	} else {
		fmt.Println("Phone number is empty")
	}
}

func (user *User) SetIsClose(isClose bool) {
	user.IsClose = isClose
}

func (user *User) SetRating(rating float64) {
	if rating >= 0 && user.Rating+rating <= 10 {
		user.Rating = rating
	} else {
		fmt.Println("Rating is not valid, it should be between 0 and 10")
	}
}

func main() {
	user := User{
		Name:        "Serii",
		Age:         64,
		PhoneNumber: "+37360562168",
		IsClose:     true,
		Rating:      4.5,
	}
	fmt.Println("User", user)
	user.Greeting()
	user.SetName("John")
	user.SetAge(30)
	user.SetPhoneNumber("+1234567890")
	user.SetIsClose(false)
	user.SetRating(2.5)
	fmt.Println("user", user)
}
