package main

import "fmt"

func main() {
	message := "First message"
	changeMessage(&message)
	fmt.Println("Message: ", message)
}

func changeMessage(msg *string) {
	*msg = "Second Message"
}
