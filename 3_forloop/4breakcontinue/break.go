package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		} else if i == 6 {
			break
		}
		fmt.Println("I: ", i)
	}

}
