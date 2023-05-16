package main

import "fmt"

func main() {
	states := make(map[string]string)
	fmt.Println("States: ", states) //tanitilib qiymat berilmasa uni qiymati NIL bo'ladi
	states["TAS"] = "Toshkent"
	states["AZN"] = "Andijan"
	states["KAS"] = "Qashqadaryo"
	fmt.Println("States: ", states)

	fmt.Println("State one: ", states["TAS"])
	fmt.Println("State one: ", states["AZN"]) //mapni ichdagi qaydur qiymatni ozini  chiqarsh uchun shundayb ishlatiladi

	kas := states["KAS"] // mapning ichdagi qiymatni boshqa bir o'zgaruvchga tenglash mumkin
	fmt.Println("O'zgaruvchiga: ", kas)

	delete(states, "XOR") // map ni ichdagi qiymatni o'chirish
	fmt.Println("States: ", states)

	for k, v := range states {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	myMap := make(map[int]string)
	myMap[11] = "O'n bir"
	myMap[22] = "Yigirma ikki"
	myMap[33] = "O'ttiz uch"

	keys := make([]string{len(states)}) // key nomli slice yaratdik
	i := 0
	for _, v := range states {
		keys[i] = v
		i++
	}
	fmt.Println("Key Slice: ", keys)
	//slice orqalik mapga kirib uni ichdagi key yoki valuelarini  slice ga tenglashtrish mumkin

}
