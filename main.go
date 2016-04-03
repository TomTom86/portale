package main

import (
    "fmt"
    "portale/controllers"
	_ "portale/routers"
	"github.com/astaxie/beego"
    
    "github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "portaledb.db")
	orm.Debug = true
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
    beego.BConfig.WebConfig.Session.SessionOn = true
    beego.BConfig.WebConfig.EnableXSRF = true
    beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

