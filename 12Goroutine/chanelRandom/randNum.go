package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	go getRandomNumber(channel)
	for randomNumber := range channel {
		fmt.Println("Tasodify son: ", randomNumber)
	}

}

func getRandomNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		number := rand.Intn(1000)
		time.Sleep(time.Second * 1)
		channel <- number

	}
	close(channel)
}
