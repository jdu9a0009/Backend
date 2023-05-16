package main

import "fmt"

func main() {
	a := 10
	b := 20
	qoshilganda := a + b
	fmt.Println(qoshilganda)

	modda := a % b // % bu mod xisoblanadi yani qoldiq lek kichik sondan katta sondi mod olinsa kichik sondi ozi qaytariladi

	fmt.Println(modda)
	kopayganda := a * b
	fmt.Println(kopayganda)

	jami := 0
	jami += 10 // jami = jami+10 ->10
	jami -= 1  // jami = jami-1 ->9
	jami *= 10 // jami = jami*10 ->90
	jami *= 10 // jami = jami/10 ->9

	fmt.Println(jami)
}
