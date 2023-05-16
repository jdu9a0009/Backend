package main

import (
	"fmt"
	"github.com/NotebookInterface/BusinessComputers"
	"github.com/NotebookInterface/GamingComputers"
	"github.com/NotebookInterface/Simple"
)

func main() {
	b := new(BusinessComp.Hp)
	b.Brand = "Hp"
	b.Model = "Spectre 15"
	b.Cpu = "Intel Core i9-1250U"
	b.Ram = 32
	b.Hdd = 2048
	b.Sdd = 512
	b.Display = 15.6
	b.Weight = 1.57
	b.Intense = false

	g := new(GamersComp.Asus)
	g.Brand = "Asus"
	g.Model = "Rog Star 17"
	g.Cpu = "Intel Core i9-1250U"
	g.Ram = 64
	g.Hdd = 1024
	g.Sdd = 512
	g.Display = 14
	g.Weight = 2.45
	g.Intense = true

	s := new(Student.Acer)
	s.Brand = "Acer"
	s.Model = "Acer Aspire 3"
	s.Cpu = "Intel Celeron"
	s.Ram = 8
	s.Hdd = 1024
	s.Sdd = 128
	s.Display = 15.6
	s.Weight = 1.57
	s.Intense = false

	Student.CompExecute(s)
	BusinessComp.CompExecute(b)
	GamersComp.CompExecute(g)
	b2 := BusinessComp.NewHp("Hp", "Envy x360", "Intel Core i5-1120U", 16, 1024, 512, 14, 1.34)
	g2 := GamersComp.NewAsus("MSI", "KATANA", "Intel Core i9-1020G", 64, 2048, 1024, 14.0, 1.89)
	s2 := Student.NewAcer("Lenovo", "V15 G2 ITL", "Intel Celeron core i3-1020", 8, 1024, 128, 15.6, 1.89)
	fmt.Println(b2)
	fmt.Println(g2)
	fmt.Println(s2)
}
