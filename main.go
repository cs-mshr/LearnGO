package main

import "log"

func main() {
	//var myOtherMap map[string]string
	myMap := make(map[string]string)

	myMap["Dog"] = "Fido"
	myMap["Cat"] = "Garfield"
	myMap["Bird"] = "Tweety"

	for key, value := range myMap {
		println(key, " = ", value)
	}

	//Integer Map
	myIntMap := make(map[string]int)
	log.Println(myIntMap["Dog"]) // default is zero

	myIntMap["Dog"]++
	log.Println(myIntMap["Dog"]) // zero -> 1

	//Map can have no specified value type also
	myMapWithoutValueType := make(map[string]interface{})
	myMapWithoutValueType["Dog"] = "Fido"
	myMapWithoutValueType["Cat"] = 1
	myMapWithoutValueType["Bird"] = 7.43
	for key, value := range myMapWithoutValueType {
		log.Println(key, " = ", value)
	}
}
