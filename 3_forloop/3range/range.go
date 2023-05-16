package main

import "fmt"

var daraja = []int{1, 2, 4, 8, 16, 32, 64, 128} // List
var users = []string{"Abdulla", "Sunnatilla", "Toshpo'lat"}
var cities = map[string]string{"UZS": "Uzbekistan", "YEN": "Yaponya", "USD": "USA"} // mapdan chapda key o'ngda qiymat  turadi o'rtasida ikkitalik nuqta bo'ladi
func main() {
	for i, value := range daraja {
		fmt.Printf(" 2 ning %d inchi darajasi = %d ga teng\n", i, value)
	}

	for _, v := range users {
		fmt.Printf("Value: %s\n", v) // \t har bir qiymatdan kegin bita tab tashaydi
	}

	for key, value := range cities {
		fmt.Println("Pul birliklari: ", key, "Davlat:", value)
	}

	for num value:= range(1,101):
    if num % 3 == 0 and num % 5 == 0:
        print("FizzBuzz: "+str(num))
    elif num % 3 == 0:
        print("Fizz:" +str(num))
    elif num % 5 == 0:
        print("Buzz:" +str(num))
    else:
        print(num)
}
