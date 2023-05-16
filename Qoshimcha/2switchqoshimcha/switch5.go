package main

import "fmt"

func main() {
	var a, b float64
	var c int
	fmt.Println("Please enter first number: ")
	fmt.Scanln(&a)

	fmt.Println("Please enter Second number: ")
	fmt.Scanf("%f", &b)

	fmt.Println("Please enter your option : ")
	fmt.Println("1 (+);\n2(-);\n3(/);\n4(*);\n")
	fmt.Scanf("%d", &c)

	switch c {
	case 1:
		fmt.Println("result: ", a+b)
	case 2:
		fmt.Println("result: ", a-b)
	case 3:
		fmt.Println("result: ", a/b)
	case 4:
		fmt.Println("result: ", a*b)
	}
}
