package main

import "fmt"

func main() {
	n := 2
	var res float32
	for i := 1; i <= n; i++ {
		res += 1 / float32(i)
	}
	fmt.Println(res)

}
