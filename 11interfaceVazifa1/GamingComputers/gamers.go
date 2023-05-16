package GamersComp

import (
	"fmt"
	"strconv"
)

type Suitable interface {
	forProgramming() bool
	forModeling() bool
	Information() string
}

func CompExecute(g Suitable) {
	fmt.Println("..................................")

	fmt.Println("Notebook full info: " + g.Information())
	message := ""
	isProg := g.forProgramming()
	if isProg {
		message = "Also suitable for Programming"
	} else {
		message = "not suitable for programming"
	}
	fmt.Println(message)

	msg := ""
	isModel := g.forModeling()
	if isModel {
		msg = "Also suitable for Programming"
	} else {
		message = "not suitable for programming"
	}
	fmt.Println(msg)
}

type Comp struct {
	Brand, Model, Cpu string
	Ram, Hdd, Sdd     int
	Display, Weight   float64
}

type IntenseGames struct {
	Intense bool
}

type Asus struct {
	Comp
	IntenseGames
}

func (_ Asus) forProgramming() bool {
	return true
}
func (_ Asus) forModeling() bool {
	return true
}

func (Rog Asus) Information() string {
	r := Rog.Brand + "\n" + Rog.Model + "\n" + Rog.Cpu + "\n" + strconv.Itoa(Rog.Ram) + "\n" + strconv.Itoa(Rog.Hdd) + "\n" + strconv.Itoa(Rog.Sdd) + "\n" + strconv.FormatFloat(Rog.Display, 'f', 1, 64) + "\n" + strconv.FormatFloat(Rog.Weight, 'f', 1, 64)
	a := "Yes you can play Highl level Games"
	if Rog.Intense {
		fmt.Println(a)
	}
	return r
}

func NewAsus(brand, model, cpu string, ram, hdd, sdd int, dispay, weight float64) *Asus {
	Rog := *&Asus{
		Comp: Comp{
			Brand:   brand,
			Model:   model,
			Cpu:     cpu,
			Ram:     ram,
			Hdd:     hdd,
			Sdd:     sdd,
			Display: dispay,
			Weight:  weight,
		}, IntenseGames: IntenseGames{},
	}
	return &Rog
}
