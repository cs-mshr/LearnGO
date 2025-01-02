package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result: ", result)

}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return -1, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}
