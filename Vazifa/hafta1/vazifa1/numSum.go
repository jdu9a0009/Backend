package main

import "fmt"

func main() {
	var son, yuz int
	fmt.Println("3 xonalik son kiriting")
	fmt.Scanln(&son)

	yuz = (son / 100)
	on := (son - yuz*100) / 10
	yigindi := yuz + on + on%10

	if son > 99 && son < 1000 {
		fmt.Println("Siz kiritgan sonning sonlar yigindisi: ", yigindi, "ga teng")
	} else {
		fmt.Println("Iltmos 3 xonalik son kiriting")
	}

}

///2- eng yaxwi usul
/*

package main

import "fmt"
func main() {
	var number int
	var result int
		fmt.Println("3 xonalik son kiriting")
	fmt.Scanln(&number)

	for number>0{
		result += number%10
		number/=10
	}
	fmr.Println("Kiritilgan 3 xonalikn son yigindisi: ",result)
}
*/
