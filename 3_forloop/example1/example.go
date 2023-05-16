package main

import "fmt"

func main() {
	// count := 0
	// fmt.Println("Iltmos son kirting")
	// fmt.Scanf("%d", &count)
	// for i := 1; i < count; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Juft Son: ", i)
	// 	} else {
	// 		fmt.Println("Toq son", i)
	// 	}
	// }

	// teskari sanaydgan

	from := 0
	fmt.Println("Iltmos son kirting !!")
	fmt.Scanln(&from)

	for i := from; i > 0; i-- {
		fmt.Println("ortga: ", i)
	}
}
