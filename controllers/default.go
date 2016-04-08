package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//ErrorController manage error page
type ErrorController struct {
    MainController
}

//Error404 func for 404 error
func (c *ErrorController) Error404() {
    c.Data["content"] = "page not found"
    c.TplName = "404.tpl"
}

//Error500 func for 500 error
func (c *ErrorController) Error500() {
    c.Data["content"] = "internal server error"
    c.TplName = "500.tpl"
}

//ErrorDb func for db error
func (c *ErrorController) ErrorDb() {
    c.Data["content"] = "database is now down"
    c.TplName = "dberror.tpl"
}


//activeContnent function build the page
func (c *MainController) activeContent(view string) {
	c.Layout = "basic-layout.tpl"
	c.Data["domainname"] = beego.AppConfig.String("appcfgdomainname")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	//c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"

	sess := c.GetSession("portale")
	if sess != nil {
		c.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		c.Data["First"] = m["first"]
		c.Data["Admin"] = m["admin"]
		c.Data["IDkey"] = m["idkey"]
        fmt.Println(m["portale"])
		c.Data["Automezzi"] = m["automezzi"]
	}
}

//Get function check if user are logged
func (c *MainController) Get() {
	c.activeContent("index")
	//******** c page requires login
	sess := c.GetSession("portale")
	if sess == nil {
		c.Redirect("/login", 302)
		return
	}
  
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}



//Notice show flash message

func (c *MainController) Notice() {
    sess := c.GetSession("portale")
	if sess != nil {
        //notice if user is logged
    	c.activeContent("noticelog")   
    } else {
        //notice if user isn't logged
    	c.activeContent("notice")    
    }


	flash := beego.ReadFromRequest(&c.Controller)
    fmt.Printf(flash.Data["notice"])
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	}
}