package main

import "fmt"

func main() {
	countries := [...]string{"UK", "US", "UZ", "Japan"}
	for i, v := range countries {
		fmt.Println("INDEX: ", i, "VALUE: ", v)
	}

	for _, v := range countries {
		fmt.Println("VALUE: ", v)
	}

	numbers := [...]int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	yigindi := 0
	faktorial := 1

	for _, value := range numbers {
		yigindi += value
		faktorial *= value
		if value%2 == 0 {
			fmt.Println("Juft Son: ", value)

		}

	}
	fmt.Println("Yigindi: ", yigindi)
	fmt.Println("Faktorial: ", faktorial)

}
