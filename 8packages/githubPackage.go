package main

import (
	"github.com/fatih/color"
	"github.com/frap3r/hellomod"
)

//githubdan package import qilshdan oldin   go mod init qilw kerak kegin esa go get "name of package" yozib  undan song gu run qilinsa ishlatsa boladi.

func main() {
	hellomod.Eng()
	hellomod.Fra()
	hellomod.Ita()

	color.Red("Bu qizil rangda chiqadi.")
	color.Cyan("bu cyanda ")

}
