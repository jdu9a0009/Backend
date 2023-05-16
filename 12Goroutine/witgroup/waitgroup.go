package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go salomBer()
		go xayrlash()
	}
	wg.Wait()
}

func salomBer() {
	fmt.Println("Salom !!!")
	wg.Done()
}

func xayrlash() {
	fmt.Println("Xayr Ko'rishguncha !!")
	wg.Done()
}
