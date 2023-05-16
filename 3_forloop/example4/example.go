package main

// faktarialni aniqlash

import "fmt"

func main() {
	son := 0
	fmt.Println("Iltimos Son Kiriting")
	fmt.Scanf("%d", &son)
	yigindi := 0
	faktorial := 1
	for i := 1; i <= son; i++ {
		yigindi += i
		faktorial += faktorial
	}
	fmt.Println("Yigindi: ", yigindi)
	fmt.Println("Faktoriyal: ", faktorial)
}
