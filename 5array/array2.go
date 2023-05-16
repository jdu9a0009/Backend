package main

import "fmt"

func array() {
	myArray := [...]int{22, 3, 4, 652, 8, 2} //  size ga "..." qoyish orqalik unga istalgan miqdordagi qiymatni kirtsak bo'ladi
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	// listning uzunligini chiqaradi
	fmt.Println(myArray[len(myArray)-1]) // shu orqalik  listing eng oxrgi qiymatni  ekranga chiqarish mumkin

	//myArray[6]=35  arraylarda xech qachon u elon qilingandan kegin unga qiymat qowib bo'lmaydi uni size belgilangach unga qowimcha qowib bolmedi

	myCars := [...]string{"Toyota", "Lexus", "Mazda", "Honda", "Nissan"}
	fmt.Println(len(myCars))
	myCars[2] = "Suzuki" // shu xolatdav ishidagi qiymatni ozgartrsa bo'ladi.
	fmt.Println(myCars)

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])

	}

}
