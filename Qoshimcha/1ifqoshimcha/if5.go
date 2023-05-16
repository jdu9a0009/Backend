package main

import "fmt"

func main() {
	a := -2
	b := -3
	c := -4

	if a > 0 && b > 0 && c > 0 {
		fmt.Println("3ta musbat son")
		fmt.Println("0ta manfiy son")
	} else if a < 0 && b < 0 && c < 0 {
		fmt.Println("0ta musbat son")
		fmt.Println("3ta musbat son")
	} else if a > 0 && b > 0 {
		fmt.Println("2ta musbat son")
		fmt.Println("1ta manfiy son")
	} else if b > 0 && c > 0 {
		fmt.Println("2ta musbat son")
		fmt.Println("1ta manfiy son")
	} else if a > 0 && c > 0 {
		fmt.Println("2ta musbat son")
		fmt.Println("1ta manfiy son")
	} else if a < 0 && b < 0 || b < 0 && c < 0 || a < 0 && c < 0 {
		fmt.Println("1ta musbat son")
		fmt.Println("2ta manfiy son")
	}
}
