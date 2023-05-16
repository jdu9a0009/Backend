package main

import "fmt"

func main() {
	var yil int
	fmt.Println("Iltmos qaysi yilni kabisa ekanligini aniqlamoqchisiz")
	fmt.Scanln(&yil)
	if yil%400 == 0 {
		fmt.Println("Bu yil 400 yilda bir keladgon Kabisa yili: ", yil)

	} else if yil%100 == 0 {
		fmt.Println("Bu yil kabisa yil emas: ", yil)
	} else if yil%4 == 0 {
		fmt.Println("Bu yil kabisa yili: ", yil)
	} else {
		fmt.Println("BU yil kabisa yili emas: ", yil)
	}

}
