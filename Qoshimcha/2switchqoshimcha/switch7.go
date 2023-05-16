package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("a: ")
	fmt.Scanln(&a)

	fmt.Println("og'irlik ")
	fmt.Scanf("%d", &b)
	switch a {
	case 1:
		fmt.Println("kgram ", b)
	case 2:
		fmt.Println("Milligram: ", b/1000000)
	case 3:
		fmt.Println("gram: ", b/1000)
	case 4:
		fmt.Println("Tonna: ", b*1000)
	case 5:
		fmt.Println("sentner: ", b*100)
	}

}
