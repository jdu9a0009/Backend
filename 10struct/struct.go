package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

// Constructor
func NewHuman() *Human {
	fmt.Println("Constructor ishga tushdi")
	n := new(Human)
	return n
}

//Parametrli Constructor
func NewHumanWithParams(name, surname string, age int) *Human {
	n := new(Human)
	n.Name = name
	n.Surname = surname
	n.Age = age
	return n
}

func main() {
	/*
		h1 := Human{Name: "John", Surname: "Wick", Age: 24}
		h2 := Human{Name: "Bernard", Surname: "Opel", Age: 34}
		data1 := h1.Name + " " + h1.Surname + " " + strconv.Itoa(h1.Age)
		data2 := h2.Name + " " + h2.Surname + " " + strconv.Itoa(h2.Age)
		fmt.Println(data1)
		fmt.Println(data2)

		var h3 = new(Human)
		h3.Name = "Saud"
		h3.Surname = "Abdulwahed"
		h3.Age = 45
		fmt.Println(h3)
		fmt.Println(h2)
		fmt.Println(h1)
	*/
	h4 := NewHuman()
	fmt.Println(h4)
	h5 := NewHumanWithParams("Asror", "Wick", 56)
	fmt.Println(h5)
}
