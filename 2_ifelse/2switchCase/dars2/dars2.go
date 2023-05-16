package main

import "fmt"

func main() {
	//Switch
	/*
		fmt.Println("1.Balans\n 2.Sms\n3.MB")
		var tanlov int
		fmt.Println("Tanlang: ")
		fmt.Scanf("%d", &tanlov)
		switch tanlov {
		case 1:
			fmt.Println("Balansingizda 43$ mavjud!")
		case 2:
			fmt.Println("67 ta SMS mavjud !")
		case 3:
			fmt.Println("789 MB mavjud !")
		default:
			fmt.Println("Notogri amal tanlandi !!")
		}
	*/
	// var ism string
	// fmt.Println("Ismingizni Bosh xarfini kiriting !")
	// fmt.Scanf("%v", &ism)
	// switch ism {
	// case "a":
	// 	fmt.Println("Salom Anvar !")
	// case "b":
	// 	fmt.Println("Salom Behruz")
	// case "t":
	// 	fmt.Println("Salom Temur")
	// default:
	// 	fmt.Println("Bunday ism mavjud emas !")
	// }

	// bankamatdan pul yechish jarayoni
	// var cash = 0
	// fmt.Println("Bankamatdan necha so'm pul chiqarmqochisiz tanlang :\n 1: 50 000 \n 2: 100 000\n 3:200 000\n 4:500 000")
	// fmt.Scanf("%v", &cash)
	// switch cash {
	// case 1:
	// 	fmt.Println("Siz 50 000 som pul oldingiz")
	// case 2:
	// 	fmt.Println("Siz 100 000 som pul oldingiz")
	// case 3:
	// 	fmt.Println("Siz 200 000 som pul oldingiz")
	// case 4:
	// 	fmt.Println(" Siz 500 000 som pul oldingiz")
	// default:
	// 	fmt.Println(" Notogri tanlov")
	// }
	const number = 5
	if number > 0 {
		println("Bu Musbat Son")
	} else if number == 0 {
		println("Bu musbat ham manfiy ham emas,bu Nolga teng")
	} else {
		fmt.Println("Bu manfiy son")
	}

	fmt.Println("If blo'kidan kegingi ko'd")

}
