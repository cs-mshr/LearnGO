package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")

	var message string
	var i int

	message = "Good bye Java"
	fmt.Println(message)

	i = 10
	fmt.Println("i is set to:", i)

	message2, m3 := saySomething()
	fmt.Println(message2, " ", m3)
}

func saySomething() (string, string) {
	return "what i should return", "else"
}
