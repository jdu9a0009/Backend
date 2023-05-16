package main

import "fmt"

func main() {
	var son int
	fmt.Println("ILtmos butun son kirting")
	fmt.Scanf("%d", &son)

	if son > 0 {
		fmt.Println("Son: ", son+1)
	} else {
		fmt.Println("2ga kamaytrldi: ", son-2)
	}
}
