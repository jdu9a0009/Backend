package main

import "fmt"

func main() {
	input := 0
	output := 0
	fmt.Println("5 xonalik istalgan sondi kiriting")
	fmt.Scanf("%d", &input)
	if input > 9999 && input < 100000 {
		for input != 0 {
			num := input % 10
			output = output*10 + num
			input /= 10
		}
		fmt.Printf("Kiritilgan sonning teskari xolati:%d\n ", output)
	} else {
		fmt.Println("Iltmos 5 xonalik son kirting")
	}

}
