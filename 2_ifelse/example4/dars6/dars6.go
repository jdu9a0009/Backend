package main

import "fmt"

func main() {
	// CHAKAN SAVDO DOKONI UCHUN MAXSUS DASTUR
	// Apelsin
	//1-10 -> 3$
	//11-100 ->2.75$
	// 101-1000->2.5$
	//1001-10 000-->2$
	//10 001-...-->1.25$

	var dona, donaNarx = 0, 0.0
	fmt.Println("Nechda dona Apelsin olishni istaysiz: ")
	fmt.Scanf("%d", &dona)

	if dona > 10000 {
		donaNarx = 1.25
	} else if dona > 1000 && dona <= 10000 {
		donaNarx = 2
	} else if dona > 100 && dona <= 1000 {
		donaNarx = 2.5
	} else if dona > 10 && dona <= 100 {
		donaNarx = 2.75
	} else if dona > 0 && dona <= 10 {
		donaNarx = 3
	} else {
		fmt.Println("Iltimos faqat Musbat son kiriting,xozirda  qarzga savdo yo'q")
	}

	var jamiNarx float32
	jamiNarx = float32(dona) * float32(donaNarx)
	fmt.Println("Jami  to'lov summasi: ", jamiNarx, "$")

}
