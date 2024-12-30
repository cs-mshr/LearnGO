package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
	BirthDate time.Time
}

func main() {

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Phone:     "123456789",
		//BirthDate: time.Now(),
	}

	log.Println(user.FirstName, user.BirthDate)
}
