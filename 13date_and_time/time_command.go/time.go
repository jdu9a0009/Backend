package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, time.May, 1, 10, 20, 30, 100, time.UTC)
	fmt.Printf("Time: %v\n", t)
	fmt.Println("*---------------------------------*")
	now := time.Now()
	fmt.Println("The time of Now: ", now)
	fmt.Println("*---------------------------------*", now.YearDay())

	fmt.Println("Oy: ", now.Month())
	fmt.Println("Kun: ", now.Day())
	fmt.Println("Hafta kunlar: ", now.Weekday())
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("Ertangi kun: ", tomorrow.Day())

}
