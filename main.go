package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/simbainnovation/strutil"
)
import "github.com/simbainnovation/tpkg"

func main() {
	fmt.Println(tpkg.Version)
	beego.Run()
	fmt.Println(strutil.Random())

}
