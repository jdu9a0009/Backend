package BusinessComp

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

type Hp struct {
	Comp
	IntenseGames
}

func (_ Hp) forProgramming() bool {
	return true
}
func (_ Hp) forModeling() bool {
	return false
}

func (Spectre Hp) Information() string {
	r := Spectre.Brand + "\n" + Spectre.Model + "\n" + Spectre.Cpu + "\n" + strconv.Itoa(Spectre.Ram) + "\n" + strconv.Itoa(Spectre.Hdd) + "\n" + strconv.Itoa(Spectre.Sdd) + "\n" + strconv.FormatFloat(Spectre.Display, 'f', 1, 64) + "\n" + strconv.FormatFloat(Spectre.Weight, 'f', 1, 64)
	a := "Yes you can play Highl level Games"
	if Spectre.Intense {
		fmt.Println(a)
	}
	return r
}

func NewHp(brand, model, cpu string, ram, hdd, sdd int, dispay, weight float64) *Hp {
	Spectre := *&Hp{
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
	return &Spectre
}
