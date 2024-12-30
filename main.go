package main

import (
	"log"
	"sort"
)

func main() {
	var myslice []string
	myslice = append(myslice, "hero", "john", "a")

	sort.Strings(myslice)
	log.Println(myslice)

	numbers := []int{1, 2, 3, 4, 5}
	sort.Ints(numbers)
	log.Println(numbers[:4])
	for val := range numbers {
		log.Println(val)
	}
}
