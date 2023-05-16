package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "A"
	a = "B"
	fmt.Println(a)
	//ozgaruvchini  qiymatini istalgan joyda ozgartrish mumkin
	const s string = "Hello"
	// s= "Salom"
	// ozgarmas ni esa bir martta qiymat bergach uni qiymatni ozgartrb bolmaydi
	// ozgarmasni  qiymatini bergach uni elon qilinmasa ham  muammo bolmaydi

	var ism string
	ism = "Hello"

	// const bir marta tnitiladi shunga uni avval tanitib kegin boshqa joyda  qiymatini berib bolmaydi   yuqordagi ozgaruvchidek qilib bolmaydi

	// var son1=22
	//const son2= son1
	// xech qachon constga ozgaruvchini qiymatini  qoyib bolmaydi
}
