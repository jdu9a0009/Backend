package main

import (
	"fmt"
	"os"
)

func try() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error bor: ", err)
	} else {
		fmt.Println("Fayl: ", file.Name())
	}

	fayl, err := os.Open("misol.txt")
	checkError(err)
	fmt.Println("File: ", fayl.Name())

}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}
