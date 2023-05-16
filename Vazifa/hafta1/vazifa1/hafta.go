package main

import "fmt"

func main() {
	var hafta int
	fmt.Println("Hafta kunini kiriting,men sizga u qaysi kun ekanligini aytaman")
	fmt.Scanln(&hafta)
	switch hafta {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Iltimos hafta kunlari 7 kundan iborat  unda boshqa son kiritmang")
	}

}
