package main

import "fmt"

func main() {
	var son, son1 float32
	var belgi string
	fmt.Println("Birinchi sondi kiriting: ")
	fmt.Scanln(&son)
	fmt.Println("Biron bir amalni yoki belgini kiriting: ")
	fmt.Scanln(&belgi)
	fmt.Println("Ikkinchi sondi kiriting: ")
	fmt.Scanln(&son1)

	if belgi == "*" {
		fmt.Println("Natija: ", son*son1)
	} else if belgi == "/" {
		fmt.Println("Natija: ", son/son1)
	} else if belgi == "+" {
		fmt.Println("Natija: ", son+son1)
	} else if belgi == "-" {
		fmt.Println("Natija: ", son-son1)
	} else {
		fmt.Println("Natija: ", son, belgi, son1)
	}

}
