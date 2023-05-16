package main

//simple program  calculate your age

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 0
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println("Please enter birth year")
	time.Sleep(2 * time.Second)
	fmt.Scan(&birthYear)
	fmt.Println("Your age: ", now.Year()-birthYear)

}
