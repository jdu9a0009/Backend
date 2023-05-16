package main

import "fmt"

func main() {
	var son, son1, son2, katta int
	// fmt.Println("Iltmos son1,son2va son3 larni ketma ketlikda kirting")
	// fmt.Scanf("%v%v%v", &son, &son1, &son2)
	fmt.Println("Iltmos 1-sondi ni kirting: ")
	fmt.Scanln(&son)
	fmt.Printf("Iltmos 2-sondi ni kirting: ")
	fmt.Scanln(&son1)
	fmt.Printf("Iltmos 3-sondi ni kirting: ")
	fmt.Scanln(&son2)

	katta = son
	if son1 > katta {
		katta = son1
	}
	if son2 > katta {
		katta = son2
	}
	fmt.Println("Kirtilgan sonlar ichda eng kattasi: ", katta)

}
