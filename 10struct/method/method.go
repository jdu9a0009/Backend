package main

import "fmt"

///Struct
type User struct {
	ID, Age                 int
	Name, Surname, Username string
	Pay                     *Payment
}

///Constructor
func NewUser() *User {
	var n = new(User)
	return n
}

//Methodlar
func (u User) GetFullName() string {
	return u.Name + " " + u.Surname
}

///Struct
type Payment struct {
	Salary, Bonus float64
}

//Methodlar
func (p Payment) GetSalaryWithBonus() float64 {
	return p.Salary + p.Bonus
}

func main() {
	u1 := User{
		ID:       1,
		Age:      34,
		Name:     "John",
		Surname:  "Wick",
		Username: "Jwick",
		Pay:      &Payment{Salary: 1000, Bonus: 300},
	}
	fmt.Println(u1.Age)
	fmt.Println(u1.Pay.Salary)
	fmt.Println(u1.GetFullName())
	fmt.Println(u1.Pay.GetSalaryWithBonus())
}
