package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hozirgi vaqt: %v\n", time.Now())
	fmt.Printf("Hozirgi vaqt Unix: %v\n", time.Now().Unix())

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Soniya: ", time.Now())

	}
}
