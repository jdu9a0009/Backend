package main

import (
	"fmt"

	"github.com/Jdustudent/Interface/special"
	"github.com/Jdustudent/Interface/wellknown"
)

func main() {
	f1 := new(car.Ferrari)
	f1.Brand = "F45"
	f1.Model = "Ferrari 1"
	f1.Color = "Red"
	f1.Speed = 300
	f1.Price = 1.2
	f1.Special = true

	m1 := new(mycar.BMW)
	m1.Brand = "M035"
	m1.Model = "BMW i5 edrive"
	m1.Color = "Black"
	m1.Speed = 260
	m1.Price = 0.6
	m1.Special = false

	f2 := car.NewFerrari("Formula", "Ferrari 5", "Black", 600, 3.2, true)
	m2 := mycar.NewBMW("BMW", "BMW i5 edrive", "White", 460, 1.6, true)
	fmt.Println(f2)
	fmt.Println(m2)
	car.CarExecute(f1)
	mycar.CarExecute(m1)

}
