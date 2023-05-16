package main

import (
	"fmt"
	"math/rand" // raqamni random qiladgon kutubxona
	"time"      // vaqt boyicha kutubxona
)

func main() {
	fmt.Println("O'yin: 0 dan 10 gacha raqamlar orasidan men o'ylagan  raqamni toping")
	// o'yinchiga qanday oynash kerakligin  bildradi
	fmt.Println("Sizda faqatgina (3)marta taxmin qilish imkoniyatiz bor")

	source := rand.NewSource(time.Now().UnixNano())

	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)
	// odan 10 gacha bo'lgan sonlar orasdan bitasni  random qilib beradi

	var guess int

	for try := 1; try <= 3; try++ {
		// 3 marta  tsikl aylandi topsa topdi topmasa togri javobni chiqazadi
		fmt.Printf("Imkoniyat:  %d\n", try)
		// nechinchi imkoniyatdan foydalanayotganini ko'rsatib turadi
		fmt.Println("Iltimos o'z taxminizni kiriting")
		fmt.Scan(&guess)
		if guess < secretNumber {
			fmt.Printf("Sizni  taxminingiz xato,biroz kattaroq son o'ylab ko'ring\n")
		} else if guess > secretNumber {
			fmt.Printf("Sizni taxminingiz xato: biroz kichikroq  son o'ylab toping")
		} else {
			fmt.Printf("Siz to'g'ri topdingiz\n")
			break
		}
		if try == 3 {
			// 3ta imkoniyatdan kegin oyin tugaydi va tepadagi tsikl tugaydi
			fmt.Printf("O'yin tugadi!!\n")
			fmt.Printf("Tog'ri javob:  %d\n", secretNumber)
			break
		}
	}
}
