package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("A sonini kiriting")
	fmt.Scanln(&a)
	fmt.Println("B sonini kiritng")
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("A katta !")
	} else if a == b {
		fmt.Println("A B ga  teng !!!")
	} else {
		fmt.Println("B katta !!")
	}
	/*
		if son := 23; son>0 {
			 fmt.Println("Son pozitiv")
		} else {
			fmt.Println("Son negativ")
		}
		bu usul faqatgina Go da bor

	*/
	// 	const number = 5
	// 	if number > 0 {
	// 		println("Musbat son")
	// 	} else if number == 0 {
	// 		println("Bu musbat ham manfiy ham emas,bu Nolga teng")
	// 	} else {
	// 		fmt.Println("Bu manfiy son")
	// 	}

	// 	fmt.Println("If blo'kidan kegingi ko'd")
	// }
	// bankamatdan pul yechish jarayoni
	const cash = 0
	switch cash {
	case condition:

	}
}
