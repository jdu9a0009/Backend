package main

import "time"

// for loopni tugash qismi aniq bo'lganda ishlatiladi
// while loop esa tugashi aniq bolmaganda misol foydalanuvcgi nimadir kirtgandagini  tsikl tugashi kerak bolsa while ishlatiladi
func main() {

	i := 0
	sum := 0
	for i < 5 {
		println("I: ", i)
		sum += i
		i++
		time.Sleep(500 * time.Millisecond)
	}
	println("Yigindi: ", sum)

}
