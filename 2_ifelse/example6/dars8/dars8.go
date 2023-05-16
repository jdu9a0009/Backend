package main

import "fmt"

func main() {
	//Kiritilgan 3ta sonning kopaytamsi yigindisi,ortalamasi,katta kichigi
	var s1, s2, s3, yigindi, kopaytma, ortalama, katta, kichik float32
	fmt.Println("Iltmos son1,son2va son3 larni ketma ketlikda kirting")
	fmt.Scanf("%v%v%v", &s1, &s2, &s3) //11 22 33
	yigindi = s1 + s2 + s3
	kopaytma = s1 * s2 * s3
	ortalama = yigindi / 3
	kichik = s1
	if s2 < kichik {
		kichik = s2
	}
	if s3 < kichik {
		kichik = s3
	}
	katta = s1
	if s2 > katta {
		katta = s2
	}
	if s3 > katta {
		katta = s3
	}
	fmt.Println("Yigindi: ", yigindi, "Ko'paytma: ", kopaytma, "O'rtalama: ", ortalama)
	fmt.Println("Katta Son: ", katta, "Kichik Son: ", kichik)
}
