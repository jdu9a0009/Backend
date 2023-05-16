package main

import "fmt"

func main() {
	//ASCII
	kichikA := 'a'
	fmt.Println(kichikA)
	kattaA := 'A'
	fmt.Println(kattaA)

	harf := 't'
	if harf <= 90 && harf >= 65 {
		fmt.Println("Kiritlgan harf KATTA !")
	} else if harf <= 122 && harf >= 97 {
		fmt.Println("Kiritilgan harf KICHIK !!")
	}
}
