// Swapping 2 variable values using pointers
package main

import "fmt"

func main() {

	ism1, ism2 := swap("Axmad", "Bobur")
	fmt.Println("Ismlar: ", ism1, ism2)

}

func swap(a, b string) (string, string) {
	return b, a // go da mutiple returning imkoni bor
}

/*func main() {
	ism1, ism2 := "Axmad", "Bobur"
	fmt.Println("Before swap: ", ism1, ism2)
	ism1, ism2 = swap(ism1, ism2)
	fmt.Println("After swap: ", ism1, ism2)
}

func swap(a, b string) (string, string) {
	return b, a
} */
