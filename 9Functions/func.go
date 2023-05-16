package main

import "fmt"

func main() {
	sayHi()
	sayMyName("Sarvarbek")
	fmt.Println(add(10, 20, "")) // o'zgaruvchga tenglamasdan ham  parametrlar ustda amal bajarsa bo'ladi
	natija := add(100, 200, "")
	fmt.Println("Natije: ", natija)
}
func sayHi() {
	fmt.Println("me: Hello evrybody I am new Gopher")
}

// ! go da 2 xil funksya bor parametrli va parametrsz yuqordagi funksya  parametrsiz  pastragsi esa parametrli funksya
// *name bu  argument yoki parametr deb ataladi
//parametr  berilib lekin uni ishlatmay ketilsa ham bo'ladi.
func sayMyName(name string) {
	fmt.Println("Gophers: Hi", name, "you are welcome")
}
func add(n1, n2 int, ism string) int {
	return n1 + n2
}
