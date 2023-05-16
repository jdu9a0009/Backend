package main

var name string = "Saud" // Global
//suranme := "Abdulwahed"  bu xolatda tanitib bolmaydi
var name1, name2, name3 = "a", "b", "c"

var (
	// bu xolatda  ozgaruvchni nomi uzun bolsa ishlatiladi
	ism1 = "a"
	ism2 = "b"
	ism3 = "c"
)

// agar o'zgaruvchi func dan tashqarda  berilgan bolsa lekin ishlatilmasa ham  muammo bolmaydi
func main() {

	//VAR TYPE -> STRING , NUMBER, BOOLEAN
	var habar string = "Salom o'quvchi Habar 1"
	var habar1 string                //DECLARE avval tanitib
	habar1 = "Salom oquvchi habar 2" //  ASSIGNMENT kegin qiymat berish

	println(habar)
	println(habar1)
	println(habar)

}
