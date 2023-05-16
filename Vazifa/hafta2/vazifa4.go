package main

import "fmt"

func main() {
	input := 0

	fmt.Println("4 xonalik istalgan sondi kiriting")
	fmt.Scanf("%d", &input)
	if input > 999 && input < 10000 {

		tort := input / 1000
		ming := (input - tort*1000) / 100
		yuz := ((input - tort*1000) / 10) % 10
		bir := input % 10

		// fmt.Printf("tort: %d\n ", tort)
		// fmt.Printf("ming: %d\n ", ming)
		// fmt.Printf("yuz: %d\n ", yuz)
		// fmt.Printf("bir: %d\n ", bir)
		if tort > ming {
			fmt.Println(tort)
		} else if ming > yuz {
			fmt.Println(ming)
		} else if yuz > bir {
			fmt.Println(bir)
		}
	} else {
		fmt.Println("Iltmos 4 xonalik son kirting")
	}

}
