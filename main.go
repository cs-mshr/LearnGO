package main

import "log"

func main() {
	isTrue := true

	if isTrue {
		log.Println("Variable is :", isTrue)
	} else {
		log.Println("Variable isn't True")
	}

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"cat", "dog", "fish"}
	for _, animal := range animals {
		log.Println(animal)
	}

}
