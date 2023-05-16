package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c -->15
}
func main() {
	//Channellar bizga go dan olgan qiymatizmizni boshqarishda yordam  beradi uni <- orqalik ifodalanadi

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)     /// created channel like map
	go sum(s[:], c)         //15
	go sum(s[:len(s)/2], c) //3
	x, y := <-c, <-c        // receive from c
	fmt.Println(x, y, x+y)

}
