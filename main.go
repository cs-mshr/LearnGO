package main

import "log"

func main() {
	var myString string
	myString = "Green"
	println(myString)

	changeString(&myString)
	println(myString)
}

func changeString(s *string) {
	log.Println("the string address is: ", s)
	log.Println("the string value is: ", *s)
	newValue := "Blue"
	*s = newValue
}
