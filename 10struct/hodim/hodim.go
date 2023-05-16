package main

import "fmt"

type Employee struct {
	Name, Surname, Proffesion string
	Age, Salary               int
}

// Constructor
func NewEmployee() *Employee {
	fmt.Println("Constructor ishga tushdi")
	n := new(Employee)
	return n
}

//Parametrli Constructor
func NewEmployeeWithParams(name, surname string, age, salary int) *Employee {
	mathTeacher := new(Employee)
	mathTeacher.Name = name
	mathTeacher.Surname = surname
	mathTeacher.Age = age
	mathTeacher.Salary = salary
	return mathTeacher
}

func main() {
	/*
		// var teacher Employee
		teacher := Employee{Name: "Islombek", Surname: "Malikov", Proffesion: "English Teacher", Age: 32, Salary: 3000}
		director := Employee{"Dovud", "Temirov", "Director", 28, 5000}
		fmt.Println(teacher)
		fmt.Println(director)
	*/

	h4 := NewEmployee()
	fmt.Println(h4)
	h5 := NewEmployeeWithParams("Asadbek", "Manopov", 34, 4000)
	fmt.Println(h5)
}
