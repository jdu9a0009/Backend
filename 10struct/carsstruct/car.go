package main

import "fmt"

func main() {
	myCar := car{Make: "Volvo", Model: "XC90", Color: "White", Year: 2022, Weight: 2343, Engine: engine{name: "T8", hp: 400}}
	fmt.Println(myCar)
}

type car struct {
	Make, Model, Color string
	Year, Weight       int
	Engine             engine
}
type engine struct {
	name string
	hp   int
}
