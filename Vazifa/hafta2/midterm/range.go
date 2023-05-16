package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("3ga ham 5 ga ham")
		} else if i%3 == 0 {
			fmt.Println("3ga bo'linadi")
		} else if i%5 == 0 {
			fmt.Println("5ga bo'linadi")
		} else {
			fmt.Println(i, ",")
		}

	}

}
