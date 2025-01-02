package main

import (
	"awesomeProject/helpers"
	"log"
)

const numPool = 100

func CalValues(intChan chan int) {
	randomNumber := helpers.RandonNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalValues(intChan)

	val := <-intChan
	log.Println(val)
}
