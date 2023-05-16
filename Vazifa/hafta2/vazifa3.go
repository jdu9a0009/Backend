package main

import "fmt"

func main() {
	var son int
	fmt.Println("Iltmos butun son kirting:  ")
	fmt.Scanf("%d", &son)

	if son%10 == 0 && son%20 != 0 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
