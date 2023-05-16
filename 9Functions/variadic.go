package main

import "fmt"

func main() {
	natija := yigindiHisobla(12, 23, 34, 45, 67, 78, 99)
	fmt.Println("Natija: ", natija)

}
func yigindiHisobla(sonlar ...int) (yigindi int) {
	for _, son := range sonlar {
		yigindi += son
	}
	return
}

// variadic degani bu funskiyaga istalgancha qiymat berish mumkin degani  aslida uni slice arraylar bn ham xal etsa boladi lekin eng tez oson yoli shu vaariadic funksyadur
