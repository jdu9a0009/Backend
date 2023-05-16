package main

import "fmt"

func main() {
	var hafta int
	fmt.Println("Hafta kunini kirting (1-7")
	fmt.Scanf("%d", &hafta)
	switch hafta {
	case 1:
		fmt.Println("Bugun  ish kuni")
	case 2:
		fmt.Println("Bugun  ish kuni")
	case 3:
		fmt.Println("Bugun  ish kuni")
	case 4:
		fmt.Println("Bugun  ish kuni")
	case 5:
		fmt.Println("Bugun  ish kuni")

	case 6:
		fmt.Println("Bugun  dam olish kuni")
	case 7:
		fmt.Println("Bugun  dam olish kuni")
	default:
		fmt.Println("Iltmos hafta kunini kirting bunday kun mavjud emas")
	}

}
