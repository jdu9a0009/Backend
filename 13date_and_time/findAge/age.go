package main

import (
	"fmt"
	"time"
)

func calculateAge(birthYear, birthMonth, birthDay int) (int, int) {
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())
	currentDay := time.Now().Day()

	age := currentYear - birthYear

	if birthMonth > currentMonth || (birthMonth == currentMonth && birthDay > currentDay) {
		age--
	}

	months := currentMonth - birthMonth
	if months < 0 {
		months += 12
	}

	return age, months
}

func main() {
	var birthYear, birthMonth, birthDay int

	fmt.Println("Enter your birth year:")
	_, err := fmt.Scanf("%d", &birthYear)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid year.")
		return
	}

	fmt.Println("Enter your birth month:")
	_, err = fmt.Scanf("%d", &birthMonth)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid month.")
		return
	}

	fmt.Println("Enter your birth day:")
	_, err = fmt.Scanf("%d", &birthDay)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid day.")
		return
	}

	age, months := calculateAge(birthYear, birthMonth, birthDay)
	fmt.Printf("Your age is %d years and %d months.\n", age, months)
}
