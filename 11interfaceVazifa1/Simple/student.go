package Student

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

type Acer struct {
	Comp
	IntenseGames
}

func (_ Acer) forProgramming() bool {
	return false
}
func (_ Acer) forModeling() bool {
	return false
}

func (Aspire Acer) Information() string {
	r := Aspire.Brand + "\n" + Aspire.Model + "\n" + Aspire.Cpu + "\n" + strconv.Itoa(Aspire.Ram) + "\n" + strconv.Itoa(Aspire.Hdd) + "\n" + strconv.Itoa(Aspire.Sdd) + "\n" + strconv.FormatFloat(Aspire.Display, 'f', 1, 64) + "\n" + strconv.FormatFloat(Aspire.Weight, 'f', 1, 64)
	a := "Yes you can play Highl level Games"
	if Aspire.Intense {
		fmt.Println(a)
	}
	return r
}

func NewAcer(brand, model, cpu string, ram, hdd, sdd int, dispay, weight float64) *Acer {
	Aspire := *&Acer{
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
	return &Aspire
}
