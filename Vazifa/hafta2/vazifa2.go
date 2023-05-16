package main

import "fmt"

func main() {
	var son int
	fmt.Println("Iltmos butun son kirting:  ")
	fmt.Scanf("%d", &son)

	if son%3 == 0 || son%5 == 0 || son%7 == 0 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
