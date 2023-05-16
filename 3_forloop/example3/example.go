package main

import "fmt"

// Kirtlgan sondan 0 gacha bolgan sonlar 4ga bolinshini chiqaring

func main() {
	son := 0
	fmt.Printf("Sonni kiriting: ")
	fmt.Scanf("%d", &son)

	for ; son > 0; son-- {
		if son%4 == 0 {
			fmt.Println(son, " Bo'linadi")
		}
	}

}
