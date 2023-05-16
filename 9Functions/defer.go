package main

import "fmt"

var isConnected = false

func main() {
	fmt.Println("Connection status:", isConnected)
	dbProcessing()
	fmt.Println("Connection status:", isConnected)
}
func dbProcessing() {
	defer disconnect()
	connect()
	fmt.Println("Defering disconnected !")
	fmt.Println("Connection status: ", isConnected)
	fmt.Println("Blah blah blah!!!")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database !!")
}
func disconnect() {
	isConnected = false
	fmt.Println("Disconnected !!!")
}
