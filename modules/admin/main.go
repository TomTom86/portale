package main

import (
	_ "portale/modules/admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

