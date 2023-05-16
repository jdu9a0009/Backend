package main

import "fmt"

func main() {
	//JUFT  YOKI TOQ SON ANIQLAYDIGON DASTUR
	son := 0
	fmt.Println("Juft yoki Toq ekanligini aniqlamoqchi bo'lgan soningizni kiriting")
	fmt.Scanf("%d", &son)
	if son%2 == 0 {
		fmt.Println(son, "Juft !!")
	} else {
		fmt.Println(son, "Toq!!")
	}
}
