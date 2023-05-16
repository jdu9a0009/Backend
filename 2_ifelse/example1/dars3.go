package main

import "fmt"

func main() {
	/*
		// || -> OR -> YOKI
		// || da bir tomoni  true bolsa  qolgan tomoni ham true bo'ladi
		fmt.Println(true || true)   // 1 + 1 = 1
		fmt.Println(true || false)  // 1 + 0 = 1
		fmt.Println(false || true)  // 0 + 1 = 1
		fmt.Println(false || false) // 0 + 0 = 0
		// && -> AND -> VA

		// && da bir tomonida false bolsa natija ham false chiqadi
		fmt.Println(true && true)   // 1 * 1 = 1
		fmt.Println(true && false)  // 1 * 0 = 0
		fmt.Println(false && true)  // 0 * 1 = 0
		fmt.Println(false && false) // 0 * 0 = 0

	*/

	// 80-100 --> 5
	//60-80 --> 4
	//40-60 --> 3
	//40dan past -- Failed

	var ball int
	fmt.Println("Ballni kiritng:")
	fmt.Scanf("%d", &ball)
	if ball > 100 {
		fmt.Println("Tabriklaymiz Grant!!")
	} else if ball >= 80 && ball <= 100 {
		fmt.Println("Baxo: 5 !")
	} else if ball >= 60 && ball < 80 {
		fmt.Println("Baxo: 4 !")
	} else if ball > 40 && ball < 60 {
		fmt.Println("Baxo: 3 !")
	} else {
		fmt.Println("FAILED !")
	}
}
