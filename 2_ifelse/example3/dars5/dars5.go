package main

import "fmt"

func main() {
	/*
		Harbiy qabul
		Jins Erkak Boy 180 dan baland
		Jins Ayol Boy 160 dan baland
		Output: Salom ism.Siz jinsdasiz.Boyingiz boy sm.Qabul qilindingiz.!
		Output: Salom ism.Siz jinsdasiz. Bo'yingiz boy sms.Qabul qilinmadingiz!!
	*/

	var ism, jins, boy = "", true, 0
	fmt.Println("Iltimos Ismingizni kiriting: ")
	fmt.Scanln(&ism)
	fmt.Println("Iltimos Jinsingizni kiriting: ")
	fmt.Scanln(&jins)
	fmt.Println("Iltimos boyingizni kiriting: ")
	fmt.Scanln(&boy)

	if (jins == "erkak" || jins == "Erkak") && boy >= 180 {
		fmt.Printf("Salom %v. Siz %v siz. Bo'yingiz %v sm\n",ism,jins,boy)
		fmt.Println("Qabul qilindingiz")
 
	} else if (jins == "erkak" || jins == "Erkak") && boy < 180 {
		fmt.Printf("Salom %v. Siz %v siz. Bo'yingiz %v sm\n",ism,jins,boy)
		fmt.Println("Qabul qilinmadingiz")
	}	 else if (jins == "ayol" || jins == "Ayol") && boy >= 160 {
		fmt.Printf("Salom %v. Siz %v siz. Bo'yingiz %v sm\n",ism,jins,boy)
		fmt.Println("Qabul qilindingiz")
 
	} else if (jins == "ayol" || jins == "Ayol") && boy < 160 {
		fmt.Printf("Salom %v. Siz %v siz. Bo'yingiz %v sm\n",ism,jins,boy)
		fmt.Println("Qabul qilinmadingiz")
	}

}
