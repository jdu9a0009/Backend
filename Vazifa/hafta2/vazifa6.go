package main

import "fmt"

func main() {
	var oy int
	fmt.Println("Qaysi oyda necha kun borligini bilmoqchisz")
	fmt.Scanf("%d", &oy)

	switch oy {
	case 1:
		fmt.Println("31 kun")
	case 3:
		fmt.Println("31 kun")
	case 5:
		fmt.Println("31 kun")
	case 7:
		fmt.Println("31 kun")
	case 8:
		fmt.Println("31 kun")
	case 10:
		fmt.Println("31 kun")
	case 12:
		fmt.Println("31 kun")
		break
	case 4:
		print("30 kun")
	case 6:
		print("30 kun")
	case 11:
		print("30 kun")

	case 2:
		print("28/29 kun")
	default:
		fmt.Println("Bunday oy yo'q")

	}

}
