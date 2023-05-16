package main

import "fmt"

func main() {
	// 80-100 --> 5
	//60-80 --> 4
	//40-60 --> 3
	//20-40dan -> 2
	// else "Siz bebahosiz"

	var ball int
	fmt.Println("Ballni kiritng:")
	fmt.Scanf("%d", &ball)
	if ball > 100 {
		fmt.Println("Siz o'ta bebaxosiz")
	} else if ball >= 80 && ball <= 100 {
		fmt.Println("Baxo: 5 !")
	} else if ball >= 60 && ball < 80 {
		fmt.Println("Baxo: 4 !")
	} else if ball >= 40 && ball < 60 {
		fmt.Println("Baxo: 3 !")
	} else if ball > 20 && ball < 40 {
		fmt.Println("Baho: 2 !")
	} else {
		fmt.Println("Siz bebahosiz ")
	}
}
