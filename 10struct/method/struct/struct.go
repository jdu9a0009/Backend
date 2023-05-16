package main

import "fmt"

type StudentsData struct {
	Name, Surname, FullAddress, JapaneseLevel string
	Age, ID                                   int
}
type Square struct {
	x, y int
}

func main() {
	s1 := StudentsData{Name: "Abduqodir", Surname: "Abdullayev", FullAddress: "Andijon vil.Asaka tumani ,23 uy", JapaneseLevel: "N2", Age: 23, ID: 22211}
	fmt.Println(s1)
	data1 := s1.Name + " " + s1.Surname + " Yapon tilidan" + s1.JapaneseLevel + " darajaga ega"
	fmt.Println(data1)
	s2 := new(StudentsData)
	s2.Name = "Abdulla"
	s2.Surname = "Sadullayev"
	s2.FullAddress = "Tashkent Chilonozr 6, 10-uy"
	s2.JapaneseLevel = "N1"
	s2.Age = 22
	s2.ID = 2212112
	p1 := Square{1, 23}
	p2 := Square{x: 23} //y ga default qiymat beriladi
	fmt.Println(p1)
	fmt.Println(p2)
}
