package main

import (
	"fmt"
	"portale/controllers"
	automezzi "portale/modules/automezzi/controllers"
	_ "portale/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	validation.SetDefaultMessage(map[string]string{
		"Required":     "Il campo deve essere riempito",
		"Min":          "Il valore minimo consentito: %d",
		"Max":          "Il valore massimo consentito: %d",
		"Range":        "Il valore deve essere compreso tra %d e %d",
		"MinSize":      "Lunghezza minima consentita: %d",
		"MaxSize":      "Lunghezza massima consentita %d",
		"Length":       "Deve essere lungo %d",
		"Alpha":        "Deve contenere lettere",
		"Numeric":      "Deve contenere numeri",
		"AlphaNumeric": "Deve essere composto da lettere o numeri",
		"Match":        "Deve essere uguale a %s",
		"NoMatch":      "Deve essere diversa da %s",
		"AlphaDash":    "Esso deve essere costituito da lettere, numeri o simboli (-_)",
		"Email":        "Deve essere un indirizzo email valido",
		"IP":           "Deve essere un indirizzo IP valido",
		"Base64":       "Deve essere in formato Base64 corretto",
		"Mobile":       "Deve essere un numero di cellulare valido",
		"Tel":          "Deve essere un numero di telefono fisso valido",
		"Phone":        "Deve essere un numero di telefono fisso o cellulare valido ",
		"ZipCode":      "Deve essere un codice postale valido",
	})

	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "portaledb.db")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.EnableXSRF = true

}

func main() {

	beego.ErrorController(&controllers.ErrorController{})
	automezzi.InitializeModule()
	beego.Run()
}
