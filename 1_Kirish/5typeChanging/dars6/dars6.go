package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "1"
	raqam, _ := strconv.Atoi(s)
	// bunda strconb orqalik 2ta  qiymat qaytadi bitasi stringni integerga ozgartrlgani yana bir esa ERROR  ,errorni esa _ bosh tanituvchi bilan otkazb yuborilayapti
	fmt.Println(raqam + 10)
	// int(s) bu suslda type ni ozgratrb bo'lmaydi
	n := 11
	println(float32(n))
	// bu xolatda  int ni float ga ozgartrsa boladi
	w := 12.9
	println(int(w))
}
