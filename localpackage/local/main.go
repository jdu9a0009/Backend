package main

import (
	bp "birinchi/bpackage"
	s "birinchi/bpackage/say"
	"fmt"
)

// / agar  package lar bir joya bolsa va ular main pakgae da bolsa ularni import qiliw wart emas
func main() {
	bp.SayHello()
	fmt.Println(s.SayName())
}
