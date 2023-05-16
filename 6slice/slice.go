package main

import "fmt"

func main() {
	myArray := [...]int{11, 22, 33, 44}
	mySlice := myArray[:]
	fmt.Println("Slice:", mySlice)
	mySlice[0] = 99
	fmt.Println("Slice: ", mySlice)

	myColors := [...]string{"red", "Yellow", "Green", "Blue", "Cyan"}
	myColorsSlice := myColors[1 : len(myColors)-1] // 1-dan oxrgdan bita oldngigacha chiqaradi
	fmt.Println("Color Slice:", myColorsSlice)

	//slice ni ozini xech qanday arraylarsz yaratish uchun make dan foydalaniladi

	nums := make([]int, 5) // slice tanitilib lekin qiymat berilmasa  uni ichi )0 bn to'ldrladi
	nums[1] = 11
	fmt.Println(nums)
	nums = append(nums, 112) // apend orqalik qoshimcha son qoshildi yani 6-qiymat bo'lib
	fmt.Println(nums)

	fmt.Println("Length", len(nums))     // uzunligi  nechtalgi
	fmt.Println("Capacity: ", cap(nums)) // sigimi  miqdori yana nechta son qo'shish mumkinlgi,xoz 10ta deydi agara biz 10tadan ortiq son qo'shsek uni Sigimi 20ga teng bo'lib qoladi
	myArray2 := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 111}
	mySlice2 := myArray2[:]
	mySlice3 := myArray2[:]

	fmt.Println("O'zgrishsiz")
	fmt.Println("My Array 2", myArray2)
	fmt.Println("My Slice 2", mySlice2)
	fmt.Println("My Slice 3", mySlice3)

	mySlice2[0] = 123

	fmt.Println("O'zgartirilgach")

	fmt.Println("My Array 2", myArray2)
	fmt.Println("My Slice 2", mySlice2) // qiymat o'zgartrlsa hammasg atasir etadi
	fmt.Println("My Slice 3", mySlice3)

}
