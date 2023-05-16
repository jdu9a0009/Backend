package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("a: ")
	fmt.Scanln(&a)

	fmt.Println("Uzunlik ")
	fmt.Scanf("%d", &b)
	switch a {
	case 1:
		fmt.Println("Desimetr: ", b*10) // 1 metr= 10 desimetr
	case 2:
		fmt.Println("Kilometr: ", b/1000) //1km=1000m
	case 3:
		fmt.Println("Metr: ", b*1)
	case 4:
		fmt.Println("millimetr: ", b*10000) // 1m=10 000 millimetr  1m= 1000 sm ,1 sm=	10millimetr
	case 5:
		fmt.Println("santimetr: ", b*1000) //1m= 1000 sm
	}

}
