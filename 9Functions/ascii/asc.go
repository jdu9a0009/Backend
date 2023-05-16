package main

import (
	"fmt"
)

func main() {
	var mark string
	fmt.Println("3ta belgi kiriting")
	fmt.Scanln(&mark)

	chars := []rune(mark)
	count := 0
	for _, char := range chars {
		if isDigit(char) {
			count++
		}
	}

	fmt.Println("Belgilar orasida ", count, " ta son bor")
}

func isDigit(char rune) bool {
	return char >= 48 && char <= 57
}

/*func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}*/

/* boshqa usuli
package main

import (
	"fmt"
)

func main() {
	var a, b, c byte
	fmt.Println("Belgilarni birint ketin joy tshlab kiriting!")
	fmt.Scanf("%c %c %c", &a, &b, &c)

	count := 0
	if isDigit(a) {
		count++
	}
	if isDigit(b) {
		count++
	}
	if isDigit(c) {
		count++
	}

	fmt.Println("Belgilar orasida ", count, " ta son bor")
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
*/
