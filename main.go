package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := "Enter all this data into a file"

	file, err := os.Create("./log.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("log.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	log.Println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
