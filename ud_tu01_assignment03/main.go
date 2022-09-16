package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readByIoutils(fileName string) {

	println("readByIoutils\n============")

	bytes, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Println("Failed to read file: ", fileName)
		os.Exit(1)
	}
	fmt.Println(string(bytes))

	println("\n============\n\n")
}

func readByOs(fileName string) {

	println("readByOs\n============")

	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println("Failed to read file: ", fileName)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

	println("\n\n============\n\n")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name and rerun...")
		os.Exit(1)
	} else {
		readByIoutils(os.Args[1])
		readByOs(os.Args[1])
	}
}
