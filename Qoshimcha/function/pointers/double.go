package main

//Write a Go function that takes a pointer to an integer as its argument and doubles the value of the integer.
import "fmt"

func main() {
	var number int
	fmt.Println("Istalgan sondi kiriting men uni 2 barobr qilib  qaytaraman")
	fmt.Scanf("%d", &number)

	double(&number)
}

func double(son *int) {
	*son = *son * 2
	fmt.Println("2 barobar: ", *son)
}
