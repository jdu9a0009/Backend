package main

import "fmt"

func main() {
	var son, son1, son2, son3, son4, katta int
	// fmt.Println("Iltmos son1,son2,son3,son4,son5 larni ketma ketlikda kirting")
	// fmt.Scanf("%v%v%v%v%v", &son, &son1, &son2,son3,son4)
	fmt.Println("Iltmos 1-sondi ni kirting: ")
	fmt.Scanln(&son)
	fmt.Printf("Iltmos 2-sondi ni kirting: ")
	fmt.Scanln(&son1)
	fmt.Printf("Iltmos 3-sondi ni kirting: ")
	fmt.Scanln(&son2)
	fmt.Printf("Iltmos 4-sondi ni kirting: ")
	fmt.Scanln(&son3)
	fmt.Printf("Iltmos 5-sondi ni kirting: ")
	fmt.Scanln(&son4)

	katta = son
	if son1 > katta {
		katta = son1
	}
	if son2 > katta {
		katta = son2
	}
	if son3 > katta {
		katta = son3
	}
	if son4 > katta {
		katta = son4
	}
	fmt.Println("Kirtilgan 5ta  sonlar ichda eng kattasi: ", katta)

}

///2- eng yaxwi usul
/*

package main

import "fmt"
size:=5
var  numbers= make([]int,size)
for( i=0; i< size; i++){
 fmr.Scanln(&numbers)
}
maxnumber=numbers[0];
for( i=0; i<size;i++){
	 if (maxnumber<numbers[i]){
		maxnumbers=numbers[i]
	 }
fmt.Println(maxnumber)

}

*/
