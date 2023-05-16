package main

import "fmt"

func main() {
	son1 := 0
	son2 := 0
	fmt.Print("Sonni kiriting: ")
	fmt.Scan(son1)

	fmt.Print("Sonni kiriting: ")
	fmt.Scan(son2)
	if son2 == 0 {
		fmt.Println("xech bir sondi 0  ga bo'lib bo'lmaydi")
	} else {
		fmt.Println(son1 / son2)
	}

}
