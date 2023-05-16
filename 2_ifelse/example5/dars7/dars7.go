package main

import "fmt"

func main() {
	//Bir biriga qoldiqsiz bolinish yoki bolinmasligini aniqlaydgon dastur
	kichikson, kattason := 0, 0
	fmt.Println("Katta sondi kiriting: ")
	fmt.Scanf("%d", &kattason)
	fmt.Println("Kichik sondi kiritng")
	fmt.Scanf("%d", &kichikson)

	if kattason%kichikson == 0 {
		fmt.Println("Bo'linadi")
	} else {
		fmt.Println("Bo'linmaydi !!")
	}

}
