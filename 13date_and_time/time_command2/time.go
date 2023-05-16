package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	xTime := time.Date(1000, 10, 10, 10, 10, 10, 100, time.UTC)
	p(xTime)
	fmt.Println("*---------------------------------*")
	now := time.Now()
	p(now.Year())
	p(now.Month(), "shu")
	p(now.Weekday())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	// check date and clock
	p(xTime.Before(now))
	p(xTime.After(now))
	p(xTime.Equal(now))

	//See differrence
	different := now.Sub(xTime)
	p(different.Hours())

	p(xTime.Add(different))
	p(xTime.Add(-different))
	nimadir := xTime.Month()
	fmt.Println(nimadir)

}
