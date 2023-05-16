package mycar

import (
	"fmt"
	"strconv"
)

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

func CarExecute(c Carface) {
	fmt.Println("Car info: " + c.Information())

	message := ""
	isRun := c.Run()
	if isRun {
		message = "Ishlayapti..."
	} else {
		message = "Ishlamayapti!!!"
	}
	fmt.Println(message)

	msg := ""

	isStop := c.Stop()
	if isStop {
		msg = "To'xtagan.."
	} else {
		msg = "To'xtamagan!!!"
	}
	fmt.Println(msg)
}

type Car struct {
	Brand, Model, Color string
	Speed               int
	Price               float64
}

type SpecialProduction struct {
	Special bool
}

type BMW struct {
	Car
	SpecialProduction
}

func (_ BMW) Run() bool {
	return true
}
func (_ BMW) Stop() bool {
	return true
}
func (f BMW) Information() string {
	r := f.Brand + "\n" + f.Color + "\n" + f.Model + "\n" + strconv.Itoa(f.Speed) + strconv.FormatFloat(f.Price, 'f', -1, 64)
	a := "Yes"
	if f.Special {
		fmt.Println(a)
	}
	return r

}
func NewBMW(brand, model, color string, speed int, price float64, special bool) *BMW {
	b := &BMW{
		Car: Car{
			Brand: brand,
			Model: model,
			Color: color,
			Speed: speed,
			Price: price,
		},
		SpecialProduction: SpecialProduction{
			Special: special,
		},
	}
	return b
}
